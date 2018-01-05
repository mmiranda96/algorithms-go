package sorting

// CountingSort sorts an integer slice using the counting algorithm, based on the Cormen implementation
func CountingSort(list []int) []int {
	result := make([]int, len(list))
	k := maxValue(list) + 1
	t := make([]int, k)

	for _, x := range list {
		t[x]++
	}

	for i := 1; i < k; i++ {
		t[i] += t[i-1]
	}

	for i := len(list); i > 0; i-- {
		result[t[list[i-1]]-1] = list[i-1]
		t[list[i-1]]--
	}

	return result
}
