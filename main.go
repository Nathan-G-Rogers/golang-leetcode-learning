package main

import (
	"fmt"

	twosum "github.com/Nathan-G-Rogers/golang-leetcode-learning/01-two-sum"
	addtwonumbers "github.com/Nathan-G-Rogers/golang-leetcode-learning/02-add-two-numbers"
)

func main() {
	fmt.Println("Hello World")

	solution := twosum.TwoSum([]int{2, 9, 11, 20}, 29)

	fmt.Println(solution)

	l1 := &addtwonumbers.ListNode{9, &addtwonumbers.ListNode{9, &addtwonumbers.ListNode{9, &addtwonumbers.ListNode{9, &addtwonumbers.ListNode{9, &addtwonumbers.ListNode{9, nil}}}}}}
	l2 := &addtwonumbers.ListNode{9, &addtwonumbers.ListNode{9, &addtwonumbers.ListNode{9, &addtwonumbers.ListNode{9, &addtwonumbers.ListNode{9, nil}}}}}

	solution2 := addtwonumbers.AddTwoNumbers(l1, l2)

	for node := solution2; node != nil; node = node.Next {
		fmt.Print(node.Val)
	}
}
