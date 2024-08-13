package hard

func FindSubstring(s string, words []string) []int {
	lenW := len(words[0])
	length := len(words)
	lenS := len(s)
	startGap := lenW*length - lenW
	if lenS < lenW*length {
		return []int{}
	}

	var marker []bool
	mapper := make(map[string]int)
	duplicateCount := make(map[string]int)
	var result []int
	for i, w := range words {
		if _, ok := mapper[w]; !ok {
			mapper[w] = i
		} else {
			duplicateCount[w]++
		}
	}

	for i := range lenW {
		count := 0
		marker = make([]bool, lenS)
		for i <= lenS-lenW {
			myWord := s[i : i+lenW]
			if _, ok := mapper[myWord]; !ok {
				count = 0
				marker = make([]bool, lenS)
				i += lenW
				continue
			}
			wordFound := false
			for k, w := range words {
				if w == myWord && !marker[k] {
					wordFound = true
					marker[k] = true
					count++
					if count >= length {
						result = append(result, i-startGap)
						marker[mapper[s[i-startGap:i-startGap+lenW]]] = false
					}
					break

				}
			}
			if !wordFound {
				if _, ok := mapper[myWord]; ok {
					index := 0
					start := i - count*lenW
					for {
						thisWord := s[start+index*lenW : start+index*lenW+lenW]
						if thisWord == myWord {
							break
						}
						for k, w := range words {
							if w == thisWord && marker[k] {
								marker[k] = false
								count--
								break
							}
						}
						index++
					}
				} else {
					count = 0
					marker = make([]bool, lenS)
				}
			}
			i += lenW
		}
	}
	return result
}
