package main

import (
	"bytes"
	"crypto/md5"
)

var alphabet = []rune{'a', 'b', 'c', 'd', '1', '2', '3'}

func RecoverPassword(h []byte) string {
	if bytes.Equal(h, hashPassword("")) {
		return ""
	}

	currentLengthStrings := []string{""}

	for {
		nextLengthStrings := make([]string, 0, len(currentLengthStrings)*len(alphabet))

		for _, s := range currentLengthStrings {
			for _, a := range alphabet {
				str := s + string(a)

				if bytes.Equal(h, hashPassword(str)) {
					return str
				}
				nextLengthStrings = append(nextLengthStrings, str)
			}
		}

		currentLengthStrings = nextLengthStrings
	}

}

func hashPassword(in string) []byte {
	h := md5.Sum([]byte(in))
	return h[:]
}
