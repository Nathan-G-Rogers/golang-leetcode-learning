package concatenationofarray

func getConcatenationFaster(nums []int) []int {
	return append(nums, nums...)
}
