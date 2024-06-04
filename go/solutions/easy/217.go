package easy

import "sort"

// Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

// Example 1:

// Input: nums = [1,2,3,1]
// Output: true
// Example 2:

// Input: nums = [1,2,3,4]
// Output: false
// Example 3:

// Input: nums = [1,1,1,3,3,4,3,2,4,2]
// Output: true

// Constraints:

// 1 <= nums.length <= 105

//containsDuplicate(nums []int) bool

func ContainsDuplicate(nums []int) bool {
	sort.Ints(nums)

	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}
