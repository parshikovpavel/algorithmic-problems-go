package _0002_add_two_numbers_test

import (
	"testing"

	source "github.com/parshikovpavel/algorithmic-problems-go/problems/0002_add_two_numbers"
	"github.com/stretchr/testify/require"
)

func generateList(n int) *source.Node {
	if n == 0 {
		return &source.Node{
			0,
			nil,
		}
	}

	current := &source.Node{}
	head := current

	for n > 0 {
		remainder := n % 10
		node := &source.Node{
			remainder,
			nil,
		}

		current.Next = node

		current = node

		n = n / 10
	}

	return head.Next
}

func Test_addTwoNumbers(t *testing.T) {
	tests := []struct {
		num1 int
		num2 int
		sum  int
	}{
		{10, 210, 220},
		{0, 10, 10},
		{99, 1, 100},
		{1, 1, 2},
		{342, 465, 807},
		{15, 0, 15},
		{0, 0, 0},
	}

	for _, tt := range tests {
		list1 := generateList(tt.num1)
		list2 := generateList(tt.num2)
		res := source.AddTwoNumbers(list1, list2)
		exp := generateList(tt.sum)
		for res != nil || exp != nil {
			if res == nil || exp == nil {
				require.FailNow(t, "Fail")
			}
			require.Equal(t, res.Val, exp.Val)
			res = res.Next
			exp = exp.Next
		}
	}
}
