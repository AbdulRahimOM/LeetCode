package easy

func MinimumOperations(nums []int) int {
    count:=0
    for i:=range nums{
        if nums[i]%3!=0{
            count++
        }
    }
    return count
}