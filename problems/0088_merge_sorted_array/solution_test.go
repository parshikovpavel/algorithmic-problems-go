package _0088_merge_sorted_array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_isValid(t *testing.T) {
	tests := []struct {
		a        []int
		m        int
		b        []int
		n        int
		expected []int
	}{
		{[]int{}, 0, []int{}, 0, []int{}},
		{[]int{0}, 0, []int{1}, 1, []int{1}},
		{[]int{1}, 1, []int{}, 0, []int{1}},
		{[]int{1, 2}, 2, []int{}, 0, []int{1, 2}},
		{[]int{0, 0}, 0, []int{1, 2}, 2, []int{1, 2}},
		{[]int{1, 0}, 1, []int{2}, 1, []int{1, 2}},
		{[]int{2, 0}, 1, []int{1}, 1, []int{1, 2}},
		{[]int{1, 2, 0}, 2, []int{3}, 1, []int{1, 2, 3}},
		{[]int{3, 0, 0}, 1, []int{1, 2}, 2, []int{1, 2, 3}},
		{[]int{1, 3, 0}, 2, []int{2}, 1, []int{1, 2, 3}},
		{[]int{2, 0, 0}, 1, []int{1, 3}, 2, []int{1, 2, 3}},
		{[]int{1, 0}, 1, []int{1}, 1, []int{1, 1}},
		{[]int{1, 2, 0}, 2, []int{1}, 1, []int{1, 1, 2}},
		{[]int{1, 0, 0}, 1, []int{1, 2}, 2, []int{1, 1, 2}},
	}
	for _, tt := range tests {
		merge(tt.a, tt.m, tt.b, tt.n)
		require.Equal(t, tt.a, tt.expected)
	}
}
