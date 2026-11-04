package _0746_min_cost_climbing_stairs

func minCostClimbingStairs1(cost []int) int {
	min := -1

	// начинай с первой ступени
	recurs(cost, 0, 0, &min)
	// или со второй ступени
	recurs(cost, 1, 0, &min)

	return min
}

func recurs(cost []int, index int, sum int, min *int) {
	// если добрались до вершины
	if index >= len(cost) {
		// если минимальное значение еще не задали, или сумма меньше минимальной
		if *min == -1 || sum < *min {
			*min = sum
		}
		return
	}

	//поднимаемся еще на одну ступени
	recurs(cost, index+1, sum+cost[index], min)

	// или на две, если не перепрыгнем верщину
	if index+2 <= len(cost) {
		recurs(cost, index+2, sum+cost[index], min)
	}
}
