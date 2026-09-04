# Package similarity

Package `similarity` provides standard information-theoretic and path-based semantic similarity and relatedness metrics operating over WordNet synsets.

## Overview

Semantic similarity measures the degree to which two concepts share meaning or taxonomic ancestry (e.g. *dog* and *wolf* are highly similar, while *dog* and *leash* are related but not taxonomically similar). Package `similarity` provides:

1. **Path-Based Metrics**:
   - **Path Similarity**: Inverse of shortest path distance ($1 / \text{path\_len}$).
   - **Wu-Palmer (WuP)**: Ratio of the depth of the Lowest Common Subsumer (LCS) to the sum of synset depths:
     $$\text{Sim}_{\text{WuP}} = \frac{2 \times \text{depth}(\text{LCS})}{\text{depth}(s_1) + \text{depth}(s_2)}$$
   - **Leacock-Chodorow (LCH)**: Negative logarithm of path distance normalized by the maximum taxonomy depth:
     $$\text{Sim}_{\text{LCH}} = -\ln\left(\frac{\text{len}(s_1, s_2)}{2 \times \text{max\_depth}}\right)$$
2. **Information Content (IC) Metrics**:
   - **Seco et al. (2004) Intrinsic IC**: Computes Information Content directly from WordNet's hierarchical hyponym structure without requiring external corpus counts:
     $$\text{IC}(s) = 1 - \frac{\ln(\text{hypo}(s) + 1)}{\ln(\text{total\_synsets})}$$
   - **Resnik**: Information Content of the Lowest Common Subsumer: $\text{Sim}_{\text{Resnik}} = \text{IC}(\text{LCS})$.
   - **Lin**: Normalized ratio of common information to individual information:
     $$\text{Sim}_{\text{Lin}} = \frac{2 \times \text{IC}(\text{LCS})}{\text{IC}(s_1) + \text{IC}(s_2)}$$
   - **Jiang-Conrath (JCN)**: Distance derived from difference in information:
     $$\text{Dist}_{\text{JCN}} = \text{IC}(s_1) + \text{IC}(s_2) - 2 \times \text{IC}(\text{LCS})$$
3. **Surface String Similarity**:
   - Normalized Levenshtein distance for orthographic comparison.

## Key Types and Functions

- `Path(res *gown.LexicalResource, s1, s2 *gown.Synset) float64`: Path-based similarity (0.0 to 1.0).
- `WuPalmer(res *gown.LexicalResource, s1, s2 *gown.Synset) float64`: Wu-Palmer similarity metric (0.0 to 1.0).
- `LeacockChodorow(res *gown.LexicalResource, s1, s2 *gown.Synset) (float64, error)`: Leacock-Chodorow similarity.
- `NewInformationContent(res *gown.LexicalResource) *InformationContent`: Builds an in-memory IC calculator.
- `Resnik(res *gown.LexicalResource, s1, s2 *gown.Synset, ic *InformationContent) float64`: Resnik semantic similarity.
- `Lin(res *gown.LexicalResource, s1, s2 *gown.Synset, ic *InformationContent) float64`: Lin similarity (0.0 to 1.0).
- `JiangConrath(res *gown.LexicalResource, s1, s2 *gown.Synset, ic *InformationContent) float64`: Jiang-Conrath distance measure.

## Example

```go
package main

import (
    "fmt"
    "log"

    "github.com/raitucarp/gown"
    "github.com/raitucarp/gown/similarity"
)

func main() {
    res, err := gown.ReadLexicalResource()
    if err != nil {
        log.Fatal(err)
    }

    dogSyn := res.LookupNoun("dog")[0].Senses[0].GetSynset()
    wolfSyn := res.LookupNoun("wolf")[0].Senses[0].GetSynset()

    // 1. Path-based Wu-Palmer
    wup := similarity.WuPalmer(res, dogSyn, wolfSyn)
    fmt.Printf("Wu-Palmer(dog, wolf): %.4f\n", wup)

    // 2. Information Content Lin similarity
    ic := similarity.NewInformationContent(res)
    lin := similarity.Lin(res, dogSyn, wolfSyn, ic)
    fmt.Printf("Lin(dog, wolf): %.4f\n", lin)
}
```
