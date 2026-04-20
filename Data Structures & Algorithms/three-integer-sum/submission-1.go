func threeSum(nums []int) [][]int {

	result := [][]int{}
	//[-1,0,1,2,-1,-4]
	sort.Ints(nums)
	//[-1,-1,0,1,2,-4]

	for i, v := range nums {
		//Evitar numero duplicados
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		//Recorrer el array por ambos lados
		leftN, rightN := i+1, len(nums)-1

		for leftN < rightN {
			sum := v + nums[leftN] + nums[rightN]

			if sum > 0 {
				rightN -= 1
			} else if sum < 0 {
				leftN += 1
			} else {
				result = append(result, []int{v, nums[leftN], nums[rightN]})
				leftN += 1
				for nums[leftN] == nums[leftN-1] && leftN < rightN {
					leftN += 1
				}
			}
		}

	}

	return result
}