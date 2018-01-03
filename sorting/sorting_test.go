package sorting

import (
	"math/rand"
	"testing"
	"time"
)

type sortFunc func([]int) []int

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func buildRandomizedList(elements int) []int {
	list := make([]int, elements, elements)

	for i := 0; i < elements; i++ {
		list[i] = r.Int()
	}

	return list
}

func benchmarkSort(i int, f sortFunc, b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		list := buildRandomizedList(i)

		b.StartTimer()
		list = f(list)
	}
}

// HeapSort
func BenchmarkHeapSort10(b *testing.B)     { benchmarkSort(10, HeapSort, b) }
func BenchmarkHeapSort100(b *testing.B)    { benchmarkSort(100, HeapSort, b) }
func BenchmarkHeapSort1000(b *testing.B)   { benchmarkSort(1000, HeapSort, b) }
func BenchmarkHeapSort10000(b *testing.B)  { benchmarkSort(10000, HeapSort, b) }
func BenchmarkHeapSort100000(b *testing.B) { benchmarkSort(100000, HeapSort, b) }

// MergeSort
func BenchmarkMergeSort10(b *testing.B)     { benchmarkSort(10, MergeSort, b) }
func BenchmarkMergeSort100(b *testing.B)    { benchmarkSort(100, MergeSort, b) }
func BenchmarkMergeSort1000(b *testing.B)   { benchmarkSort(1000, MergeSort, b) }
func BenchmarkMergeSort10000(b *testing.B)  { benchmarkSort(10000, MergeSort, b) }
func BenchmarkMergeSort100000(b *testing.B) { benchmarkSort(100000, MergeSort, b) }
