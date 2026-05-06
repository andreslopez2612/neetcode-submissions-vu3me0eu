func trap(height []int) int {

	result, l, r := 0, 0, len(height)-1
	maxL, maxR := height[0], height[r]
	for l < r {
		if maxL < maxR || maxL == maxR {
			if l > 0 {
				sum := maxL - height[l]
				if sum >= 0 {
					result += sum
				}

			}
			if maxL < height[l] {
				maxL = height[l]
				continue
			}
			l++

		}
		if maxL > maxR {
			if r < len(height)-1 {
				sum := maxR - height[r]
				if sum >= 0 {
					result += sum
				}
			}
			//fmt.Println(height[r])
			if maxR < height[r] {
				maxR = height[r]
				continue
			}
			r--

		}
	}

	return result
}
