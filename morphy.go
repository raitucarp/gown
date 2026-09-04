package gown

import (
	"strings"
)

// Morphy returns possible base lemma forms for a given surface word form and part of speech.
// It applies irregular morphological mapping, affix substitution rules, and validates against WordNet.
func (resource *LexicalResource) Morphy(word string, pos POS) []string {
	word = strings.TrimSpace(strings.ToLower(word))
	if word == "" {
		return nil
	}

	seen := make(map[string]bool)
	var results []string

	addResult := func(candidate string) {
		if candidate != "" && !seen[candidate] {
			// Verify if candidate exists in WordNet under the given POS
			if resource.hasEntryWithPos(candidate, pos) {
				seen[candidate] = true
				results = append(results, candidate)
			}
		}
	}

	// 1. Direct check: Is it already in WordNet?
	addResult(word)

	// 2. Check irregular exceptions table
	if irregulars, ok := irregularLookup(word, pos); ok {
		for _, irr := range irregulars {
			addResult(irr)
		}
	}

	// 3. Rule-based affix substitutions
	rules := getMorphRules(pos)
	for _, rule := range rules {
		if strings.HasSuffix(word, rule.suffix) {
			base := strings.TrimSuffix(word, rule.suffix) + rule.replacement
			addResult(base)

			// Also handle consonant doubling (e.g. running -> runn -> run, stopped -> stopp -> stop)
			if len(base) > 2 && base[len(base)-1] == base[len(base)-2] {
				undoubled := base[:len(base)-1]
				addResult(undoubled)
			}
		}
	}

	// If no validated lemma found in WordNet, return rule transformations anyway if nothing matched
	if len(results) == 0 {
		if irregulars, ok := irregularLookup(word, pos); ok && len(irregulars) > 0 {
			return irregulars
		}
		for _, rule := range rules {
			if strings.HasSuffix(word, rule.suffix) {
				base := strings.TrimSuffix(word, rule.suffix) + rule.replacement
				if !seen[base] {
					seen[base] = true
					results = append(results, base)
				}
			}
		}
	}

	if len(results) == 0 {
		results = append(results, word)
	}

	return results
}

// MorphyAll finds candidate lemmas across all parts of speech.
func (resource *LexicalResource) MorphyAll(word string) []string {
	word = strings.TrimSpace(strings.ToLower(word))
	if word == "" {
		return nil
	}

	seen := make(map[string]bool)
	var results []string

	poses := []POS{NounPos, VerbPos, AdjectivePos, AdverbPos}
	for _, p := range poses {
		for _, candidate := range resource.Morphy(word, p) {
			if !seen[candidate] {
				seen[candidate] = true
				results = append(results, candidate)
			}
		}
	}

	return results
}

func (resource *LexicalResource) hasEntryWithPos(word string, pos POS) bool {
	entries := resource.entriesByLemmaLower[strings.ToLower(word)]
	if pos == "" {
		return len(entries) > 0
	}
	for _, e := range entries {
		if POS(e.Lemma.PartOfSpeech) == pos {
			return true
		}
		if (pos == AdjectivePos || pos == AdjectiveSatellitePos) &&
			(e.Lemma.PartOfSpeech == string(AdjectivePos) || e.Lemma.PartOfSpeech == string(AdjectiveSatellitePos)) {
			return true
		}
	}
	return false
}

type morphRule struct {
	suffix      string
	replacement string
}

func getMorphRules(pos POS) []morphRule {
	switch pos {
	case NounPos:
		return []morphRule{
			{"s", ""},
			{"ses", "s"},
			{"xes", "x"},
			{"zes", "z"},
			{"ches", "ch"},
			{"shes", "sh"},
			{"men", "man"},
			{"ies", "y"},
			{"ves", "f"},
			{"ves", "fe"},
			{"ices", "ix"},
			{"ices", "ex"},
			{"axes", "axis"},
			{"ina", "en"},
			{"a", "um"},
			{"a", "on"},
			{"i", "us"},
		}
	case VerbPos:
		return []morphRule{
			{"s", ""},
			{"ies", "y"},
			{"es", "e"},
			{"es", ""},
			{"ed", "e"},
			{"ed", ""},
			{"ing", "e"},
			{"ing", ""},
		}
	case AdjectivePos, AdjectiveSatellitePos:
		return []morphRule{
			{"er", ""},
			{"er", "e"},
			{"est", ""},
			{"est", "e"},
			{"ier", "y"},
			{"iest", "y"},
		}
	case AdverbPos:
		return []morphRule{
			{"ly", ""},
			{"ily", "y"},
		}
	default:
		return nil
	}
}

