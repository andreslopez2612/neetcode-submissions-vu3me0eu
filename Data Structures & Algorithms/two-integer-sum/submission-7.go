func twoSum(nums []int, target int) []int {

	result := []int{}

	check := make(map[int]int)
	for i, v := range nums {
		diference := target - v
		if prev, ok := check[diference]; ok {
			result = append(result, prev, i)
			return result
		}
		check[v] = i
	}

	return nil
}