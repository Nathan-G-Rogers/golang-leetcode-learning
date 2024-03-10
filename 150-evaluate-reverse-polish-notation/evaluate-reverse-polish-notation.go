/*
150. Evaluate Reverse Polish Notation

You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.

Evaluate the expression. Return an integer that represents the value of the expression.

Note that:

The valid operators are '+', '-', '*', and '/'.
Each operand may be an integer or another expression.
The division between two integers always truncates toward zero.
There will not be any division by zero.
The input represents a valid arithmetic expression in a reverse polish notation.
The answer and all the intermediate calculations can be represented in a 32-bit integer.


Example 1:

Input: tokens = ["2","1","+","3","*"]
Output: 9
Explanation: ((2 + 1) * 3) = 9
Example 2:

Input: tokens = ["4","13","5","/","+"]
Output: 6
Explanation: (4 + (13 / 5)) = 6
Example 3:

Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
Output: 22
Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22
*/
//package evaluatereversepolishnotation
package main

import (
	"fmt"
	"strconv"
)

type RPN struct {
	values []int
}

func (r *RPN) pop() int {
	number := r.values[len(r.values)-1]
	r.values = r.values[:len(r.values)-1]
	return number
}

func evalRPN(tokens []string) int {
	total := RPN{values: []int{}}

	for i := range tokens {
		switch tokens[i] {
		case "*":
			val1 := total.pop()
			val2 := total.pop()
			total.values = append(total.values, (val1 * val2))
		case "-":
			val1 := total.pop()
			val2 := total.pop()
			total.values = append(total.values, (val1 - val2))
		case "+":
			val1 := total.pop()
			val2 := total.pop()
			total.values = append(total.values, (val1 + val2))
		case "/":
			val1 := total.pop()
			val2 := total.pop()
			total.values = append(total.values, (val2 / val1))
		default:
			num, _ := strconv.Atoi((tokens[i]))
			total.values = append(total.values, num)
		}
	}

	return total.values[0]
}

func main() {

	testArr := []string{"4", "13", "5", "/", "+"}

	fmt.Println(13 / 5)

	fmt.Println(evalRPN(testArr))

}
