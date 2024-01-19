package medianoftwosortedarrays

import (
	"math"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	length1, length2 := len(nums1), len(nums2)

	min, max := 0, length1

	for min <= max {

		i := (min + max) / 2
		j := ((length1 + length2 + 1) / 2) - i

		maxLeftNums1 := math.MinInt
		if i > 0 {
			maxLeftNums1 = nums1[i-1]
		}

		minRightNums1 := math.MaxInt
		if i < length1 {
			minRightNums1 = nums1[i]
		}

		maxLeftNums2 := math.MinInt
		if j > 0 {
			maxLeftNums2 = nums2[j-1]
		}

		minRightNums2 := math.MaxInt
		if j < length2 {
			minRightNums2 = nums2[j]
		}

		if maxLeftNums1 <= minRightNums2 && maxLeftNums2 <= minRightNums1 {
			if (length1+length2)%2 == 0 {
				return float64((calcMax(maxLeftNums1, maxLeftNums2) + calcMin(minRightNums1, minRightNums2))) / 2.0
			} else {
				return math.Max(float64(maxLeftNums1), float64(maxLeftNums2))
			}
		} else if maxLeftNums1 > minRightNums2 {
			max = i - 1
		} else {
			min = i + 1
		}

	}

	//should never get to this return unless there is an error
	return 0.0
}

func calcMax(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func calcMin(x, y int) int {
	if x > y {
		return y
	}
	return x
}
