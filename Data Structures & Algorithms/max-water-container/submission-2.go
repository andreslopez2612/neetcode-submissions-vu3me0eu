func maxArea(heights []int) int {

	c, l, r := 0, 0, len(heights)-1
	for l < r {
		val := min(heights[l], heights[r])
		total := (r - l) * val
		if total > c {
			c = total
		}
		if val == heights[l] {
			l++
		} else {
			r--
		}

	}

	return c
}