func irregularLookup(word string, pos POS) ([]string, bool) {
	switch pos {
	case NounPos:
		if val, ok := nounIrregulars[word]; ok {
			return val, true
		}
	case VerbPos:
		if val, ok := verbIrregulars[word]; ok {
			return val, true
		}
	case AdjectivePos, AdjectiveSatellitePos:
		if val, ok := adjIrregulars[word]; ok {
			return val, true
		}
	case AdverbPos:
		if val, ok := advIrregulars[word]; ok {
			return val, true
		}
	default:
		var res []string
		if v, ok := nounIrregulars[word]; ok {
			res = append(res, v...)
		}
		if v, ok := verbIrregulars[word]; ok {
			res = append(res, v...)
		}
		if v, ok := adjIrregulars[word]; ok {
			res = append(res, v...)
		}
		if v, ok := advIrregulars[word]; ok {
			res = append(res, v...)
		}
		if len(res) > 0 {
			return res, true
		}
	}
	return nil, false
}

var nounIrregulars = map[string][]string{
	"children":    {"child"},
	"men":         {"man"},
	"women":       {"woman"},
	"mice":        {"mouse"},
	"teeth":       {"tooth"},
	"feet":        {"foot"},
	"geese":       {"goose"},
	"oxen":        {"ox"},
	"criteria":    {"criterion"},
	"phenomena":   {"phenomenon"},
	"data":        {"datum"},
	"media":       {"medium"},
	"matrices":    {"matrix"},
	"vertices":    {"vertex"},
	"indices":     {"index"},
	"appendices":  {"appendix"},
	"hypotheses":  {"hypothesis"},
	"theses":      {"thesis"},
	"crises":      {"crisis"},
	"analyses":    {"analysis"},
	"diagnoses":   {"diagnosis"},
	"bases":       {"basis", "base"},
	"oases":       {"oasis"},
	"parentheses": {"parenthesis"},
	"synopses":    {"synopsis"},
	"cacti":       {"cactus"},
	"fungi":       {"fungus"},
	"nuclei":      {"nucleus"},
	"radii":       {"radius"},
	"stimuli":     {"stimulus"},
	"syllabi":     {"syllabus"},
	"alumni":      {"alumnus"},
	"alumnae":     {"alumna"},
	"larvae":      {"larva"},
	"algae":       {"alga"},
	"vertebrae":   {"vertebra"},
	"curricula":   {"curriculum"},
	"millennia":   {"millennium"},
	"wolves":      {"wolf"},
	"calves":      {"calf"},
	"halves":      {"half"},
	"knives":      {"knife"},
	"lives":       {"life"},
	"leaves":      {"leaf"},
	"loaves":      {"loaf"},
	"scarves":     {"scarf"},
	"thieves":     {"thief"},
	"wives":       {"wife"},
	"elves":       {"elf"},
	"shelves":     {"shelf"},
	"selves":      {"self"},
	"sheaves":     {"sheaf"},
	"people":      {"person"},
}

