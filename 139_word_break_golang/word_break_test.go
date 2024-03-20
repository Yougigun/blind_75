package word_break

import "testing"

func TestWordBreak(t *testing.T) {
	s := "a"
	wordDict := []string{"b"}
	expected := true
	got := wordBreak(s, wordDict)
	if got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}
}
