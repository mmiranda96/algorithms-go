package sorting

import "testing"
import "reflect"

func TestPartition(t *testing.T) {
	t.Run("Large pivot", func(t *testing.T) {
		list := []int{2, 4, 7, 1, 3, 5, 6, 8}
		pivot := partition(list, 0, len(list), nil)

		expectedPivot := 7
		if pivot != expectedPivot {
			t.Errorf("Expected %d, got %d", expectedPivot, pivot)
		}

		expectedList := []int{2, 4, 7, 1, 3, 5, 6, 8}
		if !reflect.DeepEqual(expectedList, list) {
			t.Errorf("Expected %+v, got %v", expectedList, list)
		}
	})

	t.Run("Medium movements", func(t *testing.T) {
		list := []int{2, 8, 7, 1, 3, 5, 6, 4}
		pivot := partition(list, 0, len(list), nil)

		expectedPivot := 3
		if pivot != expectedPivot {
			t.Errorf("Expected %d, got %d", expectedPivot, pivot)
		}

		expectedList := []int{2, 1, 3, 4, 7, 5, 6, 8}
		if !reflect.DeepEqual(expectedList, list) {
			t.Errorf("Expected %+v, got %v", expectedList, list)
		}
	})

	t.Run("Small pivot", func(t *testing.T) {
		list := []int{2, 8, 7, 4, 3, 5, 6, 1}
		pivot := partition(list, 0, len(list), nil)

		expectedPivot := 0
		if pivot != expectedPivot {
			t.Errorf("Expected %d, got %d", expectedPivot, pivot)
		}

		expectedList := []int{1, 8, 7, 4, 3, 5, 6, 2}
		if !reflect.DeepEqual(expectedList, list) {
			t.Errorf("Expected %+v, got %v", expectedList, list)
		}
	})
}

func TestQuickSort(t *testing.T) {
	t.Run("Unique int list", func(t *testing.T) {
		list := []int{3, 5, 0, 4, 2, 6, 7, 1, 9, 8}
		list = QuickSort(list)

		expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		if !reflect.DeepEqual(expected, list) {
			t.Errorf("Expected %+v, got %+v", expected, list)
		}
	})

	t.Run("Repeating int list", func(t *testing.T) {
		list := []int{2, 3, 1, 2, 2, 1, 0, 1, 0, 3}
		list = QuickSort(list)

		expected := []int{0, 0, 1, 1, 1, 2, 2, 2, 3, 3}
		if !reflect.DeepEqual(expected, list) {
			t.Errorf("Expected %+v, got %+v", expected, list)
		}
	})
}
