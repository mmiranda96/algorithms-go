package sorting

import (
	"math"

	"github.com/mmiranda96/algorithms-go/queue"
)

func maxValue(list []int) int {
	max := list[0]
	for _, i := range list {
		if i > max {
			max = i
		}
	}

	return max
}

func radix(x, d int) int {
	if d < 0 {
		return 0
	}

	return (x / int(math.Pow10(d))) % 10
}

// RadixSort sorts an integer slice using the radix algorithm, based on the Cormen implementation
func RadixSort(list []int) []int {
	result := make([]int, len(list))
	copy(result, list)

	queues := make([]*queue.Queue, 10)
	for i := 0; i < 10; i++ {
		queues[i] = &queue.Queue{}
		queues[i].Init()
	}

	d := int(math.Ceil(math.Log10(float64(maxValue(result)))))
	for i := 0; i < d; i++ {
		for _, x := range result {
			r := radix(x, i)
			queues[r].Enqueue(x)
		}

		for j, k := 0, 0; j < 10; j++ {
			for !queues[j].IsEmpty() {
				x, _ := queues[j].Dequeue()
				result[k] = x.(int)
				k++
			}
		}
	}

	return result
}
