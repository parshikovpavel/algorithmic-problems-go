package _0088_av_merge_sorted_array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_isValid(t *testing.T) {
	tests := []struct {
		a        []int
		b        []int
		expected []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{}, []int{1}, []int{1}},
		{[]int{1}, []int{}, []int{1}},
		{[]int{1, 2}, []int{}, []int{1, 2}},
		{[]int{}, []int{1, 2}, []int{1, 2}},
		{[]int{1}, []int{2}, []int{1, 2}},
		{[]int{2}, []int{1}, []int{1, 2}},
		{[]int{1, 2}, []int{3}, []int{1, 2, 3}},
		{[]int{3}, []int{1, 2}, []int{1, 2, 3}},
		{[]int{1, 3}, []int{2}, []int{1, 2, 3}},
		{[]int{2}, []int{1, 3}, []int{1, 2, 3}},
		{[]int{1}, []int{1}, []int{1, 1}},
		{[]int{1, 2}, []int{1}, []int{1, 1, 2}},
		{[]int{1}, []int{1, 2}, []int{1, 1, 2}},
	}
	for _, tt := range tests {
		require.Equal(t, tt.expected, merge(tt.a, tt.b))
	}
}
