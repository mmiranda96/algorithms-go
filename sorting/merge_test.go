package sorting

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	t.Run("Full merge", func(t *testing.T) {
		list := []int{1, 2, 4, 7, 8, 0, 3, 5, 6, 9}
		p, q, r := 0, len(list)/2, len(list)
		merge(list, p, q, r)

		expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		if !reflect.DeepEqual(expected, list) {
			t.Errorf("Expected %+v, got %+v", expected, list)
		}
	})

	t.Run("Partial merge", func(t *testing.T) {
		list := []int{8, 4, 5, 6, 1, 7, 2, 3, 4, 8, 0, 9, 11, 3, 4, 1}
		p, q, r := 4, 6, 8
		merge(list, p, q, r)

		expected := []int{8, 4, 5, 6, 1, 2, 3, 7, 4, 8, 0, 9, 11, 3, 4, 1}
		if !reflect.DeepEqual(expected, list) {
			t.Errorf("Expected %+v, got %+v", expected, list)
		}
	})
}

func TestMergeSort(t *testing.T) {
	t.Run("Unique int list", func(t *testing.T) {
		list := []int{3, 5, 0, 4, 2, 6, 7, 1, 9, 8}
		list = MergeSort(list)

		expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		if !reflect.DeepEqual(expected, list) {
			t.Errorf("Expected %+v, got %+v", expected, list)
		}
	})

	t.Run("Repeating int list", func(t *testing.T) {
		list := []int{2, 3, 1, 2, 2, 1, 0, 1, 0, 3}
		list = MergeSort(list)

		expected := []int{0, 0, 1, 1, 1, 2, 2, 2, 3, 3}
		if !reflect.DeepEqual(expected, list) {
			t.Errorf("Expected %+v, got %+v", expected, list)
		}
	})
}
