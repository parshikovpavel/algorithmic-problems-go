package _0347_av_top_k_frequent_elements

import "sort"

func top1(nums []int, k int) []int {
	counters := make(map[int]int)

	for _, num := range nums {
		counters[num]++
	}

	uniqueNums := make([]int, 0, len(counters))
	for num := range counters {
		uniqueNums = append(uniqueNums, num)
	}

	sort.Slice(uniqueNums, func(i, j int) bool {
		return counters[uniqueNums[i]] > counters[uniqueNums[j]]
	})

	if k > len(uniqueNums) {
		return uniqueNums
	}

	return uniqueNums[:k]
}
