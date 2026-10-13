package main

func numTilePossibilities(tiles string) int {
	//для отслеживания уже использованных символов
	used := make([]bool, len(tiles))

	// для хранения уникальных последовательносткй
	result := make(map[string]struct{})

	// запускаем рекурсию, вначале - пустая строка в последовательности
	recurs(tiles, result, used, "")

	return len(result)
}

func recurs(tiles string, result map[string]struct{}, used []bool, head string) {
	// когда получили последовательность с длиной равной количеству символов,
	// делаем возврат - больше присоединять нечего
	if len(head) == len(tiles) {
		return
	}

	for i := range len(tiles) {
		// если символ использован - пропускаем его
		if used[i] {
			continue
		}
		// присоединяем символ к текущей последовательности
		head += string(tiles[i])
		if _, exist := result[head]; exist {
			// если такая последовательность уже была сохранена ранее
			// убираем символ из последовательности и пропускаем эту ветку
			head = head[:len(head)-1]
			continue
		}

		// если такой последовательности не было - сохраняем ее
		result[head] = struct{}{}
		// помечаем символ - как используемый
		used[i] = true
		// исследуем дальше эту ветку
		recurs(tiles, result, used, head)
		// возвращаем все назад - (1) символ не используется (2) убираем из последовательности
		used[i] = false
		head = head[:len(head)-1]
	}
}
