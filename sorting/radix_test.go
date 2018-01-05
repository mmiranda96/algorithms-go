package sorting

import (
	"reflect"
	"testing"
)

func TestMaxValue(t *testing.T) {
	t.Run("Non-repeating list", func(t *testing.T) {
		list := []int{0, 1, 3, 2, 4, 5, 7, 6, 8, 9}
		max := maxValue(list)

		expected := 9
		if max != expected {
			t.Errorf("Expected %d, got %d", expected, max)
		}
	})

	t.Run("Repeating list", func(t *testing.T) {
		list := []int{0, 1, 1, 3, 1, 2, 1, 3, 2, 1, 3}
		max := maxValue(list)

		expected := 3
		if max != expected {
			t.Errorf("Expected %d, got %d", expected, max)
		}

	})
}

func TestRadix(t *testing.T) {
	t.Run("Valid radix", func(t *testing.T) {
		value := 103221
		r := radix(value, 3)

		expected := 3
		if r != expected {
			t.Errorf("Expected %d, got %d", expected, r)
		}
	})

	t.Run("Zero r", func(t *testing.T) {
		value := 103221
		r := radix(value, 0)

		expected := 1
		if r != expected {
			t.Errorf("Expected %d, got %d", expected, r)
		}
	})

	t.Run("Invalid radix", func(t *testing.T) {
		value := 103221
		r := radix(value, -2)

		expected := 0
		if r != expected {
			t.Errorf("Expected %d, got %d", expected, r)
		}
	})
}

func TestRadixSort(t *testing.T) {
	t.Run("Unique int list", func(t *testing.T) {
		list := []int{3, 5, 0, 4, 2, 6, 7, 1, 9, 8}
		list = RadixSort(list)

		expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		if !reflect.DeepEqual(expected, list) {
			t.Errorf("Expected %+v, got %+v", expected, list)
		}
	})

	t.Run("Repeating int list", func(t *testing.T) {
		list := []int{2, 3, 1, 2, 2, 1, 0, 1, 0, 3}
		list = RadixSort(list)

		expected := []int{0, 0, 1, 1, 1, 2, 2, 2, 3, 3}
		if !reflect.DeepEqual(expected, list) {
			t.Errorf("Expected %+v, got %+v", expected, list)
		}
	})
}
