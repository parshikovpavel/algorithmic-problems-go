package _0020_valid_parentheses

func isValidBraketSeq(input string) bool {
	brakets := make(map[string]int)

	for _, char := range input {
		switch char {
		case '(':
			brakets["circle"]++
		case ')':
			if brakets["circle"] <= 0 {
				return false
			}

			brakets["circle"]--
		default:
			return false
		}
	}

	return brakets["circle"] == 0
}
