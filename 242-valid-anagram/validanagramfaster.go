package validanagram

func isAnagramFaster(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	newMap := make(map[rune]int)

	for i, _ := range s {
		newMap[rune(s[i])]++

	}

	for _, x := range t {
		if number, ok := newMap[x]; !ok || number == 0 {
			return false
		}
		newMap[x]--

	}

	return true

}
