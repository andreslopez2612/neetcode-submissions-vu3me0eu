func groupAnagrams(strs []string) [][]string {

	check := make(map[string][]string)

	for _, v := range strs {
		check[sortStr(v)] = append(check[sortStr(v)], v)
	}

	result := [][]string{}

	for _, v := range check {
		result = append(result, v)
	}

	return result
}


func sortStr(s string) string {

	val := strings.Split(s, "")

	sort.Slice(val, func(i, j int) bool {
		return val[i] < val[j]
	})

	//slices.Sort(val)
	return strings.Join(val, "")
}
