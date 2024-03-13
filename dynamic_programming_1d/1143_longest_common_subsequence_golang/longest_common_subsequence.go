package longest_common_subsequence

// https://chat.openai.com/share/938d8cc3-9264-4044-8aa4-0d035cd03561
// Each cell dp[i][j] represents the length of the longest common subsequence 
// between text1[0...i-1] and text2[0...j-1].
// formulation
// dp[i,j] = dp[i-1,j -1]+1 if text1[i] == text2[j]
// dp[i,j] = max(dp[i-1,j],dp[i,j-1]) if text1[i] != text2[j]
// Function to find the length of the longest common subsequence
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	// Create a 2D slice (matrix) with dimensions (m+1) x (n+1)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Fill the matrix
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				// If characters match, add 1 to the result from the previous characters
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// If not, take the maximum value from either excluding the current char from text1 or text2
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// The answer is in dp[m][n]
	return dp[m][n]
}

// Helper function to find the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
