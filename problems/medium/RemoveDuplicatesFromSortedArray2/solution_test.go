package removeduplicatesfromsortedarray2

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		nums := []int{1, 1, 1, 2, 2, 3}
		wanted := 5
		got := removeDuplicates(nums)
		assert(t, got, wanted, nums[:5], []int{1, 1, 2, 2, 3})
	})

	t.Run("Example 2", func(t *testing.T) {
		nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
		wanted := 7
		fmt.Printf("nums: %v\n", nums)
		got := removeDuplicates(nums)

		assert(t, got, wanted, nums[:7], []int{0, 0, 1, 1, 2, 3, 3})
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
