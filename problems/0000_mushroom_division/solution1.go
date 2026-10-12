package main

func division(a []int) int {
	// заполняем изначально максимальное значение счастья, когда не меняем местами элементы
	// установить в ноль нельзя, максимальное значение счастья может быть отрицательным
	maxHapiness := hapinessCalculate(a, -1, -1)

	// перебираем все пары элементов и меняем их местами, чтобы найти максимальное значение счастья
	for i := range len(a) {
		for j := i + 1; j < len(a); j++ {
			hapiness := hapinessCalculate(a, i, j)
			maxHapiness = max(maxHapiness, hapiness)
		}
	}

	return maxHapiness
}

func hapinessCalculate(a []int, i int, j int) int {
	var hapiness int
	for index, value := range a {
		// если индекс равен i или j, то меняем местами элементы
		switch index {
		case i:
			value = a[j]
		case j:
			value = a[i]
		}
		hapiness += sign(index) * value
	}

	return hapiness
}

func sign(i int) int {
	// возвращаем ПЛЮС, если индекс четный, и МИНУС, если индекс нечетный
	if i%2 == 0 {
		return 1
	}

	return -1
}
