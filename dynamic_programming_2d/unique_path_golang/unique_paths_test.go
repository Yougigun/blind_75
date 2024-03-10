package unique_paths

import "testing"

func TestUniquePaths(t *testing.T) {
	cases := []struct {
		m         int
		n         int
		expected  int
		case_name string
	}{
		{
			case_name: "case 1",
			m:         3,
			n:         7,
			expected:  28,
		},
	}

	for _, c := range cases {
		got := uniquePaths(c.m, c.n)
		if got != c.expected {
			t.Errorf("%v, expected: %v, but got: %v", c.case_name, c.expected, got)
		}
	}
}
