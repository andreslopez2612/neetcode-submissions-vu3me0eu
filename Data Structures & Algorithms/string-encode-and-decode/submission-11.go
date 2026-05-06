type Solution struct{}

func (s *Solution) Encode(strs []string) string {

	res := ""
	for _, str := range strs {
		res += strconv.Itoa(len(str)) + "#" + str
	}
	return res
}

func (s *Solution) Decode(encoded string) []string {

	result := []string{}
	character := "#"
	i := 0

	for i < len(encoded) {
		j := i

		for string(encoded[j]) != character {
			j++
		}

		length, _ := strconv.Atoi(string(encoded[i:j]))
		i = j + 1
		result = append(result, encoded[i : i+length])
		i += length
	}

	return result
}
