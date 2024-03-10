package longest_incr_subsequence

// longest increasing subsequence must start from one certain index
// so if we can find all the longest subsequence for each index then we can solve the problem.
//
// [10,9,2,5,3,7,101,18]
// f(index) is the longest length
// f(index) = max(f(index+1), f(index+2), ..., f(index + ..)) + 1, for nums[index] < nums[index+ a], for a from 1 to len(num)-1
// boundary:
// f(len(nums)) = 1
// f(x) if x out of bound the it is invalid
// using bottom-top dp from dp[len(nums)-1]
// time complexity: O(m^2)
func lengthOfLIS(nums []int) int {
	dp := make([]int, 0, len(nums)-1)
	for i := 0; i < len(nums); i++ {
		dp = append(dp, 1)
	}
	for i := len(nums) - 1; i >= 0; i-- {
		// max can not be 1. because the f(index) could be all invalid.
		max := 0
		for j := i + 1; j < len(nums); j++ {
			if nums[i] >= nums[j] {
				continue
			}
			if max < dp[j] {
				max = dp[j]
			}
		}
		dp[i] = max + 1
	}
	max := 0
	for _, v := range dp {
		if max < v {
			max = v
		}
	}
	return max
}

func lengthOfLISOfDFS(nums []int) int {
	dp := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		dp = append(dp, -1)
	}
	dp[len(nums)-1] = 1
	var dfs func(index int) int
	dfs = func(index int) int {
		if dp[index] != -1 {
			return dp[index]
		}

		max := 0
		for i, v := range nums[index+1:] {
			i = index + 1 + i
			if nums[index] < v {
				r := dfs(i)
				if max < r {
					max = r
				}
			}
		}
		dp[index] = max + 1
		return max + 1
	}
	for i := range nums {
		dp[i] = dfs(i)
	}
	max := 0
	for _, v := range dp {
		if max < v {
			max = v
		}
	}
	return max
}
