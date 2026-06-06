func topKFrequent(nums []int, k int) []int {
	group := make(map[int]int)

	for _, v := range nums {
		group[v] += 1
	}

	type list struct {
		key   int
		value int
	}
	resultMap := make([]list, 0, len(group))
	for i, v := range group {
		resultMap = append(resultMap, list{i, v})
	}

	sort.Slice(resultMap, func(i, j int) bool {
		return resultMap[i].value > resultMap[j].value
	})
	result := []int{}

	for i := range k {
		result = append(result, resultMap[i].key)
	}

	return result
}