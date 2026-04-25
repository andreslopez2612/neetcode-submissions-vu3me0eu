func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	vs, vt := strings.Split(s, ""), strings.Split(t, "")

	sort.Strings(vs)
	sort.Strings(vt)

	for i := 0; i < len(s); i++ {
		if vs[i] != vt[i] {
			return false
		}
	}

	return true
}
