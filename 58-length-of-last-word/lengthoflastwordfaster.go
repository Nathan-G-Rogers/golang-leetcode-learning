package lengthoflastword

func FasterLengthOfLastWord(s string) int {
	w := len(s)
	count := 0

	for i := w - 1; i >= 0; i-- {
		if count == 0 && s[i] == ' ' {
			continue
		} else {
			if s[i] == ' ' {
				break
			}

			count++
		}
	}

	return count
}
