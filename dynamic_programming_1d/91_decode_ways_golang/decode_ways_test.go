package decode_ways

import "testing"

func TestNumDecodings(t *testing.T) {
	got := numDecodings("12")
	expected := 2
	if got != expected {
		t.Errorf("case 1 - expected: %v, but got: %v", expected, got)
	}
}
