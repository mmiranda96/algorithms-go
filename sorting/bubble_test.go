package sorting

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	t.Run("Unique int list", func(t *testing.T) {
		list := []int{3, 5, 0, 4, 2, 6, 7, 1, 9, 8}
		list = BubbleSort(list)

		expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		if !reflect.DeepEqual(expected, list) {
			t.Errorf("Expected %+v, got %+v", expected, list)
		}
	})

	t.Run("Repeating int list", func(t *testing.T) {
		list := []int{2, 3, 1, 2, 2, 1, 0, 1, 0, 3}
		list = BubbleSort(list)

		expected := []int{0, 0, 1, 1, 1, 2, 2, 2, 3, 3}
		if !reflect.DeepEqual(expected, list) {
			t.Errorf("Expected %+v, got %+v", expected, list)
		}
	})
}
