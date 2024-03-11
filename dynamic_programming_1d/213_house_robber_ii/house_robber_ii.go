package house_robber_ii

// [2,3,2]
// 2
// 2 -> 3 => X
// 2 -> 2 => X
// 3
// 3 -> 2 => X
// 2
// 2 -> 3 => X == 3 -> 2 ,
// so we can treat circle as a line of house but with one more condition that beginner and end are now adjacent

// formulation
// f(n): most stolen money starting from nth house
// f1(n) = max(f(n+2),...,f(len(nums)-2)) + nums[n], without last house => this just calculate for f(0). because this dp does not count last house. for other would be the most stolen money without last house
// f2(n + 1) = max(f(n+2),...,f(len(nums)-2), f(len(nums)-1)) + nums[n+1], other index can just be calculated as normal.
// boundary
// f(len(nums)-1) = nums[len(nums)-1]
// f(len(nums)-2) = nums[len(nums)-2]
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}
	// init dp
	dp := make([]int, len(nums))
	// boundary
	dp[len(nums)-1] = 0
	dp[len(nums)-2] = nums[len(nums)-2]

	// build dp with stolen last house
	for i := len(nums) - 1; i >= 0; i-- {
		// formulation
		max := 0
		for j := i + 2; j < len(nums)-1; j++ {

			if max < dp[j] {
				max = dp[j]
			}
		}
		dp[i] = max + nums[i]
	}

	// init dp
	dp2 := make([]int, len(nums))
	// boundary
	dp2[len(nums)-1] = nums[len(nums)-1]
	dp2[len(nums)-2] = nums[len(nums)-2]
	// build dp without stolen last house
	// dp[len(nums)-1] = 0
	for i := len(nums) - 2; i > 0; i-- {
		// formulation
		max := 0
		for j := i + 2; j < len(nums); j++ {

			if max < dp2[j] {
				max = dp2[j]
			}
		}
		dp2[i] = max + nums[i]
	}


	max := 0
	for _, v := range dp2 {
		if max < v {
			max = v
		}
	}

	if max < dp[0] {
		max = dp[0]
	}
	return max
}