var verbIrregulars = map[string][]string{
	"went":       {"go"},
	"gone":       {"go"},
	"ran":        {"run"},
	"ate":        {"eat"},
	"eaten":      {"eat"},
	"wrote":      {"write"},
	"written":    {"write"},
	"spoke":      {"speak"},
	"spoken":     {"speak"},
	"took":       {"take"},
	"taken":      {"take"},
	"gave":       {"give"},
	"given":      {"give"},
	"saw":        {"see"},
	"seen":       {"see"},
	"knew":       {"know"},
	"known":      {"know"},
	"brought":    {"bring"},
	"bought":     {"buy"},
	"thought":    {"think"},
	"caught":     {"catch"},
	"taught":     {"teach"},
	"fought":     {"fight"},
	"sought":     {"seek"},
	"found":      {"find"},
	"built":      {"build"},
	"sent":       {"send"},
	"lent":       {"lend"},
	"spent":      {"spend"},
	"left":       {"leave"},
	"met":        {"meet"},
	"led":        {"lead"},
	"read":       {"read"},
	"stood":      {"stand"},
	"understood": {"understand"},
	"sat":        {"sit"},
	"set":        {"set"},
	"put":        {"put"},
	"cut":        {"cut"},
	"hit":        {"hit"},
	"hurt":       {"hurt"},
	"let":        {"let"},
	"cost":       {"cost"},
	"shut":       {"shut"},
	"beat":       {"beat"},
	"beaten":     {"beat"},
	"bit":        {"bite"},
	"bitten":     {"bite"},
	"hid":        {"hide"},
	"hidden":     {"hide"},
	"chose":      {"choose"},
	"chosen":     {"choose"},
	"drove":      {"drive"},
	"driven":     {"drive"},
	"rode":       {"ride"},
	"ridden":     {"ride"},
	"rose":       {"rise"},
	"risen":      {"rise"},
	"broke":      {"break"},
	"broken":     {"break"},
	"fell":       {"fall"},
	"fallen":     {"fall"},
	"shook":      {"shake"},
	"shaken":     {"shake"},
	"threw":      {"throw"},
	"thrown":     {"throw"},
	"grew":       {"grow"},
	"grown":      {"grow"},
	"blew":       {"blow"},
	"blown":      {"blow"},
	"drew":       {"draw"},
	"drawn":      {"draw"},
	"flew":       {"fly"},
	"flown":      {"fly"},
	"wore":       {"wear"},
	"worn":       {"wear"},
	"tore":       {"tear"},
	"torn":       {"tear"},
	"swore":      {"swear"},
	"sworn":      {"swear"},
	"began":      {"begin"},
	"begun":      {"begin"},
	"swam":       {"swim"},
	"swum":       {"swim"},
	"sang":       {"sing"},
	"sung":       {"sing"},
	"sank":       {"sink"},
	"sunk":       {"sink"},
	"drank":      {"drink"},
	"drunk":      {"drink"},
	"rang":       {"ring"},
	"rung":       {"ring"},
	"sprung":     {"spring"},
	"shrank":     {"shrink"},
	"shrunk":     {"shrink"},
	"spun":       {"spin"},
	"swung":      {"swing"},
	"hung":       {"hang"},
	"struck":     {"strike"},
	"shone":      {"shine"},
	"crept":      {"creep"},
	"slept":      {"sleep"},
	"wept":       {"weep"},
	"swept":      {"sweep"},
	"kept":       {"keep"},
	"felt":       {"feel"},
	"knelt":      {"kneel"},
	"dreamt":     {"dream"},
	"burnt":      {"burn"},
	"learnt":     {"learn"},
	"dealt":      {"deal"},
	"meant":      {"mean"},
	"heard":      {"hear"},
	"said":       {"say"},
	"paid":       {"pay"},
	"laid":       {"lay"},
	"sold":       {"sell"},
	"told":       {"tell"},
	"lost":       {"lose"},
	"shot":       {"shoot"},
	"got":        {"get"},
	"gotten":     {"get"},
	"forgot":     {"forget"},
	"forgotten":  {"forget"},
	"became":     {"become"},
	"came":       {"come"},
	"overcame":   {"overcome"},
	"done":       {"do"},
	"did":        {"do"},
	"does":       {"do"},
	"been":       {"be"},
	"was":        {"be"},
	"were":       {"be"},
	"is":         {"be"},
	"are":        {"be"},
	"am":         {"be"},
	"had":        {"have"},
	"has":        {"have"},
}

var adjIrregulars = map[string][]string{
	"better":   {"good", "well"},
	"best":     {"good", "well"},
	"worse":    {"bad", "ill"},
	"worst":    {"bad", "ill"},
	"further":  {"far"},
	"furthest": {"far"},
	"farther":  {"far"},
	"farthest": {"far"},
	"less":     {"little"},
	"least":    {"little"},
	"more":     {"many", "much"},
	"most":     {"many", "much"},
}

var advIrregulars = map[string][]string{
	"better":   {"well"},
	"best":     {"well"},
	"worse":    {"badly"},
	"worst":    {"badly"},
	"further":  {"far"},
	"furthest": {"far"},
	"farther":  {"far"},
	"farthest": {"far"},
	"more":     {"much"},
	"most":     {"much"},
	"less":     {"little"},
	"least":    {"little"},
}
