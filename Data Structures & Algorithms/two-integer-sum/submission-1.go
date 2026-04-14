func twoSum(nums []int, target int) []int {
    output := []int{}

	for i,_ := range nums{
		for j := i + 1; j < len(nums); j++{
			if nums[i] + nums[j] == target {
				output = append(output, i, j)
				return output
			}
		}
	}
	return output
}
