package main

import "fmt"

func main() {
	const (
		n = 5
		k = 3
	)

	combination := make([]int, k)
	var result [][]int

	recurs(n, k, 0, 0, combination, &result)

	fmt.Println(result)
}

func recurs(n int, k int, currentK int, prev int, combination []int, result *[][]int) {
	// base case – текущий размер Combination `currentK` достиг требуемого размера `k`.
	if currentK == k {
		res := make([]int, k)
		copy(res, combination)
		*result = append(*result, res)
		return
	}

	// количество элементов, которые нам еще нужно добавить
	need := k - currentK
	// осталось всего чисел для перебора
	remain := n - prev
	// имеет смысл еще перебирать столько элементов на данной итерации
	available := remain - need

	// recursive case.
	// Перебираем в цикле числа из set'а, начиная с позиции `prev` – позиция после числа, взятого на
	// предыдущем шаге рекурсии (в этом случае не возможны повторения *combination*'s).
	// Перебор продолжаем, пока не переберем все числа до `n`, либо не останется в set достаточно чисел,
	// чтобы получить Combination размера `k`.
	for i := prev; i <= prev+available; i++ {
		// Добавляем очередное число `i` в текущую комбинацию `combination`,
		combination[currentK] = i + 1
		// вызываем функцию рекурсивно, при этом стартовая позиция для очередного перебора  `i + 1`.
		recurs(n, k, currentK+1, i+1, combination, result)
	}
}
