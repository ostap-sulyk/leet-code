package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		nums := []int{2, 7, 11, 15}
		target := 9

		want := []int{0, 1}
		got := TwoSum(nums, target)
		assert(t, got, want)
	})

	t.Run("Example 2", func(t *testing.T) {
		nums := []int{3, 2, 4}
		target := 6

		want := []int{1, 2}
		got := TwoSum(nums, target)
		assert(t, got, want)
	})
	t.Run("Example 3", func(t *testing.T) {
		nums := []int{3,3}
		target := 6

		want := []int{0, 1}
		got := TwoSum(nums, target)
		assert(t, got, want)
	})
}

func assert(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
