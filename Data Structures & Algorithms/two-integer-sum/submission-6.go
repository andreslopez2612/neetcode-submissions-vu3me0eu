func twoSum(nums []int, target int) []int {

	result := []int{}

	if len(nums) == 0 {
		return result
	}

	check := make(map[int]int)
	for i := range nums {
		diference := target - nums[i]
		if _, ok := check[diference]; ok {
			result = append(result, check[diference], i)
			return result
		}

		check[nums[i]] = i
	}

	return result
}
