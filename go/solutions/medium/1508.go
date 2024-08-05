package medium

import "sort"

func RangeSum(nums []int, n int, left int, right int) int {
	sumArray := []int{}
	localSumArray := make([]int, len(nums))
	copy(localSumArray, nums)
	maxx := 0
	for _, v := range nums {
		if v > maxx {
			maxx = v
		}
	}
	sumArray = append(sumArray, nums...)

	for xtraDigits := 1; xtraDigits < n; xtraDigits++ {
		locMin := 101

		for i, v := range nums[xtraDigits:] {
			localSumArray[i] += v
		}
		for _, v := range localSumArray[:n-xtraDigits] {
			if v < locMin {
				locMin = v
			}
		}
		if len(sumArray) >= right && locMin >= maxx {
			break // no need to find more
		}

		sumArray = append(sumArray, localSumArray[:n-xtraDigits]...)
	}
	sort.Ints(sumArray)
	result := 0
	for _, v := range sumArray[left-1 : right] {
		result += v
	}
	return result % 1000000007
}
