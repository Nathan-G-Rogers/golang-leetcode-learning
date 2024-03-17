package main

import (
	"fmt"
	"strings"
)

func generateParenthesis(n int) []string {
	stack := []string{}
	result := []string{}

	var recurse func(openN int, closedN int)

	recurse = func(openN int, closedN int) {
		if openN == n && closedN == n {
			result = append(result, strings.Join(stack, ""))
			return
		}

		if openN < n {
			stack = append(stack, "(")
			recurse(openN+1, closedN)
			stack = stack[:len(stack)-1]

		}
		if closedN < openN {
			stack = append(stack, ")")
			recurse(openN, closedN+1)
			stack = stack[:len(stack)-1]
		}
	}
	recurse(0, 0)

	return result
}

func main() {
	test := (generateParenthesis(2))

	fmt.Println(test)
}
