package main

import (
	"bytes"
	"crypto/md5"
)

func RecoverPassword1(h []byte) string {
	// в цикле перебираем глубины, на которую будем опускать в DFS на текущей итерации
	for length := 0; ; length++ {
		// сразу создаем строку нужно длины
		currentString := make([]rune, length)
		// запускаем рекурсию
		if recurs(h, 0, length, currentString) {
			// если рекурсия вернула true, строка найдена
			return string(currentString)
		}
	}
}

// length - текущая длина строки (глубина DFS)
// maxLength - максимальная глубина, на которую опускаемся на текущей итерации
func recurs(h []byte, length int, maxLength int, currentString []rune) bool {
	// если составили строку требуемой длины, проверяем ее хеш
	if length == maxLength {
		if bytes.Equal(h, hashPassword1(string(currentString))) {
			return true
		}

		return false
	}

	// в цикле перебираем все ветви DFS на текущем уровне
	// присоединяем к строке поочереди все символы
	for _, a := range alphabet {
		currentString[length] = a

		// запускаем рекурсию дальше, добавляем следующие символы в строку
		if recurs(h, length+1, maxLength, currentString) {
			return true
		}
	}

	return false
}

func hashPassword1(in string) []byte {
	h := md5.Sum([]byte(in))
	return h[:]
}
