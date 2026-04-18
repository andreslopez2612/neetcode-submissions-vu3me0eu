func topKFrequent(nums []int, k int) []int {
	//Map for add the frequent number
	count := make(map[int]int)

	//Adding the frequent number
	for _, v := range nums {
		count[v] = count[v] + 1
	}

	//Created map for organize the frequent numbers
	//Started with len 0 for prevent the error
	newMap := make([][2]int, 0)

	for i, v := range count {
		newMap = append(newMap, [2]int{v, i})
	}

	sort.Slice(newMap, func(i, j int) bool {
		return newMap[i][0] > newMap[j][0]
	})

	//Created map for return the result
	result := []int{}
	for i := range k {
		result = append(result, newMap[i][1])
	}

	return result
}

