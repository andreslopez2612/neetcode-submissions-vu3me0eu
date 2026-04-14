func isPalindrome(s string) bool {

	str := []string{}

	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			str = append(str, strings.ToLower(string(v)))
		}
	}

	for j := 0; j < len(str); j++ {
		num := (len(str) - 1) - j
		if str[j] == str[num] {
			continue
		}
		return false
	}

	return true
}
