func productExceptSelf(nums []int) []int {

	//v := []int{}
	//slice = append(slice[:i], slice[i+1:]...)
	arrMap := make(map[int]int)

	for i := range nums {
		pro := 1
		for j := range nums {

			if i == j {
				continue
			}
			pro *= nums[j]
		}
		arrMap[i] = pro
	}
	length := len(arrMap)

	result := make([]int, length)

	for i, v := range arrMap {
		result[i] = v
	}

	return result
}
