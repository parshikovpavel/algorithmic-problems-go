package _0046_permutations

func permute(nums []int) [][]int {
	// для хранения массива permutations
	result := make([][]int, 0)

	// для хранения текущей permutation
	head := make([]int, 0, len(nums))

	//для флагов, что элемент уже используется в permutation
	used := make([]bool, len(nums))

	// запускаем рекурсию
	recurs(nums, head, used, &result)
	return result
}

func recurs(nums, head []int, used []bool, result *[][]int) {
	// если permutation содержит количество элементов, равное общему количеству
	if len(head) == len(nums) {
		// копируем permutation из head
		permutation := make([]int, len(head))
		copy(permutation, head)

		// вставляем в result
		*result = append(*result, permutation)
		return
	}

	for i, num := range nums {
		// если элемент уже используется в permutation, пропускаем его
		if used[i] {
			continue
		}

		// помечаем, что элемент используется в permutation
		used[i] = true
		// добавляем элемент в permutation
		head = append(head, num)

		// запускаем рекурсию для добавления последующих элементов в permutation
		recurs(nums, head, used, result)

		// убираем элемент из permutation
		head = head[:len(head)-1]
		// помечаем, что теперь он не используется
		used[i] = false
	}
}
