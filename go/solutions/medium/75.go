package medium

func SortColors(nums []int) {
	var count0,count1 int
	for _, v := range nums {
		switch v {
		case 0:
			count0++
		case 1:
			count1++
		}
	}
	var i int
	for i < count0{
		nums[i] = 0
        i++
	}
    count0+=count1
    for i < count0{
		nums[i] = 1
        i++
	}
    count0=len(nums)
	for i < count0{
		nums[i] = 2
        i++
	}
}