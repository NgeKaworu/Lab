package removeduplicateletters

import "testing"

func TestRemoveDuplicateLetters(t *testing.T) {
	tests := []struct {
		s, sol string
	}{
		{"bcabc", "abc"},
		{"cbacdcbc", "acdb"},
	}

	for _, v := range tests {
		actual := removeDuplicateLetters(v.s)
		if actual != v.sol {
			t.Errorf("removeDuplicateLetters(%v) = %v; expected %v", v.s, actual, v.sol)
		}
	}
}
