package easy

func LongestPalindrome(s string) int {
	count := 0
	a := make([]int, 200)
	for _, v := range s {
		a[v]++
	}
	singleMarked := false
	for _, v := range a {
		if v > 0 {
			if v %2 == 1 {
				if !singleMarked {
					singleMarked = true
					count+=v
				}else{
                    count+=v-1
                }
                
			} else {
				count += v
			}
		}
	}
	return count
}