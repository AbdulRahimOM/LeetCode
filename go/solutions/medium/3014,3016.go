package medium

import "sort"

func MinimumPushes(word string) int {
	letterCount := make([]int, 26)
	for _, v := range word {
		letterCount[v-'a']++
	}
	sort.Slice(letterCount, func(i, j int) bool {
		return letterCount[i] > letterCount[j]
	})

	var cost int
	for i, count := range letterCount {
		cost += count * (i/8 + 1)
	}
	return cost
}
