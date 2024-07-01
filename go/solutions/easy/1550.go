package easy

func ThreeConsecutiveOdds(arr []int) bool {
    length:=len(arr)
	i := -1
	for i < length-3 {
		i++
		if arr[i]%2 == 1 {
			i++
			if arr[i]%2 == 1 {
				i++
				if arr[i]%2 == 1 {
					return true
				}
			}
		}
	}
	return false
}