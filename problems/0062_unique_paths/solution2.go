package _0062_unique_paths

func uniquePaths2(m int, n int) int {
	// создаем массив количества путей размера m*n
	numPaths := make([][]int, m)
	for i := range m {
		numPaths[i] = make([]int, n)
		// в ячейки первого столбца можно попасть только 1 путем
		numPaths[i][0] = 1
	}
	// в ячейки первой строки можно попасть только 1 путем
	for i := range n {
		numPaths[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// количество путей в ячейку = пути до ячейки сверху + пути до ячейки слева
			numPaths[i][j] = numPaths[i-1][j] + numPaths[i][j-1]
		}
	}

	return numPaths[m-1][n-1]
}
