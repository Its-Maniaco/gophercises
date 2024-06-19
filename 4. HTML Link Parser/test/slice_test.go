package test

import (
	"slices"
	"testing"
)

func BenchmarkAppend(b *testing.B) {
	tester := make([]int, 0)
	for i := 0; i != b.N; i++ {
		_ = append(tester, i)
	}
	b.ReportAllocs()
}

func BenchmarkConcat(b *testing.B) {
	tester := make([]int, 0)
	for i := 0; i != b.N; i++ {
		_ = slices.Concat(tester, []int{i})
	}
	b.ReportAllocs()
}
