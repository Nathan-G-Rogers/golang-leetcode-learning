/*
739. Daily Temperatures
Medium
Topics
Companies
Hint
Given an array of integers temperatures represents the daily temperatures, return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature. If there is no future day for which this is possible, keep answer[i] == 0 instead.

Example 1:

Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]
Example 2:

Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]
Example 3:

Input: temperatures = [30,60,90]
Output: [1,1,0]

Constraints:

1 <= temperatures.length <= 105
30 <= temperatures[i] <= 100
*/
package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))
	stack := make([]int, 0)

	for i, v := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < v {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[idx] = i - idx
		}
		stack = append(stack, i)
	}
	return answer
}

func main() {
	test := dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	fmt.Println(test)
}
