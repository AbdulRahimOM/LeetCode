package easy

func MissingNumber(nums []int) int {
    count:=0
    n:=len(nums)
    for _,v:=range nums{
        count+=v
    }
    return n*(n+1)/2-count
}