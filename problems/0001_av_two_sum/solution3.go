package _0001_av_two_sum

func twoSum3(nums []int, target int) [][2]int {
	hash := make(map[int]struct{}, len(nums))
	var res [][2]int

	for _, num := range nums {
		if _, ok := hash[target-num]; ok {
			res = append(res, [2]int{num, target - num})
		}

		hash[num] = struct{}{}
	}

	return res
}
