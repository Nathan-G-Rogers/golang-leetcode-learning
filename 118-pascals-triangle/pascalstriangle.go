//package pascalstriangle

package main

import "fmt"

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{{1}}
	}
	return getAllRows(numRows - 1)

}

func getAllRows(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{{1}}
	}

	oList := getAllRows((numRows - 1))

	leng := len(oList[len(oList)-1])
	oList = append(oList, []int{1})

	for i := 0; i < leng-1; i++ {
		oList[len(oList)-1] = append(
			oList[len(oList)-1], oList[len(oList)-2][i]+oList[len(oList)-2][i+1])
	}

	oList[len(oList)-1] = append(oList[len(oList)-1], 1)

	return oList
}

func main() {
	fmt.Println(generate(5))
}
