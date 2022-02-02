package _0088_merge_sorted_array

func merge(a []int, m int, b []int, n int) {
	aI := m - 1
	bI := n - 1
	resI := m + n - 1

	for aI >= 0 && bI >= 0 {
		if a[aI] > b[bI] {
			a[resI] = a[aI]
			aI--
		} else {
			a[resI] = b[bI]
			bI--
		}
		resI--
	}

	copy(a[:resI+1], b[:bI+1])
}
