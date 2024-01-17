package longestsubstringnorepeats

func LengthOfLongestSubstring(s string) int {
	start, longest := 0, 0

	subMap := make(map[rune]int)

	for i, v := range s {
		if index, ok := subMap[v]; ok && index >= start {
			start = index + 1
		}
		subMap[v] = i
		solution := i - start + 1

		if solution >= longest {
			longest = solution
		}
	}

	return longest
}
