# [20. Valid Parentheses](https://leetcode.com/problems/valid-parentheses/)

## Condition

Given a string `s` containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`, determine if the input string is valid.

An input string is valid if:

1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.

**Example 1:**

```
Input: s = "()"
Output: true
```

**Example 2:**

```
Input: s = "()[]{}"
Output: true
```

**Example 3:**

```
Input: s = "(]"
Output: false
```

**Example 4:**

```
Input: s = "([)]"
Output: false
```

**Example 5:**

```
Input: s = "{[]}"
Output: true
```

**Constraints:**

- `1 <= s.length <= 104`
- `s` consists of parentheses only `'()[]{}'`.

## Условие

Дана строка `s`, содержащая только символы `'('`, `')'`, `'{'`, `'}'`, `'['` и `']'`, определите, является ли входная строка валидной.

Входная строка валидна, если:

1. Открытые скобки должны быть закрыты скобками того же типа.
2. Открытые скобки должны быть закрыты в правильном порядке.

**Пример 1:**

```
Input: s = "()"
Output: true
```

**Пример 2:**

```
Input: s = "()[]{}"
Output: true
```

**Пример 3:**

```
Input: s = "(]"
Output: false
```

**Пример 4:**

```
Input: s = "([)]"
Output: false
```

**Пример 5:**

```
Input: s = "{[]}"
Output: true
```

**Ограничения:**

- `1 <= s.length <= 104`
- `s` состоит только из скобок `'()[]{}'`.

## Решение

```go
func isValid(s string) bool {
	// Отображение close bracket -> open bracket
	pairs := map[rune]rune{
		')' : '(',
		'}' : '{',
		']' : '[',
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
```







