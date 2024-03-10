package coin_change

import (
	"math"
)

// f(n) is minium nums of coins
// f(11) = min(f(10), f(9), f(6))
// f(n) = min(f(n-1) , f(n-2)  , f(n-5) ) + 1
// boundary:
// f(0) = 0
// f(x for any x <0 ), invalid
// using dynamic programming. either bottom-top(for loop) or top-bottom(dfs)
// time complexity: O(n x amount)

func coinChange(coins []int, amount int) int {
	dp := make([]int, 0, amount+1)
	for i := 0; i <= amount; i++ {
		dp = append(dp, -1)
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		min := math.MaxInt32
		for _, c := range coins {
			diff := i - c
			if diff < 0 {
				continue
			}
			if dp[diff] == -1 {
				continue
			}
			if min > (dp[diff] + 1) {
				min = dp[diff] + 1
			}
		}
		if min != math.MaxInt32 {
			dp[i] = min
		}
	}
	return dp[amount]
}
