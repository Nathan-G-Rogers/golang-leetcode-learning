// package longestconsecutivesequence
package main

import "fmt"

// import "fmt"

// func longestConsecutive(nums []int) int {

// 	if len(nums) == 1 {
// 		return 1
// 	}

// 	var longest int
// 	newMap := make(map[int]bool)

// 	for _, v := range nums {

// 		newMap[v] = true
// 	}

// 	for _, v := range nums {

// 		if newMap[v-1] {
// 			continue
// 		}

// 		count := 1
// 		tempNum := v + 1

// 		for newMap[tempNum] {
// 			count++
// 			tempNum++
// 		}

// 		if count > longest {
// 			longest = count
// 		}
// 	}
// 	return longest
// }

func main() {
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}

func longestConsecutive(nums []int) int {

	if len(nums) == 1 {
		return 1
	}

	var longest int
	newMap := make(map[int]bool)

	for _, v := range nums {

		newMap[v] = true
	}

	for i := range newMap {

		if newMap[i-1] {
			continue
		}

		count := i

		for newMap[count+1] {
			count++
		}

		longest = max(longest, count-i+1)
	}
	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
