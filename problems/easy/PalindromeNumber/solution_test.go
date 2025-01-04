package palindromenumber

import "testing"

func TestIsPalindrome(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		got := isPalindrome(121)
		want := true
		assert(t, got, want)
	})
	t.Run("Example 2", func(t *testing.T) {
		got := isPalindrome(-121)
		want := false
		assert(t, got, want)
	})
	
	t.Run("Example 3", func(t *testing.T) {
		got := isPalindrome(10)
		want := false
		assert(t, got, want)
	})

}
func assert(t *testing.T, got, want bool) {
	t.Helper()

	if got != want {
		t.Errorf("Got: %t, Want: %t", got, want)
	}
}
