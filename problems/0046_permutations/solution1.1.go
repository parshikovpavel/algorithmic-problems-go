package _0046_permutations

func permute2(nums []int) [][]int {
	return recurs2(nums, len(nums), nil)
}

// передаем оригинальную длину и текущий массив результатов, возвращаем обновленный результат
func recurs2(nums []int, originalLength int, res [][]int) [][]int {
	// если элементы закончились
	if len(nums) == 1 {
		// добавляем новый результат, его длина - оригинальная длина
		permute := append([]int(nil), nums[:originalLength]...)
		return append(res, permute)
	}

	for i := range len(nums) {
		// переставляем очередной элемент в конец
		nums[i], nums[len(nums)-1] = nums[len(nums)-1], nums[i]
		// подрезаем slice до этой длины
		res = recurs2(nums[:len(nums)-1], originalLength, res)
		// возвращаем элемент назад на место
		nums[i], nums[len(nums)-1] = nums[len(nums)-1], nums[i]
	}
	return res
}
