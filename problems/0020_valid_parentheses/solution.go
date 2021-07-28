package _0020_valid_parentheses

func isValid(s string) bool {
	// Отображение close bracket -> open bracket
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// Стек
	stack := make([]rune, 0)

	for _, r := range s {
		p, ok := pairs[r]

		// Если не нашли по ключу в map, значит open bracket
		if !ok {
			// Кладем ее в stack
			stack = append(stack, r)
			// иначе - это close bracket
			// если stack пуст или элемент на его вершине - не соответствующая open bracket
		} else if len(stack) == 0 || stack[len(stack)-1] != p {
			// возвращаем false
			return false
		} else {
			// удаляем из stack элемент с вершины
			stack = stack[:len(stack)-1]
		}
	}

	// в итоге должны получить пустой stack
	return len(stack) == 0
}
