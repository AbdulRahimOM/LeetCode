package hard

func LongestValidParentheses(s string) int {
	for i := range s {
		if s[i] == '(' {
            s = s[i:]
			break
		}
	}

	rankInitiatedAt := []int{0}
	rank := 0
	longest := 0
	for i :=range s {
		if s[i] == '(' {
			rank++
			rankInitiatedAt = append(rankInitiatedAt, i+1)
		} else {
			rank--
			if rank >= 0 {
				longest = max(longest, i+1-rankInitiatedAt[rank])
				rankInitiatedAt=rankInitiatedAt[:rank+1]
			} else {
				rank = 0
				rankInitiatedAt = []int{i+1}
			}
		}
	}
	return longest
}