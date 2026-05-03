func longestConsecutive(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)
	fmt.Println(nums)

	count := 1
	check := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}

		if nums[i] == nums[i-1]+1 {
			check++
		} else {
			check = 1
		}

		if check > count {
			count = check
		}
	}

	return count
}
