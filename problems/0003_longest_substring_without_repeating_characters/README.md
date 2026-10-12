# 3. Longest Substring Without Repeating Characters

## Condition

Given a string, find the length of the **longest substring** without repeating characters.

**Example 1:**

```
Input: "abcabcbb"
Output: 3 
Explanation: The answer is "abc", with the length of 3. 
```

**Example 2:**

```
Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
```

**Example 3:**

```
Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3. 
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
```

## Условие

Дана строка, найти длину самой **длинной подстроки** без повторяющихся символов.

**Пример 1:**

```
Input: "abcabcbb"
Output: 3 
Explanation: Ответ "abc", с длиной 3. 
```

**Пример 2:**

```
Input: "bbbbb"
Output: 1
Explanation: Ответ "b", с длиной 1.
```

**Example 3:**

```
Input: "pwwkew"
Output: 3
Explanation: Ответ "wke", с длиной 3. 
             Заметьте, что ответ должен быть подстрокой, "pwke" - это набор символов, а не подстрока.
```

[Тесты](solution_test.go)

| #    | Время              | Память                        |
| ---- | ------------------ | ----------------------------- |
| 1    | *O(n<sup>3</sup>)* | *O(n)*                        |
| 2    | *O(n+n) = O(n)*    | *O(n)* ( для хранения `hash`) |
| 3    | *O(n)*             | *O(n)*                        |



## Решение 1. Brute force

Организуем полный перебор всех подстрок в строке. Для этого делаем два вложенных цикла: `i = 0 ÷ length-1`, `j = i ÷ length -1`.

Для каждой подстроки проверяем, что все символы в подстроке - уникальны. Используем функцию `areCharsUnique()`.

[Реализация](solution.go):

```php
func LengthOfLongestSubstring(s string) int {
	runes := []rune(s)

	maxLength := 0

	for i := 0; i < len(runes); i++ {
		for j := i + 1; j <= len(runes); j++ {
			if areCharsUnique(runes[i:j]) {
				maxLength = max(maxLength, j-i)
			}
		}
	}

	return maxLength
}

func areCharsUnique(runes []rune) bool {
	hash := make(map[rune]struct{}, len(runes))

	for _, r := range runes {
		_, ok := hash[r]
		if ok {
			return false
		}
		hash[r] = struct{}{}
	}

	return true
}
```



## Решение 2. Sliding window

[Sliding window](https://github.com/parshikovpavel/cheat-sheets/blob/master/Algorithm.md#sliding-window)

Инициализируем два указателя `i` (левая граница *sliding window*) и `j` (правая граница *sliding window*). Инициализируем *Set* `hash`, где будем хранить символы из диапазона `[i; j]`. 

Двигаем указатель `j` вправо на одну позицию и проверяем:

- если символа `s[j+1]` нет в `hash`, то добавляем его в `hash` и увеличиваем максимальную длину `maxLength`
- если символ `s[j+1]` есть в `hash`, то значит где-то раньше в диапазоне  `[i; j]`  этот символ встречался. Начинаем уменьшать размер окна, для этого сдвигаем левую границу `i` вправо на один символ и удаляем этот символ из `hash`. Правую границу `j` на этом шаге не изменяем.

[Реализация](solution2.go):

```go
func LengthOfLongestSubstring2(s string) int {
	runes := []rune(s)

	i := 0
	j := 0
	maxLength := 0
	hash := make(map[rune]struct{}, len(runes))

	for j < len(runes) {
		chJ := runes[j]
		_, ok := hash[chJ]
		if ok { // если символ есть в hash
			chI := runes[i]
			delete(hash, chI) //  # удаляем из hash символ по левой границе
			i++               // двигаем левую границу, уменьшаем размер sliding window
		} else { // Если нет такого символа в hash
			j++                    // двигаем правую границу sliding window
			hash[chJ] = struct{}{} // Добавляем его в hash
		}
		maxLength = max(maxLength, j-i)
	}

	return maxLength
}
```

## Решение 3. Sliding window

[Sliding window](https://github.com/parshikovpavel/cheat-sheets/blob/master/Algorithm.md#sliding-window)

Улучшенный вариант [Решения 2](). 

Аналогично инициализируем два указателя `i` (левая граница *sliding window*) и `j` (правая граница *sliding window*). 

Однако используем *HashTable* `hash`, где будем хранить отображения `['символ_из_строки' => позиция_символа_в_строке]`. Причем в `hash` хранятся как символы, который входят в диапазон `[i; j)`, так и возможно те, которые вышли за пределы этого диапазона.

Теперь, если мы вдруг определяем, что символ `s[j]` ранее нам уже встречался и он входит в диапазон `[i; j)` с индексом `j′`, мы сразу сдвигаем указатель `i` на позицию `j′+1` за этим символом (чтобы исключить символ из диапазона `[i; j]`, при этом пропускаем сразу все символы из диапазона `[i; j′]`) и заносим в `hashTable` новую позицию этого символа `j`. 

Нам не нужно исключать из `hashTable` все остальные символы из диапазона `[i; j′]` после сдвига указателя, т.к. достаточно проверить условие `позиция_символа_в_строке >= i`, чтобы гарантировать, что символ из `hashTable` находится в диапазоне `[i; j]`.

```go
func LengthOfLongestSubstring3(s string) int {
	runes := []rune(s)

	i := 0
	j := 0
	maxLength := 0
	hash := make(map[rune]int, len(runes))

	for j = 0; j < len(runes); j++ { // двигаемся по строке
		chJ := runes[j]
		prev, ok := hash[chJ]
		if ok && prev >= i { // если символ уже ранее нам встречался и находится в диапазоне [i; j)
			i = prev + 1 // сразу сместить указатель i на позицию за этим символом, чтобы исключить его из диапазона [i; j)
		}
		hash[chJ] = j                     // сохранить позицию j текущего символа в hash
		maxLength = max(maxLength, j-i+1) // обновить значение максимальной длины
	}

	return maxLength
}

```

