package easy

import "fmt"

func AlternateDigitSum(n int) int {
	sum := 0
	str := fmt.Sprintf("%d", n)
	sign := 1
	for _, v := range str {
		sum += int(v-48) * sign
		sign = sign * -1
	}
	return sum
}
