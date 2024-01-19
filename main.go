package main

import (
	"fmt"

	medianoftwosortedarrays "github.com/Nathan-G-Rogers/golang-leetcode-learning/04-median-of-two-sorted-arrays"
)

func main() {
	// fmt.Println("Hello World")

	// solution := twosum.TwoSum([]int{2, 9, 11, 20}, 29)

	// fmt.Println(solution)

	// l1 := &addtwonumbers.ListNode{9, &addtwonumbers.ListNode{9, &addtwonumbers.ListNode{9, &addtwonumbers.ListNode{9, &addtwonumbers.ListNode{9, &addtwonumbers.ListNode{9, nil}}}}}}
	// l2 := &addtwonumbers.ListNode{9, &addtwonumbers.ListNode{9, &addtwonumbers.ListNode{9, &addtwonumbers.ListNode{9, &addtwonumbers.ListNode{9, nil}}}}}

	// solution2 := addtwonumbers.AddTwoNumbers(l1, l2)

	// for node := solution2; node != nil; node = node.Next {
	// 	fmt.Print(node.Val)
	// }

	//solution3 := longestsubstringnorepeats.LengthOfLongestSubstring("abcdabcdf")

	//fmt.Println(solution3)

	solution4 := medianoftwosortedarrays.FindMedianSortedArrays([]int{1, 2}, []int{3, 4})

	fmt.Println(solution4)

}
