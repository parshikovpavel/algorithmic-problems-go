package _0020_valid_parentheses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isValid1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid parenthesis",
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			name: "valid parenthesis",
			args: args{
				s: "((",
			},
			want: false,
		},
		{
			name: "valid parenthesis",
			args: args{
				s: "",
			},
			want: true,
		},
		{
			name: "valid parenthesis",
			args: args{
				s: "(())",
			},
			want: true,
		},
		{
			name: "valid parenthesis",
			args: args{
				s: "(",
			},
			want: false,
		},
		{
			name: "valid parenthesis",
			args: args{
				s: ")",
			},
			want: false,
		},
		{
			name: "valid parenthesis",
			args: args{
				s: "()()",
			},
			want: true,
		},
		{
			name: "valid parenthesis",
			args: args{
				s: ")(",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isValidBraketSeq(tt.args.s))
		})
	}
}
