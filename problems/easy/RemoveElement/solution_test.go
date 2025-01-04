package removeelement

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	t.Run("Example 1:", func(t *testing.T) {
		// input
		nums, val := []int{3, 2, 2, 3}, 3
		want := 2
		got := removeElement(nums, val)

		assert(t, got, want, nums[0:2], []int{2, 2})

	})
	t.Run("Example 2:", func(t *testing.T) {
		// input
		nums, val := []int{0,1,2,2,3,0,4,2}, 2
		want := 5
		got := removeElement(nums, val)

		fmt.Printf("nums: %v\n", nums)
		assert(t, got, want, nums[0:5], []int{0,1,3,0,4})
	})

}

func assert(t *testing.T, got, want int, numsOld, numsExpected []int) {
	t.Helper()
	if !reflect.DeepEqual(numsOld, numsExpected) {
		t.Errorf("Got %v but Want: %v", numsOld, numsExpected)
	}
	if got != want {
		t.Errorf("Got %d but Want: %d", got, want)
	}
}
