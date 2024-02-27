package validanagram

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	newMapS := make(map[rune]int)
	newMapT := make(map[rune]int)

	for i, _ := range s {
		newMapS[rune(s[i])]++
		newMapT[rune(t[i])]++
	}

	for _, x := range s {
		if _, ok := newMapS[x]; !ok {
			return false
		}

		if _, okt := newMapT[x]; !okt {
			return false
		}

		if newMapS[x] != newMapT[x] {
			return false
		}

	}

	return true

}
