func groupAnagrams(strs []string) [][]string {
	group := make(map[string][]string)

	for _, v := range strs {
		stringHead := order(v)
		if _, ok := group[stringHead]; ok {
			group[stringHead] = append(group[stringHead], v)
			continue
		}
		group[stringHead] = append(group[stringHead], v)

	}

	result := [][]string{}

	for _, v := range group {
		result = append(result, v)
	}

	return result
}

func order(s string) string {
	res := strings.Split(s, "")
	//slices.Sort(res)
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	return strings.Join(res, "")
}
