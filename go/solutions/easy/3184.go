package easy

func CountCompleteDayPairs_I(hours []int) int {
	count:=0
	length:=len(hours)
	for i:=range hours{
		for j:=i+1;j<length;j++{
			if (hours[i]+hours[j])%24==0{
				count++
			}
		}
	}
	return count
}

//better algorithm available in medium/3185.go