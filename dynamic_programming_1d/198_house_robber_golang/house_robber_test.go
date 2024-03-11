package house_robber

import "testing"

func TestRob(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	expected := 4
	got := rob(nums)
	if got!=expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}
}
