package countvowelstringsinranges

import (
	"reflect"
	"testing"
)

func TestVowelStrings(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		words := []string{"aba", "bcb", "ece", "aa", "e"}
		queries := [][]int{{0, 2}, {1, 4}, {1, 1}}
		want := []int{2, 3, 0}
		got := vowelStrings(words, queries)

		assert(t, got, want)
	})

	t.Run("Example 2", func(t *testing.T) {
		words := []string{"a", "e", "i"}
		queries := [][]int{{0, 2}, {0, 1}, {2, 2}}
		want := []int{3, 2, 1}
		got := vowelStrings(words, queries)
		assert(t, got, want)
	})
}

func assert(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got: %v Want: %v", got, want)
	}
}
