func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	lt1 := sortString(strings.Split(s, ""))
	lt2 := sortString(strings.Split(t, ""))
	for i := 0; i < len(lt1); i++ {
		if lt1[i] != lt2[i] {
			return false
		}
	}
	return true
}

func sortString(s []string) []string {
	//Dividirlas
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	return s
}
