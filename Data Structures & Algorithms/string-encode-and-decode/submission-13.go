type Solution struct{}

func (s *Solution) Encode(strs []string) string {

	var result strings.Builder

	for _, v := range strs {
		//Number to Int - Interger to ASCII
		val := strconv.Itoa(len(v))
		result.WriteString(val)
		result.WriteString("#")
		result.WriteString(v)
	}

	return result.String()
}

func (s *Solution) Decode(encoded string) []string {

	result := []string{}

	i := 0
	for i < len(encoded) {
		j := i

		for string(encoded[j]) != "#" {
			j++
		}
		length, _ := strconv.Atoi(encoded[i:j])
		index := j + 1
		final := j + 1 + length

		result = append(result, encoded[index:final])
		i = final
	}

	return result
}

