package maximumscoreaftersplittingastring

import "testing"

func TestMaxScore(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		input := "011101"
		want := 5
		got := maxScore(input)
		assert(t, got, want)
	})

	t.Run("Example 2", func(t *testing.T) {
		input := "011101"
		want := 5
		got := maxScore(input)
		assert(t, got, want)
	})

	t.Run("Example 3", func(t *testing.T) {
		input := "1111"
		want := 3
		got := maxScore(input)
		assert(t, got, want)
	})
}

func assert(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
