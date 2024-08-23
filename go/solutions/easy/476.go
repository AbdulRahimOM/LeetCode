package easy

import (
	"math"
	"strconv"
	"strings"
)

// When submitted: Runtime: 0 ms, memory usage: 2.16 MB (beats 100%,66.67%)
func FindComplement(num int) int {
	n := math.Log2(float64(num))
	return num ^ (1<<uint(n+1) - 1)
}

// old attempt
// When submitted: Runtime: 1 ms, memory usage: 2.2 MB (beats 100%, 66.67%)
func FindComplement_OldAttempt_II(num int) int {
	binaryStr := strconv.FormatInt(int64(num), 2)
	result := 0
	for _, c := range binaryStr {
		if c == '0' {
			result = result*2 + 1
		} else {
			result *= 2
		}
	}
	return result
}

// even older attempt
// When submitted: Runtime: 1 ms, memory usage: 2.2 MB (beats 100%, 66.67%)
func FindComplement_OldAttempt_I(num int) int {
	binaryStr := strconv.FormatInt(int64(num), 2)
	builder := strings.Builder{}
	for _, c := range binaryStr {
		if c == '0' {
			builder.WriteRune('1')
		} else {
			builder.WriteRune('0')
		}
	}
	binaryStr = builder.String()
	result, _ := strconv.ParseInt(binaryStr, 2, 64)
	return int(result)
}
