# Итерация по символам строки для подбора пароля

## Условие

Есть база данных с хешами паролей пользователей. Пароли хешируются функцией `hashPassword`. Известен набор символов, которые могут использоваться в паролях, — переменная `alphabet`.

Необходимо реализовать функцию `RecoverPassword`, которая восстанавливает пароль по известному хешу, так чтобы тест `TestRecoverPassword` завершался успешно.

```go
package main

import (
	"crypto/md5"
	"testing"
)

var alphabet = []rune{'a', 'b', 'c', 'd', '1', '2', '3'}

func RecoverPassword(h []byte) string {
	// TODO: implement me
	return ""
}

func TestRecoverPassword(t *testing.T) {
	for _, exp := range []string{
		"a",
		"12",
		"abc333d",
	} {
		t.Run(exp, func(t *testing.T) {
			act := RecoverPassword(hashPassword(exp))
			if act != exp {
				t.Error("recovered:", act, "expected:", exp)
			}
		})
	}
}

func hashPassword(in string) []byte {
	h := md5.Sum([]byte(in))
	return h[:]
}
```



Обозначим:

- `k` — количество символов в алфавите, у тебя **7**.
- `L` — длина найденного пароля.

| #    | Время                                                        | Память                                                       |
| ---- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| 1    | `O(L · kᴸ)` - экспоненциальная                               | `O(L · kᴸ)` - экспоненциально, хранит строки текущего и следующего уровней |
| 2    | `O(L · kᴸ)` - экспоненциальная (чуть больше, чем у 1 варианта, т.к. на каждую глубину заново опускаем от корня) | `O(L)` - храним только строку текущей длины, не храним все строки |
|      |                                                              |                                                              |



## Решение 1. Итеративный BFS

[Реализация](solution1.go):

Реализация итеративного BFS по дереву строк по уровням (текущий уровень + следующий уровень) ([тут теория с примером](https://github.com/parshikovpavel/cheat-sheets/blob/master/Algorithm.md#Breadth-first%20search))

```go
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
```



## Решение 2. DFS с итеративным углублением

[Реализация](solution2.go):

Используем DFS с итеративным углублением — *Iterative Deepening Depth-First Search*, **IDDFS**. Получается аналог BFS с последовательным прохождением по уровням.

Ты последовательно увеличиваешь предел глубины `0, 1, 2, …`, каждый раз запускаешь DFS от корня и проверяешь кандидатов на заданной глубине.

Как BFS, алгоритм находит сначала самый короткий подходящий пароль. При этом использует память `O(L)`, как DFS. 

```go
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
```

## Решение 3. Перебор строк в лексикографическом порядке

Идея — **перебирать пароли как числа в системе счисления, где основание равно размеру алфавита**.

Для каждой длины создаём массив индексов символов. Например, для алфавита `a` (`0`), `b` (`1`), `c` (`2`) и длины 2:

```
Индексы    Пароль
[0, 0]     aa
[0, 1]     ab
[0, 2]     ac
[1, 0]     ba
[1, 1]     bb
...
[2, 2]     cc
```

Для каждого кандидата вычисляем хеш и сравниваем с искомым. Если совпал — возвращаем пароль.

Чтобы получить следующий кандидат, увеличиваем последний индекс. Если он вышел за пределы алфавита, обнуляем его и переносим единицу в предыдущую позицию — как при переходе от `99` к `100`.

Когда все варианты текущей длины закончились, переходим к следующей длине. Так можно перебрать все строки, включая повторения букв, **без рекурсии и хранения всех кандидатов в памяти**.
