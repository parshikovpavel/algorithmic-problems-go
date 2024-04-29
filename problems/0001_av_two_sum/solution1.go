package _001_av_two_sum

func twoSum1(nums []int, target int) (result [][2]int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); i++ {
			if nums[i]+nums[j] == target {
				result = append(result, [2]int{nums[i], nums[j]})
			}
		}
	}

	return
}
