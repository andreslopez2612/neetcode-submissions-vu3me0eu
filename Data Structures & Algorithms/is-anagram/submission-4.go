import (
	"slices"
)


func isAnagram(s string, t string) bool {

	if len(s) == 0 || len(t) == 0 || len(s) != len(t) {
		return false
	}

	v1, v2 := strings.Split(s, ""), strings.Split(t, "")

	slices.Sort(v1)
	slices.Sort(v2)

	/*

	sort.Strings(vs)
	sort.Strings(vt)

	*/

	for v := range v1 {
		if v1[v] != v2[v] {
			return false
		}
	}

	return true
}
