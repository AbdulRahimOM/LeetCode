package medium

import "sort"

func MinDifference(nums []int) int {
	sort.Ints(nums)
	n := len(nums) - 4
	if n < 0 {
		return 0
	}
	return min(nums[n]-nums[0], nums[n+1]-nums[1], nums[n+2]-nums[2], nums[n+3]-nums[3])
}

// func MinDifference_1(nums []int) int {	//initial attempt (1)
// 	sort.Ints(nums)
// 	length := len(nums)
// 	if length <= 4 {
// 		return 0
// 	}
// 	n := length - 4
// 	minDiff := nums[n] - nums[0]
// 	for i := 1; i < 4; i++ {
// 		diff := nums[i+n] - nums[i]
// 		if diff < minDiff {
// 			minDiff = diff
// 		}
// 	}
// 	return minDiff
// }
