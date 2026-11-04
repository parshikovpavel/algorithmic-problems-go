package _0347_av_top_k_frequent_elements

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTop1(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1, 1, 1, 2, 2, 3}, 0, []int{}},
		{[]int{1, 1, 1, 2, 2, 3}, 4, []int{1, 2, 3}},
		{[]int{1}, 2, []int{1}},
	}
	for _, tt := range tests {
		actual := top1(tt.nums, tt.k)
		require.ElementsMatch(t, tt.expected, actual)
	}
}
