package longestpalindromicsubstring

//general idea is to start from the middle and work outwards 2 characters at a time

func LongestPalindrome(s string) string {
	ans := []int{0, 0}

	for i := range s {

		odd_len := Expand(i, i, s)
		if odd_len > ans[1]-ans[0]+1 {
			dist := odd_len / 2
			ans[0] = i - dist
			ans[1] = i + dist
		}

		even_len := Expand(i, i+1, s)
		if even_len > ans[1]-ans[0]+1 {
			dist2 := (even_len / 2) - 1
			ans[0] = i - dist2
			ans[1] = i + dist2 + 1
		}
	}
	return s[ans[0] : ans[1]+1]
}

func Expand(i int, j int, str string) (length int) {
	left := i
	right := j

	for left >= 0 && right < len(str) && str[left] == str[right] {
		left--
		right++
	}
	return right - left - 1
}
