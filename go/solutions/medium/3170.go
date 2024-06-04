// https://leetcode.com/problems/lexicographically-minimum-string-after-removing-stars/
package medium

import "strings"

func ClearStars(s string) string {
	var charCount [26]int32
	var minIndex int32
	bytes := []byte(s)
	for i, v := range s {
		if v != '*' {
			charCount[v-97]++
			if minIndex > v-97 {
				minIndex = v - 97
			}
		} else {
			for minIndex < 26 {
				if charCount[minIndex] > 0 {
					charCount[minIndex]--
					for j := i - 1; j >= 0; j-- {
						if bytes[j] == byte(minIndex)+97 {
							bytes[j] = '*'
							break
						}
					}
					break
				} else {
					minIndex++
				}
			}
		}
	}
	var builder strings.Builder
	for _, v := range bytes {
		if v != '*' {
			builder.WriteByte(v)
		}
	}
	return builder.String()
}
