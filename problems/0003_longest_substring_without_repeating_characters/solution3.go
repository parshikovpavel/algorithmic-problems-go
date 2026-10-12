package _0003_longest_substring_without_repeating_characters

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
