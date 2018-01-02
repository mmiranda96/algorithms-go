package sorting

type heap []int

func (h heap) left(i int) int {
	if 2*i > h[0] {
		return 0
	}

	return 2 * i
}

func (h heap) right(i int) int {
	if 2*i+1 > h[0] {
		return 0
	}

	return 2*i + 1
}

func (h heap) swap(x, y int) {
	t := h[x]
	h[x] = h[y]
	h[y] = t
}

func (h heap) heapify(x int) {
	var largest int
	y, z := h.left(x), h.right(x)

	if y > 0 && h[y] > h[x] {
		largest = y
	} else {
		largest = x
	}

	if z > 0 && h[z] > h[largest] {
		largest = z
	}

	if largest != x {
		h.swap(x, largest)
		h.heapify(largest)
	}
}

func buildHeap(list []int) heap {
	// Heap size is stored in the first value
	h := heap(append([]int{len(list)}, list...))

	for i := (len(list)) / 2; i > 0; i-- {
		h.heapify(i)
	}

	return h
}

// HeapSort sorts an integer slice using a heap structure, based on the Cormen algorithm
func HeapSort(list []int) []int {
	h := buildHeap(list)
	for i := h[0]; i > 1; i-- {
		h.swap(1, i)
		h[0]--
		h.heapify(1)
	}

	return []int(h)[1:]
}
