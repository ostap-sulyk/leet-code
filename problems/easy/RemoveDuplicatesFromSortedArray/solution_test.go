package removeduplicatesfromsortedarray

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		nums := []int{1, 1, 2}
		want := 2
		got := removeDuplicates(nums)
		assert(t, got, want, nums[:2], []int{1, 2})
	})
	
	t.Run("Example 1", func(t *testing.T) {
		nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
		want := 5
		got := removeDuplicates(nums)
		assert(t, got, want, nums[:5], []int{0, 1, 2, 3, 4})
	})
}

func assert(t *testing.T, got, want int, originalNums, expectedNums []int) {
	t.Helper()
	if got != want {
		t.Errorf("Got: %d Want: %d", got, want)
	}
	if !reflect.DeepEqual(originalNums, expectedNums) {
		t.Errorf("Original Numbers: %v Expected: %v", originalNums, expectedNums)
	}
}
