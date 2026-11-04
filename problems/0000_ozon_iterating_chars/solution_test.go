package main

import "testing"

func TestRecoverPassword(t *testing.T) {
	for _, exp := range []string{
		"a",
		"12",
		"abc333d",
	} {
		t.Run(exp, func(t *testing.T) {
			act := RecoverPassword2(hashPassword(exp))
			if act != exp {
				t.Error("recovered:", act, "expected:", exp)
			}
		})
	}
}
