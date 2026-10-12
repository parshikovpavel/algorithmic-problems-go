package _0003_longest_substring_without_repeating_characters

func LengthOfLongestSubstring1(s string) int {
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
