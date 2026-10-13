package main

func numTilePossibilities(tiles string) int {
	// Заводим array под подсчет количества символов, один элемент - под каждый символ (из 26)
	var charCount [26]int
	// счетчик количества sequence's
	var totalCount int

	// Считаем количества каждого символа в строке
	for _, c := range tiles {
		charCount[c-'A']++
	}

	// Запускаем рекурсию
	findSequences(charCount, &totalCount)

	return totalCount
}

func findSequences(charCount [26]int, totalCount *int) {
	// идем последовательно по каждому символу
	for pos := range len(charCount) {
		// если есть такой доступный для подстановки символ
		if charCount[pos] == 0 {
			continue
		}

		// то считаем, что его добавили и уменьшаем счетчик
		charCount[pos]--
		// получили еще одну sequence, увеличиваем счетчик
		*totalCount++

		findSequences(charCount, totalCount)

		// возвращаем символ назад в массив счетчиков
		charCount[pos]++
	}
}
