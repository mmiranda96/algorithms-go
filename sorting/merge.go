package sorting

func merge(list []int, p, q, r int) {
	x := q - p
	y := r - q

	left := make([]int, x)
	right := make([]int, y)

	copy(left, list[p:q])
	copy(right, list[q:r])

	for i, j, k := 0, 0, p; k < r; k++ {
		if !(j < y) || (i < x && left[i] < right[j]) {
			list[k] = left[i]
			i++
		} else {
			list[k] = right[j]
			j++
		}
	}
}

func mergeSortAux(list []int, p, r int) {
	if p+1 < r {
		q := (p + r) / 2

		// Divide
		mergeSortAux(list, p, q)
		mergeSortAux(list, q, r)

		// Conquer
		merge(list, p, q, r)
	}
}

// MergeSort sorts an integer slice using the divide-and-conquer merge algorithm, based on the Cormen version
func MergeSort(list []int) []int {
	result := make([]int, len(list))
	copy(result, list)
	mergeSortAux(result, 0, len(result))

	return result
}
