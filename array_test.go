package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("fixed size [slices]", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}

		got := Sum(nums)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, nums)
		}
	})

	t.Run("array of any size", func(t *testing.T) {
		nums := []int{1, 2, 3}

		got := Sum(nums)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, nums)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
