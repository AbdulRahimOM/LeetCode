package medium

import "strings"

func ReverseParentheses(s string) string {
	length := len(s)
	bytes := []byte(s)
	findAndReverseClosures(bytes, 0, length-1)
	return removeParenthesesAndGetString(bytes)
}
func findAndReverseClosures(s []byte, start, end int) {
	for i := start; i <= end; i++ {
		if s[i] == '(' {
			localStart := i
			rank := 0
			for i = i + 1; i <= end; i++ {
				if s[i] == '(' {
					rank++
				} else if s[i] == ')' {
					if rank == 0 {
						findAndReverseClosures(s, localStart+1, i-1)
						reverseText(s, localStart+1, i-1)
						break
					} else {
						rank--
					}
				}
			}

		}
	}
}
func reverseText(s []byte, start, end int) {
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
func removeParenthesesAndGetString(s []byte) string {
	var builder strings.Builder
	for _, b := range s {
		if b != '(' && b != ')' {
			builder.WriteByte(b)
		}
	}
	return builder.String()
}
