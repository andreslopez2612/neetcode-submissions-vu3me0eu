func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	str := ""

	for _, c := range s {
		if ("a" <= string(c) && string(c) <= "z") || ("0" <= string(c) && string(c) <= "9") {
			str += string(c)
		}
	}

	r := len(str) - 1

	for i := 0; i < r; i++ {
		if str[i] == str[r] {
			r--
			continue
		}
		return false
	}

	return true
}
