package easy

import "sort"

func MinimumAverage(nums []int) float64 {
	sort.Ints(nums)
	length := len(nums)
	minAvg := float64(nums[0]+nums[length-1]) / 2

	for i := 1; i < length/2; i++ {
		minAvg = min(minAvg, float64(nums[i]+nums[length-i-1])/2)
	}

	return minAvg
}
