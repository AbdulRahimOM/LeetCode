package easy

func CanAliceWin(nums []int) bool {
    sumSingles:=0
	sumTens:=0
	for _,v:=range nums{
		if v<100&&v>9{
			sumTens+=v
		}else if v<10 && v>0{
			sumSingles+=v
		}
	}
	return sumSingles>sumTens || sumTens>sumSingles
}