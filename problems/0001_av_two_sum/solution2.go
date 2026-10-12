package _0001_av_two_sum

func twoSum2(nums []int, target int) [][2]int {
	hash := make(map[int]int, len(nums))
	var res [][2]int

	for i, num := range nums {
		hash[num] = i
	}

	for i, num := range nums {
		if j, ok := hash[target-num]; ok {
			if i < j {
				res = append(res, [2]int{num, target - num})
			}
		}
	}

	return res
}
