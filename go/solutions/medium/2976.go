package medium

const overlimit = 1000001 *26

// When submitted: Runtime: 50ms, Memory: 6.99MB (Beats 77%, 100%)
func MinimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	costGrid := [26][26]int{}
	for i := range 26 {
		for j := range 26 {
			costGrid[i][j] = overlimit
		}
	}
	for i, v := range original {
		if costGrid[v-'a'][changed[i]-'a'] > cost[i] {
			costGrid[v-'a'][changed[i]-'a'] = cost[i]
		}
	}
	for middle := range 26 {
		for start := range 26 {
			for end := range 26 {
				costGrid[start][end] = min(costGrid[start][end], costGrid[start][middle]+costGrid[middle][end])
			}
		}
	}

	var minCost int
	for i := range source {
		if source[i] == target[i] {
			continue
		}
		cost := costGrid[source[i]-'a'][target[i]-'a']
		if cost == overlimit {
			return -1
		}
		minCost += cost
	}
	return int64(minCost)
}