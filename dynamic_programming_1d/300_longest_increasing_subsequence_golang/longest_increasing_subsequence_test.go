package longest_incr_subsequence

import "testing"

func TestLenOfLIS(t *testing.T) {

	test_cases := []struct {
		title    string
		nums     []int
		expected int
	}{
		{
			title:    "case 1",
			nums:     []int{10, 9, 2, 5, 3, 7, 101, 18},
			expected: 4,
		},
	}

	for _, tc := range test_cases {
		got := lengthOfLISOfDFS(tc.nums)
		if got != tc.expected {
			t.Errorf("title: %v, expected: %v, but got: %v", tc.title, tc.expected, got)
		}
	}
}
