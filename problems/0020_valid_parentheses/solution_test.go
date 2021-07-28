package _0020_valid_parentheses

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_isValid(t *testing.T) {
	tests := []struct {
		s        string
		expected bool
	}{
		{"()", true},
		{"((", false},
		{"", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"{", false},
		{"]", false},
		{")(", false},
		{"[}{]", false},
	}
	for _, tt := range tests {
		require.Equal(t, tt.expected, isValid(tt.s))
	}
}
