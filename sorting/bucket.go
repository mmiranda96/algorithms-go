package sorting

// BucketSort sorts an integer slice using the bucket algorithm, based on the Cormen implementation
func BucketSort(list []int) []int {
	n := len(list)
	max := float64(maxValue(list))
	buckets := make([][]int, n)
	for _, x := range list {
		i := int(float64(n-1) * float64(x) / max)
		buckets[i] = append(buckets[i], x)
	}

	for i, bucket := range buckets {
		buckets[i] = InsertionSort(bucket)
	}

	result := make([]int, 0)
	for _, bucket := range buckets {
		result = append(result, bucket...)
	}

	return result
}
