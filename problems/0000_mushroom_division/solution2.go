package main

func division(a []int) int {
	// текущее значение радости Васи
	happiness := 0
	// минимальный гриб Васи
	minVasyaIdx := 0
	// максимальный гриб Маши
	maxMashaIdx := 1

	// проходимся по всем грибам
	// считаем (1) текущее значение радости
	// (2) минимальный гриб Васи (3) максимальный гриб Маши
	for i, val := range a {
		if i%2 == 0 {
			// это гриб Васи
			happiness += val
			if a[minVasyaIdx] > val {
				minVasyaIdx = i
			}
			continue
		}

		// это гриб Маши
		happiness -= val
		if a[maxMashaIdx] < val {
			maxMashaIdx = i
		}
	}

	// если грибов меньше 2, то поменять нечего (по условию, такого не должно быть)
	if len(a) < 2 {
		return happiness
	}

	// если минимальный гриб Васи - меньше максимального гриба Маши
	if a[minVasyaIdx] < a[maxMashaIdx] {
		// меняем их местами
		// итоговое счастье - исходное счастье плюс 2 * максимальный гриб Маши минус 2 * минимальный гриб Васи
		return happiness + 2*a[maxMashaIdx] - 2*a[minVasyaIdx]
	}

	return happiness
}
