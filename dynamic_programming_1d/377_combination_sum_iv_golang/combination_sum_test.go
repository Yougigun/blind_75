package combine_sum

import "testing"

func TestCombinationSum4(t *testing.T) {
	test_cases := []struct {
		title    string
		nums     []int
		target   int
		expected int
	}{
		{
			title:    "case 1",
			nums:     []int{1, 2, 3},
			target:   4,
			expected: 7,
		},
	}

	for _, tc := range test_cases {
		got := combinationSum4(tc.nums, tc.target)
		if got != tc.expected {
			t.Errorf("title: %v, expected: %v,but got: %v", tc.target, tc.expected, got)
		}
	}

}
