package _0003_longest_substring_without_repeating_characters_test

import (
	"testing"

	source "github.com/parshikovpavel/algorithmic-problems-go/problems/0003_longest_substring_without_repeating_characters"
	"github.com/stretchr/testify/require"
)

func Test_addTwoNumbers(t *testing.T) {
	tests := []struct {
		str    string
		length int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"w", 1},
		{"", 0},
		{"deffyu", 3},
		{"ytyyjkl", 4},
		{"atyghakly", 7},
		{"abccade", 4},
	}

	for _, tt := range tests {
		maxLength := source.LengthOfLongestSubstring3(tt.str)

		require.Equal(t, tt.length, maxLength)
	}
}
