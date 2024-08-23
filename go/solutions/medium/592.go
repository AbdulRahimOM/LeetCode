package medium

import (
	"fmt"
	"strconv"
)

func FractionAddition(expression string) string {
	numerator := 0
	denominator := 1
	numStart := 0
	finalNumerator := 0
	finalDenominator := 1
	for i, c := range expression {
		if c == '/' {
			numerator, _ = strconv.Atoi(expression[:i])
			expression = expression[i+1:]
			break
		}
	}
	for i, c := range expression {
		switch c {
		case '+', '-':
			denominator, _ = strconv.Atoi(expression[numStart:i])
			numStart = i

			finalNumerator = finalNumerator*denominator + numerator*finalDenominator
			finalDenominator *= denominator
		case '/':
			numerator, _ = strconv.Atoi(expression[numStart:i])
			numStart = i + 1
		}
	}
	denominator, _ = strconv.Atoi(expression[numStart:])
	finalNumerator = finalNumerator*denominator + numerator*finalDenominator
	finalDenominator *= denominator

	if finalNumerator < 0 {
		gcd := GCD(-finalNumerator, finalDenominator)
		return fmt.Sprintf("-%d/%d", -finalNumerator/gcd, finalDenominator/gcd)
	}
	gcd := GCD(finalNumerator, finalDenominator)
	return fmt.Sprintf("%d/%d", finalNumerator/gcd, finalDenominator/gcd)
}
func GCD(a, b int) int { // Greatest Common Divisor
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}
