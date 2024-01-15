package main

import (
	"fmt"

	twosum "github.com/Nathan-G-Rogers/golang-leetcode-learning/01-two-sum"
)

func main() {
	fmt.Println("Hello World")

	solution := twosum.TwoSum([]int{2, 9, 11, 20}, 29)

	fmt.Println(solution)
}
