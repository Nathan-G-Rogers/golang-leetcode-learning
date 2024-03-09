/*125. Valid Palindrome
Easy
Topics
Companies
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.



Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.


Constraints:

1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.
*/

package main

import (
	"fmt"
	"strings"
)

func clean(s []byte) string {
	j := 0

	for _, b := range s {
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') {
			s[j] = b
			j++
		}
	}

	return strings.ToLower(string(s[:j]))
}

func isPalindrome(s string) bool {
	cleaned := clean([]byte(s))

	if len(cleaned) <= 1 {
		return true
	}

	leng := len(cleaned)

	var i int

	switch {
	case leng%2 == 0:
		i = (leng / 2) - 1
	default:
		i = (leng / 2)
	}

	j := (leng / 2)

	for i >= 0 && j < leng {
		fmt.Println(i, j)
		if cleaned[i] != cleaned[j] {
			return false
		}
		i--
		j++
	}
	return true
}

func main() {
	s := "aa"

	b := isPalindrome(s)

	s1 := "ab"

	b1 := isPalindrome(s1)
	fmt.Println(b)
	fmt.Println(b1)
}
