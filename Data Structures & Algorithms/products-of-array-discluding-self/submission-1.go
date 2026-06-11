func productExceptSelf(nums []int) []int {
	result := []int{}

	prefix := 1

	for _, v := range nums {
		result = append(result, prefix)
		prefix = v * prefix
	}

	subfix := 1

	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = result[i] * subfix
		subfix = nums[i] * subfix
	}

	return result
}
