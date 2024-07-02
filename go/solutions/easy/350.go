package easy

func Intersect(nums1 []int, nums2 []int) []int {
	map1 := make(map[int]int16)
	results := []int{}
	for _, v := range nums1 {
		map1[v] = map1[v] + 1
	}
	for k, count1 := range map1 {
		var count2 int16
		for _, v := range nums2 {
			if v == k {
				count2++
				results = append(results, k)
				if count2 == count1 {
					break
				}
			}
		}
	}
	return results
}