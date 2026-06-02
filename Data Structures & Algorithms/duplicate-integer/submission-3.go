func hasDuplicate(nums []int) bool {
	check := make(map[int]bool)

	for _, v := range nums {
		if check[v] {
			return true
		}

		check[v] = true
	}

	return false
}
