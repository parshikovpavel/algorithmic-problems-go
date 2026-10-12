package _0003_longest_substring_without_repeating_characters

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
