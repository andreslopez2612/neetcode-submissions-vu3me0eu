func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	for i := range nums {
		target := nums[i]
		if target > 0 {
			break
		}
		if i > 0 && target == nums[i-1] {
			continue
		}
		indexJ, indexK := i+1, len(nums)-1
		for indexJ < indexK {

			sum := target + nums[indexJ] + nums[indexK]

			if sum > 0 {
				indexK--
				continue
			} else if sum < 0 {
				indexJ++
			} else {
				result = append(result, []int{nums[i], nums[indexJ], nums[indexK]})
				indexJ++
				indexK--
				for indexJ < indexK && nums[indexJ] == nums[indexJ-1] {
					indexJ++
				}
			}

		}
	}

	return result
}