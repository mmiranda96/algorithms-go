package sorting

func InsertionSort(list []int) []int {
	result := make([]int, len(list))
	copy(result, list)

	n := len(result)
	for i := 0; i < n; i++ {
		for j := i; j > 0; j-- {
			if result[j] < result[j-1] {
				swap(result, j, j-1)
			}
		}
	}

	return result
}
