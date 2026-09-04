package discourse

import (
	"github.com/raitucarp/gown/text"
)

// SentenceTopic summarizes the salient topic keywords for a sentence.
type SentenceTopic struct {
	SentenceID int      `json:"sentence_id"`
	Sentence   string   `json:"sentence"`
	Keywords   []string `json:"keywords"`
}

// TopicTransition measures the topic shift between consecutive sentences.
type TopicTransition struct {
	FromSentenceID int     `json:"from_sentence_id"`
	ToSentenceID   int     `json:"to_sentence_id"`
	Continuity     float64 `json:"continuity"` // 0.0 (complete drift) to 1.0 (identical topic)
	IsDrift        bool    `json:"is_drift"`
}

// TrackTopics extracts content keywords per sentence and calculates inter-sentential continuity.
func TrackTopics(documentText string) ([]SentenceTopic, []TopicTransition) {
	sentences := text.SentenceSegment(documentText)
	var topics []SentenceTopic

	for i, s := range sentences {
		keywords := text.ExtractContentWords(s)
		topics = append(topics, SentenceTopic{
			SentenceID: i + 1,
			Sentence:   s,
			Keywords:   keywords,
		})
	}

	var transitions []TopicTransition
	for i := 0; i < len(topics)-1; i++ {
		t1 := topics[i]
		t2 := topics[i+1]

		sim := text.JaccardSimilarity(t1.Keywords, t2.Keywords)
		transitions = append(transitions, TopicTransition{
			FromSentenceID: t1.SentenceID,
			ToSentenceID:   t2.SentenceID,
			Continuity:     sim,
			IsDrift:        sim < 0.15,
		})
	}

	return topics, transitions
}
