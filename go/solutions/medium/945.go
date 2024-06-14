package medium

import "sort"

func MinIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	var moves int32
	var max int = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		} else {
			max++
			moves += int32(max - nums[i])
		}
	}
	return int(moves)
}
