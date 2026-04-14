func hasDuplicate(nums []int) bool {
    duplicate := make(map[int]bool)

	for _, v := range nums {
		if duplicate[v] {
			return true
		}
		duplicate[v] = true
	}

	return false
}
