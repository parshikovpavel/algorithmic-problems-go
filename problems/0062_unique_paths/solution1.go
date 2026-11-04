package _0062_unique_paths

func uniquePaths1(m int, n int) int {
	memo := make([][]int, m)
	for i := range m {
		memo[i] = make([]int, n)
	}
	return recursion(m-1, n-1, memo)
}

func recursion(i int, j int, memo [][]int) int {
	// проверяем - сохранено ли уже значение для клетки
	if memo[i][j] != 0 {
		return memo[i][j]
	}

	// для первой строки и первого столбца - только один способ попасть в них
	if i == 0 || j == 0 {
		memo[i][j] = 1
	} else {
		// для внутренних клеток - количество путей перемещения - это сумма путей в клетку слева и сверху
		memo[i][j] = recursion(i-1, j, memo) + recursion(i, j-1, memo)
	}

	return memo[i][j]
}
