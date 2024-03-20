package mac_prod_sub_array

import "testing"

func TestMaxProdSubArray(t *testing.T) {
	nums := []int{2, 3, -2, 4}
	expected := 6
	got := maxProduct(nums)
	if got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}
}
