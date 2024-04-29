package _001_av_two_sum

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_twoSum(t *testing.T) {
	tests := []struct {
		nums []int
		target int
		expected [][2]int
	}{
		{
			[]int{2, 4, 5, 3},
			7,
			[][2]int{[ [2, 5], [4, 3] ]
			```



```
			Input: [-5, 2, 12, 5] 7
			Output: [ [-5, 12], [2, 5] ]


		},


	}
	for _, tt := range tests {
		require.Equal(t, tt.expected, merge(tt.a, tt.b))
	}
}
