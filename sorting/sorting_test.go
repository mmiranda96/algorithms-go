package sorting

import (
	"math/rand"
	gosort "sort"
	"testing"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func buildRandomizedList(elements int) []int {
	list := make([]int, elements, elements)

	for i := 0; i < elements; i++ {
		list[i] = r.Int()
	}

	return list
}

func BenchmarkHeapSort(b *testing.B) {
	b.Run("Randomized int list", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			list := buildRandomizedList(i)
			list = HeapSort(list)

			isSorted := gosort.IntsAreSorted(list)
			if !isSorted {
				b.Errorf("Failed sorting with %d elements", i)
			}
		}
	})
}
