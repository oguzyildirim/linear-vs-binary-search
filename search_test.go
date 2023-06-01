package main

import (
	"math/rand"
	"sort"
	"testing"
)

func BenchmarkLinearSearch(b *testing.B) {
	const n = 128
	v := make([]int, 0, n)

	rand.Seed(0xDEADC0DE)
	for i := 0; i < n; i++ {
		v = append(v, rand.Intn(1_000_000))
	}
	sort.Ints(v)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		num := rand.Intn(1_000_000)
		for _, val := range v {
			if val >= num {
				break
			}
		}
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	const n = 128
	v := make([]int, 0, n)

	rand.Seed(0xDEADC0DE)
	for i := 0; i < n; i++ {
		v = append(v, rand.Intn(1_000_000))
	}
	sort.Ints(v)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		randNum := rand.Intn(1_000_000)
		sort.Search(len(v), func(j int) bool {
			return v[j] >= randNum
		})
	}
}
