package unique_paths

// Approach
// using depth-first-search to go through all possible route
// there must be many routes the robot has already taken.
// so we need a dp to remember each cells's valid number of routes to corner
// using 2d array instead of map to reduce the computation time.
// formula : f(pos_r, pos_c)= f(pos_r+1, pos_c) + f(pos_r, pos_c + 1)
// boundary condition:
// f(m-1, n-1) = 1
// f(our of matrix) = 0
// using post-order to sum up the all the ways for current pos
//
// time complexity : O(mxn)
// space complexity : O(mxn)
func uniquePaths(m int, n int) int {
	dp := [][]int{}

	for i := 0; i < m; i++ {
		r := []int{}
		for j := 0; j < n; j++ {
			r = append(r, 0)
		}
		dp = append(dp, r)
	}

	dp[m-1][n-1] = 1
	var dfs func(pos_r int, pos_c int) int
	dfs = func(pos_r, pos_c int) int {
		if pos_r >= m || pos_r < 0 || pos_c >= n || pos_c < 0 {
			return 0
		}

		if dp[pos_r][pos_c] != 0 {
			return dp[pos_r][pos_c]
		}
		dp[pos_r][pos_c] = dfs(pos_r, pos_c+1) + dfs(pos_r+1, pos_c)
		return dp[pos_r][pos_c]
	}

	return dfs(0, 0)
}
