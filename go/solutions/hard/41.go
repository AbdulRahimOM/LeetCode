package hard

import "sort"

// Runtime: 52 ms, Memory: 7.94 MB (Beats 71.03%, 78.50%)
func FirstMissingPositive(nums []int) int {
	sort.Ints(nums)	// why can I use sort.Ints() here? Question says nothing about it
	missing := 1
	for _, n := range nums {
		if n == missing {
			missing++
		} else if n > missing {
			return missing
		}
	}
	return missing
}
