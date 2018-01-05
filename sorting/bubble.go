package sorting

// BubbleSort sorts an integer slice using the naive bubble sort algoritm
func BubbleSort(list []int) []int {
	result := make([]int, len(list))
	copy(result, list)

	n := len(result)
	for i := 1; i < n; i++ {
		for j := 0; j < n-i; j++ {
			if result[j] > result[j+1] {
				swap(result, j, j+1) // Function declared and tested in quick.go
			}
		}
	}

	return result
}
