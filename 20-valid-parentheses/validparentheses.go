/*
20. Valid Parentheses

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.

Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false

Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/
package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	validParan := map[rune]rune{'}': '{',
		']': '[',
		')': '(',
	}

	for _, v := range s {
		if _, ok := validParan[v]; !ok {
			stack = append(stack, v)
			continue

		}

		if len(stack) == 0 {
			return false
		}

		if stack[len(stack)-1] != validParan[v] {
			return false

		}

		stack = stack[:len(stack)-1]

		fmt.Println(stack)
		fmt.Println(v)

	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("(])"))
}
