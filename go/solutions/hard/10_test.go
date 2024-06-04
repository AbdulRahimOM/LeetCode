package hard

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	tests := []struct {
		s        string
		p        string
		expected bool
	}{
		// {"aa", "a", false},
		{"aa", "a*", true},
		// {"ab", ".*", true},
		// {"aab", "c*a*b", true},
		// {"mississippi", "mis*is*p*.", false},
		// {"mississippi", "mis*is*ip*.", true},
	}

	for _, test := range tests {
		t.Run(test.s+"_"+test.p, func(t *testing.T) {
			result := isMatch(test.s, test.p)
			if result != test.expected {
				t.Errorf("isMatch(%q, %q) = %v; want %v", test.s, test.p, result, test.expected)
			}
		})
	}
}
