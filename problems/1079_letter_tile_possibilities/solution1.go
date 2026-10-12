package main

func numTilePossibilities(tiles string) int {
	//для флагов, что элемент уже используется в permutation
	used := make([]bool, len(tiles))

	result := make(map[string]struct{})

	recurs(tiles, result, used, "")

	return len(result)
}

func recurs(tiles string, result map[string]struct{}, used []bool, head string) {
	if len(head) == len(tiles) {
		return
	}

	for i := range len(tiles) {
		if used[i] {
			continue
		}
		head += string(tiles[i])
		if _, exist := result[head]; exist {
			head = head[:len(head)-1]
			continue
		}

		result[head] = struct{}{}
		used[i] = true
		recurs(tiles, result, used, head)
		used[i] = false
		head = head[:len(head)-1]
	}
}
