package sorting

import (
	"reflect"
	"testing"
)

func TestLeft(t *testing.T) {
	h := heap([]int{5, 4, 1, 2, 3, 0})

	t.Run("Existing left child", func(t *testing.T) {
		left := h.left(1)

		expected := 2
		if expected != left {
			t.Errorf("Expected %d, got %d", expected, left)
		}
	})

	t.Run("Non-existing left child", func(t *testing.T) {
		left := h.left(3)

		expected := 0
		if expected != left {
			t.Errorf("Expected %d, got %d", expected, left)
		}
	})
}

func TestRight(t *testing.T) {
	h := heap([]int{5, 4, 1, 2, 3, 0})

	t.Run("Existing right child", func(t *testing.T) {
		left := h.right(1)

		expected := 3
		if expected != left {
			t.Errorf("Expected %d, got %d", expected, left)
		}
	})

	t.Run("Non-existing right child", func(t *testing.T) {
		left := h.right(3)

		expected := 0
		if expected != left {
			t.Errorf("Expected %d, got %d", expected, left)
		}
	})
}

func TestSwap(t *testing.T) {
	h := heap([]int{5, 4, 1, 2, 3, 0})

	t.Run("Happy path", func(t *testing.T) {
		h.swap(1, 2)

		expected := heap([]int{5, 1, 4, 2, 3, 0})
		if !reflect.DeepEqual(expected, h) {
			t.Errorf("Expected %+v, got %+v", expected, h)
		}
	})
}

func TestHeapify(t *testing.T) {

}

func TestBuildHeap(t *testing.T) {
	t.Run("Unique int list", func(t *testing.T) {
		list := []int{3, 5, 0, 4, 2, 6, 7, 1, 9, 8}
		h := buildHeap(list)

		expected := heap([]int{10, 9, 8, 7, 5, 3, 6, 0, 1, 4, 2})
		if !reflect.DeepEqual(expected, h) {
			t.Errorf("Expected %+v, got %+v", expected, h)
		}
	})

	t.Run("Repeating int list", func(t *testing.T) {
		list := []int{2, 3, 1, 2, 2, 1, 0, 1, 0, 3}
		h := buildHeap(list)

		expected := heap([]int{10, 3, 3, 1, 2, 2, 1, 0, 1, 0, 2})
		if !reflect.DeepEqual(expected, h) {
			t.Errorf("Expected %+v, got %+v", expected, h)
		}
	})
}

func TestHeapSort(t *testing.T) {
	t.Run("Unique int list", func(t *testing.T) {
		list := []int{3, 5, 0, 4, 2, 6, 7, 1, 9, 8}
		list = HeapSort(list)

		expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		if !reflect.DeepEqual(expected, list) {
			t.Errorf("Expected %+v, got %+v", expected, list)
		}
	})

	t.Run("Repeating int list", func(t *testing.T) {
		list := []int{2, 3, 1, 2, 2, 1, 0, 1, 0, 3}
		list = HeapSort(list)

		expected := []int{0, 0, 1, 1, 1, 2, 2, 2, 3, 3}
		if !reflect.DeepEqual(expected, list) {
			t.Errorf("Expected %+v, got %+v", expected, list)
		}
	})
}
