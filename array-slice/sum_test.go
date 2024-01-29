package arrayslice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Summing 1-5", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Sum 2 arrays", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		deepCompare(got, want, t)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{9, 0})
		want := []int{2, 0}

		deepCompare(got, want, t)
	})

	t.Run("sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{})
		want := []int{0, 0}

		deepCompare(got, want, t)
	})
}

func deepCompare(got []int, want []int, t *testing.T) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v and want %v", got, want)
	}
}
