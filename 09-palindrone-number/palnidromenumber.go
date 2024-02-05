package palindromenumber

import "fmt"

func isPalindrome(x int) bool {
	str := fmt.Sprint(x)

	length := len(str)

	if length == 0 {
		return false
	}

	if length == 1 {
		return true
	}

	for i, j := 0, length-1; i < length && j >= 0; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
