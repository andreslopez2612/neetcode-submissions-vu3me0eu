import (
	"slices"
)
func longestConsecutive(nums []int) int {

	list := make(map[int][]int)

	slices.Sort(nums)

	n := 0

	for i, v := range nums {
		if len(list[n]) == 0 {
			list[n] = append(list[n], v)
			continue
		}
		if v == nums[i-1] {
			continue
		}
		if v-1 == nums[i-1] {
			list[n] = append(list[n], v)
			continue
		} else {
			n++
			list[n] = append(list[n], v)
		}
	}

	result := 0

	for i := range list {
		if len(list[i]) > result {
			result = len(list[i])
		}
	}

	return result
}