package _0746_min_cost_climbing_stairs

func minCostClimbingStairs2(cost []int) int {
	// мемоизируем минимальное значение, за которое можем со ступеньки добраться до вершины
	memo := make(map[int]int, len(cost))

	// возвращаем минимальное значение
	return min(
		// за которое можем добраться с первой ступени
		recurs2(cost, 0, memo),
		// или со второй ступени
		recurs2(cost, 1, memo),
	)

}

func recurs2(cost []int, index int, memo map[int]int) int {
	// если добрались до вершины
	if index >= len(cost) {
		// дальше стоимость 0, больше ступеней нет
		return 0
	}

	// если уже мемоизировали минимальную стоимость до вершины, сразу возвращаем ее
	if m, ok := memo[index]; ok {
		return m
	}

	// мемоизируем минимальную стоимость, за которую с текущей степни можем добраться до вершины
	memo[index] = cost[index] + min(
		recurs2(cost, index+1, memo),
		recurs2(cost, index+2, memo),
	)

	// возвращаем эту стоимость
	return memo[index]
}
