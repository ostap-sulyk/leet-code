package majorityelement

import (
	"testing"
)

func TestMajorityElement(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		nums := []int{3, 2, 3}
		want := 3
		got := majorityElement(nums)

		assert(t, got, want)
	})
	t.Run("Example 1", func(t *testing.T) {
		nums := []int{2, 2, 1, 1, 1, 2, 2}
		want := 2
		got := majorityElement(nums)

		assert(t, got, want)

	})
}

func assert(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("Got: %d, Want: %d", got, want)
	}
}
