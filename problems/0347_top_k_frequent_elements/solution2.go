package _0347_av_top_k_frequent_elements

import "sort"

// IntHeap – min-heap из int'ов.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int)) // h = [1, 2, 5, 3]
}

func (h *IntHeap) Pop() any {
	old := *h // h = [2, 3, 5, 1]
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1] // h = [2, 3, 5]
	return x
}

func top2(nums []int, k int) []int {
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
