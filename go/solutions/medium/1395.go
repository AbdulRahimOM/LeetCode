package medium

func NumTeams(rating []int) int {
	var result int32
	n := len(rating)
	rightGreats := make(map[int]int32, n)
	leftGreats := make(map[int]int32, n)

	for i:=len(rating)-1;i>=0;i-- {
		rate:=rating[i]
		var newCount int32
		for mapRate, count := range rightGreats {
			if rate<mapRate {
				result += count
				newCount ++
			}
		}
		rightGreats[rate] = newCount
	}
	for _, rate := range rating {
		var newCount int32
		for mapRate, count := range leftGreats {
			if rate<mapRate {
				result += count
				newCount ++
			}
		}
		leftGreats[rate] = newCount
	}
	return int(result)
}
