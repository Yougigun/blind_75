package palindromic_substring

import "testing"

// Test case 1
func TestLongestPalindromicSubstring1(t *testing.T) {
	result := longestPalindrome("babad")
	if result != "aba" && result != "bab" {
		t.Errorf("Got: %s, want: aba or bab", result)
	}
}
