func topKFrequent(nums []int, k int) []int {

	result := []int{}

	if len(nums) == 0 {
		return result
	}

	check := make(map[int]int)

	for _, v := range nums {
		check[v] += 1
	}

	// 1. Convertir a slice de pares
	type pair struct {
		key   int
		value int
	}

	//Importante
	pairs := make([]pair, 0, len(check))
	for k, v := range check {
		pairs = append(pairs, pair{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].value > pairs[j].value
	})

	for i := range k {
		result = append(result, pairs[i].key)
	}

	return result
}
