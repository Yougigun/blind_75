package combine_sum

// approach:
// f(n): the number of possible combinations that up to target
// formulation:
// f(n) = sum of f(n-nums[i]) for 0 <= i < len(nums)
// boundary:
// f(0) = 1
// using bottom-up(for-loop to build dp. i.e. f(n) )
// using top-down(dfs and memory to reduce duplicate computation)
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		sum := 0
		for _, v := range nums {
			diff := i - v
			if diff < 0 || dp[diff] == 0 {
				continue
			}
			sum = dp[diff] + sum

		}
		dp[i] = sum
	}

	return dp[target]
}
