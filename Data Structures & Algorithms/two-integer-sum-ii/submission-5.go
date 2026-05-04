func twoSum(numbers []int, target int) []int {

	for i := range numbers {
		check := target - numbers[i]
		for j := range numbers {
			if numbers[j] == check {
				return []int{i + 1, j + 1}
			}
		}
	}

	return []int{}
}
