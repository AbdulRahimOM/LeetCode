package medium

func MinOperations_II(nums []int) int {
	length := len(nums)
	count := 0
	i := 0
    
    current:=1
    for i < length {
        if nums[i] == current {
            i++
        } else {
            current = nums[i]
            count++
            i++
        }
    }

	return count
}