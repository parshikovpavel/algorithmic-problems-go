package _0001_av_two_sum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_twoSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected [][2]int
	}{
		{
			[]int{2, 4, 5, 3},
			7,
			[][2]int{{2, 5}, {3, 4}},
		},
		{
			[]int{-5, 2, 12, 5},
			7,
			[][2]int{{-5, 12}, {2, 5}},
		},
		{
			[]int{4},
			8,
			[][2]int{},
		},
	}

	for _, tt := range tests {
		res := twoSum2(tt.nums, tt.target)
		for i, arr := range res {
			res[i] = [2]int{min(arr[0], arr[1]), max(arr[0], arr[1])}
		}

		require.ElementsMatch(t, tt.expected, res)
	}
}
