package twosum

func TwoSum(nums []int, target int) []int {
	var solution []int

	m := make(map[int]int)

	for i, v := range nums {
		x := target - v

		_, ok := m[x]
		if ok {
			solution = append(solution, m[x], i)
			break
		} else {
			m[v] = i
		}
	}

	return solution
}
