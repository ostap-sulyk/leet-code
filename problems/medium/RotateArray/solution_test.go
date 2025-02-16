package rotatearray

import (
	"reflect"
	"testing"
)

func TestRotateArray(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6, 7}
		k := 3
		output := []int{5, 6, 7, 1, 2, 3, 4}

		rotate(nums, k)
		assert(t, nums, output)
	})
	
	t.Run("Example 2", func(t *testing.T) {
		nums := []int{-1, -100, 3, 99}
		k := 2
		output := []int{3, 99, -1, -100}
		rotate(nums, k)
		assert(t, nums, output)
	})

}

func assert(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got: %v, want: %v", got, want)
	}
}
