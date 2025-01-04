package mergesortedarray

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		nums1 := []int{1, 2, 3, 0, 0, 0}
		nums2 := []int{2, 5, 6}
		m := 3
		n := 3
		merge(nums1, m, nums2, n)

		want := []int{1, 2, 2, 3, 5, 6}
		got := nums1
		assert(t, &got, &want)
	})

	t.Run("Example 2", func(t *testing.T) {
		nums1 := []int{1}
		nums2 := []int{}
		m := 1
		n := 0

		merge(nums1, m, nums2, n)

		want := []int{1}
		got := nums1
		assert(t, &got, &want)
	})

	t.Run("Example 3", func(t *testing.T) {
		nums1 := []int{0}
		m := 0
		nums2 := []int{1}
		n := 1
		
		merge(nums1, m, nums2, n)

		want := []int{1}
		got := nums1
		assert(t, &got, &want)
	})
}

func assert(t *testing.T, got, want *[]int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v Want %v", got, want)
	}
}
