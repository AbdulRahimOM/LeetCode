package medium

func MinimumDeletions(s string) int {
	starta := 0
	endb := len(s) - 1
	for i, v := range s {
		if v == 'b' {
			starta = i
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'a' {
			endb = i
			break
		}
	}
	countSlice := make([]int, endb-starta+2)

	count := 0
	for i, v := range s[starta : endb+1] {
		if v == 'b' {
			count++
		}
		countSlice[i+1] = count
	}
	count = 0
	for i := endb; i >= starta; i-- {
		if s[i] == 'a' {
			count++
		}
		countSlice[i-starta] += count

	}

	min := countSlice[0]
	for _, v := range countSlice {
		if v < min {
			min = v
		}
	}
	return min
}