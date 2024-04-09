package palindromic_substring

// Initialize a 2D boolean array dp where dp[i][j] is true if the substring s[i..j]
// is a palindrome, and false otherwise.

// Every single character is a palindrome by itself, so we set dp[i][i] = true for all i.

// A two-character substring s[i..i+1] is a palindrome if s[i] == s[i+1], so we
// check and set dp[i][i+1] = true if s[i] == s[i+1].

// For substrings longer than 2 characters, we use the
// formula: dp[i][j] = dp[i+1][j-1] && s[i] == s[j].
// This means a substring s[i..j] is a palindrome if s[i] == s[j] and the
// substring s[i+1..j-1] is also a palindrome.
// Note that we need to check all substrings of length 3, 4, 5, ..., n.
// Because we have the two-character substrings already covered then we can calculate the longer substrings.
// 2->3->4->5->...->n

// We keep track of the start and length (or end) of the longest palindromic substring found so far.
// When we find a longer palindrome, we update these variables.

// Finally, we return the longest palindromic substring using the start and length/end values.

// Analysis
// time complexity: O(n^2)
// space complexity: O(n^2)

func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	start := 0
	maxLength := 1

	// All substrings of length 1 are palindromes
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	// Check for substrings of length 2
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			start = i
			maxLength = 2
		}
	}

	// Check for lengths greater than 2
	for length := 3; length <= n; length++ {
		for i := 0; i < n-length+1; i++ {
			// Get the ending index of substring from starting index i and length
			j := i + length - 1
			// Checking for sub-string from ith index to jth index if s[i+1] to s[j-1] is a palindrome
			if dp[i+1][j-1] && s[i] == s[j] {
				dp[i][j] = true
				if length > maxLength {
					start = i
					maxLength = length
				}
			}
		}
	}

	return s[start : start+maxLength]
}
