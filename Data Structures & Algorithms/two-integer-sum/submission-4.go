func twoSum(nums []int, target int) []int {
	result := []int{}

	if len(nums) == 0 {
		return append(result, 0, 0)
	}

	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j > 0; j-- {
			if nums[i]+nums[j] == target && i < j{
				return append(result, i, j)
			}
		}
	}

	return result
}
