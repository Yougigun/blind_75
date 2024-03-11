package house_robber_ii

import "testing"

func TestRobHouseII(t *testing.T) {
	got := rob([]int{2, 3, 2})
	expected := 3

	if got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}
}
