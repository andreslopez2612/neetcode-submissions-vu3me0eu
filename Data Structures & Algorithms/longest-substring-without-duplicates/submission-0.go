func lengthOfLongestSubstring(s string) int {
	check := make(map[byte]bool)

	l := 0
	result := 0

	for r := 0; r < len(s); r++ {

		for check[s[r]] {
			delete(check, s[l])
			l++
		}

		check[s[r]] = true

		if r-l+1 > result {
			result = r - l + 1
		}
	}

	return result
}
