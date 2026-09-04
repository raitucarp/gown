package gown_test

import (
	"testing"

	"github.com/raitucarp/gown"
	"github.com/raitucarp/gown/similarity"
)

func BenchmarkLookupExact(b *testing.B) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		b.Fatalf("Failed to load resource: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = res.LookupExact("dog")
	}
}

func BenchmarkLookupMorphy(b *testing.B) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		b.Fatalf("Failed to load resource: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = res.Lookup("running", gown.WithPOS(gown.VerbPos))
	}
}

func BenchmarkOrthographicCV(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = gown.OrthographicCV("extravagant")
	}
}

func BenchmarkSimilarityWuPalmer(b *testing.B) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		b.Fatalf("Failed to load resource: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = similarity.Compare(res, "dog", "cat", similarity.WithMetric(similarity.MetricWuPalmer))
	}
}
