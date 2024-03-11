package house_robber

// example
// Input: nums = [1,2,3,1]
// Output: 4
// Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
// Total amount you can rob = 1 + 3 = 4.

// first solution: brute-force
// using dfs to find all possible combinations and find the maximum.

// second solution
// 1 -> 2 = X
// 1 -> 3 = 4
// 1 -> 3 -> 1 = X
// 1 -> 1 = 2
// 2 -> 3 = X
// 2 -> 1 = 3
// definition of f(n) : that robber can get the most money from n(house n).
// f(n) = max(f(n+2),...,f(len(nums)-1)) + nums[n]
// boundary:
// f((nums)-1) = nums[len(nums)-1]
// 1. init dp with 0 ?
// using bottom-up(tabulation). for-loop
// return the max of dp[n]
func rob(nums []int) int {
	dp := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		max := 0
		for j := i + 2; j <= (len(nums) - 1); j++ {
			if max < dp[j] {
				max = dp[j]
			}
		}
		dp[i] = max + nums[i]
	}

	max := 0
	for _, v := range dp {
		if max < v {
			max = v
		}
	}
	return max
}
