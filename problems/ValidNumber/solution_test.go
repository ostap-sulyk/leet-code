package validnumber

import (
	"testing"
)

func TestValidNumber(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		intput := "0"
		want := true
		got := IsNumber(intput)

		assert(t, got, want)
	})
	
	t.Run("Example 2", func(t *testing.T) {
		intput := "e"
		want := false
		got := IsNumber(intput)

		assert(t, got, want)
	})
	
	t.Run("Example 3", func(t *testing.T) {
		intput := "."
		want := false
		got := IsNumber(intput)

		assert(t, got, want)
	})
}

func assert(t testing.TB, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("Got %t Want %t", got, want)
	}
}
