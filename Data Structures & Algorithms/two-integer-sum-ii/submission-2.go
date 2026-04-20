func twoSum(numbers []int, target int) []int {

	result := []int{}

	for i := range numbers {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				result = append(result, i+1, j+1)
				return result
			}
		}
	}
	return result
}

