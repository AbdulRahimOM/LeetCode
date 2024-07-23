package easy

import "sort"

// When submitted: Runtime: 5 ms, Memory: 3.26 MB (beats 74%, 30%)
func FrequencySort(nums []int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}
	numFreqs := make([][2]int, 0, len(freqMap))
	for num, freq := range freqMap {
		numFreqs = append(numFreqs, [2]int{num, freq})
	}

	sort.Slice(numFreqs, func(i, j int) bool {
		if numFreqs[i][1] == numFreqs[j][1] {
			return numFreqs[i][0] > numFreqs[j][0]
		} else {
			return numFreqs[i][1] < numFreqs[j][1]
		}
	})

	nums = nums[:0]
	for i := range numFreqs {
		for range numFreqs[i][1] {
			nums = append(nums, numFreqs[i][0])
		}
	}
	return nums
}
