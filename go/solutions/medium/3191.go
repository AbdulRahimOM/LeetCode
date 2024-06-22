package medium

func MinOperations_I(nums []int) int {
	length := len(nums)
	count := 0
	for i := 0; i < length-2; i++ {
		if nums[i] == 0 {
            // nums[i] = 1	// <---no need of this over-write. Add count and pass to next
			toggle(nums, i+1)
			toggle(nums, i+2)
			count++
		}
	}
    if nums[length-1] == 1 && nums[length-2] == 1{
        return count
    }
	return -1
}
func toggle(nums []int, i int) {
	if nums[i] == 0 {
		nums[i] = 1
	} else {
		nums[i] = 0
	}

}
