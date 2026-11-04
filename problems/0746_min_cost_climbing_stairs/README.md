# 746. Min Cost Climbing Stairs

## Situation

You are given an integer array `cost` where `cost[i]` is the cost of `ith` step on a staircase.

Once you pay the cost, you can either climb **one** or **two** steps.

You can either start from the step with index 0, or the step with index 1.

Return the **minimum** cost to reach the top of the staircase, which is the position just past the last step (index `cost.length`).

 

**Example 1:**

```
Input: cost = [10,15,20]
Output: 15
Explanation: You will start at index 1.
- Pay 15 and climb two steps to reach the top.
The total cost is 15.
```

**Example 2:**

```
Input: cost = [1,100,1,1,1,100,1,1,100,1]
Output: 6
Explanation: You will start at index 0.
- Pay 1 and climb two steps to reach index 2.
- Pay 1 and climb two steps to reach index 4.
- Pay 1 and climb two steps to reach index 6.
- Pay 1 and climb one step to reach index 7.
- Pay 1 and climb two steps to reach index 9.
- Pay 1 and climb one step to reach the top.
The total cost is 6.
```

 

**Constraints:**

- `2 <= cost.length <= 1000`
- `0 <= cost[i] <= 999`



## Условие задачи

Вам дан массив *integer array* `cost`, где `cost[i]`— стоимость `i-й` ступени на лестнице.

После оплаты вы можете подняться на **одну** или **две** ступеньки.

Вы можете начать либо с шага с индексом 0, либо с шага с индексом 1.

Верните **минимальную** стоимость достижения вершины лестницы, которая находится сразу за последней ступенькой (индекс `cost.length`).

 

**Пример 1:**

```
Input: cost = [10,15,20]
Output: 15
 Пояснение: Вы начнете с индекса 1.
— Заплатите 15 и поднимитесь по двум ступенькам, чтобы добраться до вершины.
Общая стоимость составляет 15.
```

**Пример 2:**

```
Input: cost = [1,100,1,1,1,100,1,1,100,1]
Output: 6
 Пояснение: Вы начнете с индекса 0.
— Заплатите 1 и поднимитесь на две ступеньки, чтобы достичь индекса 2.
— Заплатите 1 и поднимитесь на две ступеньки, чтобы достичь индекса 4.
— Заплатите 1 и поднимитесь на две ступеньки, чтобы достичь индекса 6.
— Заплатите 1 и поднимитесь на одну ступеньку, чтобы достичь индекса 7.
— Заплатите 1 и поднимитесь на две ступеньки, чтобы достичь пункта 9.
— Заплатите 1 и поднимитесь на одну ступеньку, чтобы добраться до вершины.
Общая стоимость составляет 6.
```

 

**Ограничения:**

- `2 <= cost.length <= 1000`
- `0 <= cost[i] <= 999`





| #    | Время                                                        | Память |
| ---- | ------------------------------------------------------------ | ------ |
| 1    | `O(φⁿ)` — **экспоненциальная сложность**, где `φ ≈ 1,618` — золотое сечение. | `O(n)` |
| 2    | `O(n)`                                                       | `O(n)` |
|      |                                                              |        |





## Решение 1.  Brute-force

[Код](solution1.go)

```go
func minCostClimbingStairs(cost []int) int {
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

```



## Решение 2. Мемоизация + расчет минимальной стоимости до вершины

[Код](solution2.go)

```go
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
```







