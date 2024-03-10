package coin_change

import "testing"

func TestCoinChange(t *testing.T) {
	test_cases := []struct {
		title    string
		coins    []int
		amount   int
		expected int
	}{
		{
			title:    "case 1",
			coins:    []int{1, 2, 5},
			amount:   11,
			expected: 3,
		},
	}

	for _, tc := range test_cases {
		got := coinChange(tc.coins, tc.amount)
		if tc.expected != got {
			t.Errorf("title: %v, expected: %v, but got: %v", tc.title, tc.expected, got)
		}
	}
}
