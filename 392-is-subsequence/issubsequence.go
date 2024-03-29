package issubsequence

func isSubsequence(s string, t string) bool {

	for i, j := 0, 0; i < len(s) && j < len(t); {
		if s[i] == t[j] {
			i++
		}
		j++

		if i == len(s) {
			return true
		}
	}

	return false
}
