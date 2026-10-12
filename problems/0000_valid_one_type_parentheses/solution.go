package _0020_valid_parentheses

func isValidBraketSeq(input string) bool {
	var brackets int

	for _, char := range input {
		switch char {
		case '(':
			brackets++
		case ')':
			brackets--

			if brackets < 0 {
				return false
			}
		default:
			return false
		}
	}

	return brackets == 0
}
