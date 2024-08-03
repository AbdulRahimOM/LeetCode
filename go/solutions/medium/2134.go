package medium

// MinSwaps II
func MinSwapsII(nums []int) int {
	length := int32(len(nums))
	var counts [2]int32

	for _, v := range nums {
		if v == 0 {
			counts[0]++
		}
	}
	counts[1] = length - counts[0]
	if counts[0] == 0 || counts[1] == 0 {
		return 0
	}

	var start, end int32
	nums = append(nums, nums...)
	if nums[length-1] == 0 && nums[0] == 1 {
		start = 0
	} else {
		for nums[start] != 0 {
			start++
		}
		for nums[start] != 1 {
			start++
		}
		nums = nums[start:]
		start = 0
	}

	var minim int32 = 100000000
	var cnt [2]int32

	for start < length {
		for {
			cnt[nums[end]]++
			end++
			if cnt[0]+cnt[1] == counts[1] {
				minim = min(minim, cnt[0])
				break
			}
		}

		for nums[start] != 0 {
			cnt[nums[start]]--
			start++
		}
		for nums[start] != 1 {
			cnt[nums[start]]--
			start++
		}
	}

	return int(minim)
}

// func min(a, b int32) int32 {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
