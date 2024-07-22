package easy

import "sort"

// When submitted: Runtime: 18 ms, Memory: 7.87 MB
func SortPeople(names []string, heights []int) []string {
	length := len(names)
	heightIndexMap := make(map[int]int, length)
	for i, height := range heights {
		heightIndexMap[height] = i
	}
	sort.Ints(heights)
	sortedNames := make([]string, length)
	for i, height := range heights {
		sortedNames[length-1-i] = names[heightIndexMap[height]]
	}

	return sortedNames
}

// When submitted: Runtime: 18 ms, Memory: 7.22 MB
func SortPeople_Attempt1(names []string, heights []int) []string {
	heightIndexList := make([][2]int, len(heights))
	for i, height := range heights {
		heightIndexList[i] = [2]int{i, height}
	}
	sort.Slice(heightIndexList, func(i, j int) bool {
		return heightIndexList[i][1] > heightIndexList[j][1]
	})
	sortedNames := make([]string, len(names))
	for i, heightIndex := range heightIndexList {
		sortedNames[i] = names[heightIndex[0]]
	}
	return sortedNames
}
