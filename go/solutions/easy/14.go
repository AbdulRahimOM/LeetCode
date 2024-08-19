package easy

import "strings"

//When submitted: Runtime: 0ms, Memory: 2.46MB (Beats 100%, 99.89%)
func LongestCommonPrefix(strs []string) string {
	prefix := strings.Builder{}
	minLen := 200
	for _, str := range strs {
		minLen = minimumOf(minLen, len(str))
	}
	if minLen == 0 {
		return ""
	}
	for i := range minLen {
		letter := strs[0][i]
		for _, str := range strs {
			if str[i] != letter {
				return prefix.String()
			}
		}
		prefix.WriteByte(letter)
	}
	return prefix.String()
}
func minimumOf(a, b int) int {
	if a < b {
		return a
	}
	return b
}
