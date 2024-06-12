package easy

func RelativeSortArray(arr1 []int, arr2 []int) []int {
	netCountArr:=make([]int,1001)
	for _, v1 := range arr1 {
			netCountArr[v1]++
	}
	j := 0
	for _, v := range arr2 {
		vCount := netCountArr[v]
		netCountArr[v] = 0
		for vCount > 0 {
			arr1[j] = v
			j++
			vCount--
		}
	}
	for i, v := range netCountArr {
		for v > 0 {
			arr1[j] = i
			j++
			v--
		}
	}
	return arr1
}
