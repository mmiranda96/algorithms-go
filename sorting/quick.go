package sorting

import (
	"math/rand"
	"time"
)

func swap(list []int, x, y int) {
	t := list[x]
	list[x] = list[y]
	list[y] = t
}

func partition(list []int, p, r int, g *rand.Rand) int {
	if g != nil && p < r {
		swap(list, p+g.Intn(r-p), r-1)
	}

	pivot := r - 1
	x := list[pivot]
	i := p - 1
	for j := p; j < pivot; j++ {
		if list[j] <= x {
			i++
			swap(list, i, j)
		}
	}
	i++
	swap(list, i, pivot)

	return i
}

func quickSortAux(list []int, p, r int, g *rand.Rand) {
	if p < r {
		q := partition(list, p, r, g)
		quickSortAux(list, p, q, g)
		quickSortAux(list, q+1, r, g)
	}
}

// QuickSort sorts an integer slice using the divide-and-conquer partition algorithm, based on the Cormen version
// This version chooses the last element of the subslice as pivot
func QuickSort(list []int) []int {
	result := make([]int, len(list))
	copy(result, list)
	quickSortAux(result, 0, len(result), nil)

	return result
}

// QuickSortRand sorts an integer slice using the divide-and-conquer partition algorithm, based on the Cormen version.
// This version chooses a random element of the subslice as pivot
func QuickSortRand(list []int) []int {
	result := make([]int, len(list))
	copy(result, list)
	g := rand.New(rand.NewSource(time.Now().UnixNano()))
	quickSortAux(result, 0, len(result), g)

	return result
}
