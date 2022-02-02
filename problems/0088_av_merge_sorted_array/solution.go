package _0088_av_merge_sorted_array

func merge(a []int, b []int) []int {
	resLength := len(a) + len(b)

	res := make([]int, resLength)

	aI := 0
	bI := 0
	resI := 0

	for aI < len(a) && bI < len(b) {
		if a[aI] < b[bI] {
			res[resI] = a[aI]
			aI++
		} else {
			res[resI] = b[bI]
			bI++
		}
		resI++
	}

	copy(res[resI:], a[aI:])
	copy(res[resI:], b[bI:])

	return res
}
