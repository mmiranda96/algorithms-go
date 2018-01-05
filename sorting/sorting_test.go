package sorting

import (
	"math/rand"
	"sort"
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
		b.StopTimer()
		if !sort.IntsAreSorted(list) {
			b.Error("List not sorted")
		}
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

// QuickSort
func BenchmarkQuickSort10(b *testing.B)     { benchmarkSort(10, QuickSort, b) }
func BenchmarkQuickSort100(b *testing.B)    { benchmarkSort(100, QuickSort, b) }
func BenchmarkQuickSort1000(b *testing.B)   { benchmarkSort(1000, QuickSort, b) }
func BenchmarkQuickSort10000(b *testing.B)  { benchmarkSort(10000, QuickSort, b) }
func BenchmarkQuickSort100000(b *testing.B) { benchmarkSort(100000, QuickSort, b) }

// QuickSortRand
func BenchmarkQuickSortRand10(b *testing.B)     { benchmarkSort(10, QuickSortRand, b) }
func BenchmarkQuickSortRand100(b *testing.B)    { benchmarkSort(100, QuickSortRand, b) }
func BenchmarkQuickSortRand1000(b *testing.B)   { benchmarkSort(1000, QuickSortRand, b) }
func BenchmarkQuickSortRand10000(b *testing.B)  { benchmarkSort(10000, QuickSortRand, b) }
func BenchmarkQuickSortRand100000(b *testing.B) { benchmarkSort(100000, QuickSortRand, b) }

// BubbleSort
func BenchmarkBubbleSort10(b *testing.B)     { benchmarkSort(10, BubbleSort, b) }
func BenchmarkBubbleSort100(b *testing.B)    { benchmarkSort(100, BubbleSort, b) }
func BenchmarkBubbleSort1000(b *testing.B)   { benchmarkSort(1000, BubbleSort, b) }
func BenchmarkBubbleSort10000(b *testing.B)  { benchmarkSort(10000, BubbleSort, b) }
func BenchmarkBubbleSort100000(b *testing.B) { benchmarkSort(100000, BubbleSort, b) }

// InsertionSort
func BenchmarkInsertionSort10(b *testing.B)     { benchmarkSort(10, InsertionSort, b) }
func BenchmarkInsertionSort100(b *testing.B)    { benchmarkSort(100, InsertionSort, b) }
func BenchmarkInsertionSort1000(b *testing.B)   { benchmarkSort(1000, InsertionSort, b) }
func BenchmarkInsertionSort10000(b *testing.B)  { benchmarkSort(10000, InsertionSort, b) }
func BenchmarkInsertionSort100000(b *testing.B) { benchmarkSort(100000, InsertionSort, b) }
