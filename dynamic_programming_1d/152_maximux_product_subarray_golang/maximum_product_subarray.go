package mac_prod_sub_array

import "math"
// dp[i][j] indicates product of arr[i..j]
// dp[i][i+1] == arr[i]
// dp[i][j] == dp[i][j-1] x arr[j]
// dp[i][j] == dp[i-1][j] / arr[i]
// time complexity: O(n^2), n is len(nums)
func maxProduct(nums []int) int {
	max := math.MinInt32
	for i := 0; i < len(nums); i++ {
		temp := nums[i]
		for j := i; j < len(nums); j++ {
			if i != j {
				temp *= nums[j]
			}
			if max < temp {
				max = temp
			}
		}
	}
	return max
}
