package word_break

// dp[i] indicates if s[0..i] can be segmented
// dp[i] = dp[i-len(word)]&& s[i:i+len(word)] == word . check any word, once it true then dp[i] == true
// boundary: dp[]
// O(n x m), n is len(s) and m is len(wordDict)
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		for _, word := range wordDict {
			if (i + 1 - len(word)) < 0 {
				continue
			}
			if s[:i+1] == word {
				dp[i] = true
				break
			}

			if (i + 1) > len(word) {
				if dp[i-len(word)] && s[i+1-len(word):i+1] == word {
					dp[i] = true
					break
				}
			}
		}
	}

	return dp[len(s)-1]
}
