func maxArea(heights []int) int {
	maxNumber := 0

	l, r := 0, len(heights)-1

	for l < r {

		if heights[l] < heights[r] {
			cal := (r - l) * heights[l]
			if cal > maxNumber {
				maxNumber = cal
			}
			l++
		} else {
			cal := (r - l) * heights[r]
			if cal > maxNumber {
				maxNumber = cal
			}
			r--
		}
	}

	return maxNumber
}