package mergestringsalternately

import (
	"testing"
)

func TestMergeStringsAlternately(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		word1, word2 := "abc", "pqr"

		want := "apbqcr"
		got := MergeAlternately(word1, word2)
		assert(t, got, want)
	})

	t.Run("Example 2", func(t *testing.T) {
		word1, word2 := "ab", "pqrs"

		want := "apbqrs"
		got := MergeAlternately(word1, word2)
		assert(t, got, want)
	})

	t.Run("Example 3", func(t *testing.T) {
		word1, word2 := "abcd", "pq"
		want := "apbqcd"

		got := MergeAlternately(word1, word2)
		assert(t, got, want)
	})

}

func assert(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
