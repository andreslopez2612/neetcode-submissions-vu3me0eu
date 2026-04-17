func groupAnagrams(strs []string) [][]string {
	check := make(map[string][]string)

	for _, v := range strs {
		sorted := sortedS(v)
		check[sorted] = append(check[sorted], v)
	}

	var result [][]string

	for _, g := range check {
		result = append(result, g)
	}

	return result
}

func sortedS(s string) string {

	p := strings.Split(s, "")

	sort.Slice(p, func(i, j int) bool {
		return p[i] < p[j]
	})

	c := strings.Join(p, "")

	return c
}