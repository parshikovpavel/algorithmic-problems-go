package _0046_permutations

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_permute(t *testing.T) {
	tests := []struct {
		nums     []int
		expected [][]int
	}{
		{
			[]int{1, 2, 3},
			[][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			[]int{0, 1},
			[][]int{{0, 1}, {1, 0}},
		},
		{
			[]int{1},
			[][]int{{1}},
		},
	}

	for _, tt := range tests {
		res := permute2(tt.nums)

		require.ElementsMatch(t, tt.expected, res)
	}
}
