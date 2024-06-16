package medium

/* 	I couldn't solve this problem with my earlier approaches (time limit exceeded)
*	my alternative approach didn't work, gave mistakes (which is also given below (named as OldApproach_countCompleteDayPairs))

*	I learned this algorithm from the 1st ranker
*	This and similar approach would be useful in many complement pairs scenarios, etc.
 */

func CountCompleteDayPairs_II(hours []int) int64 {
	hoursCollection := make([]int, 24)
	count := 0
	for i := range hours {
		remainder := hours[i] % 24
		count += hoursCollection[(24-remainder)%24]
		hoursCollection[remainder]++
	}
	return int64(count)
}

// This is the old approach which gave wrong answer in a testcase, couldn't understand the mistake as the input was large
func OldApproach_countCompleteDayPairs(hours []int) int64 {
	hoursCollection := make([]int, 24)
	for i := range hours {
		hoursCollection[hours[i]%24]++
	}
	var count int = combination(hoursCollection[0]) + combination(hoursCollection[12])
	for i := 1; i < 12; i++ {
		count += hoursCollection[i] * hoursCollection[24-i]
	}
	return int64(count)
}
func combination(n int) int {
	return factorial(n) / (2 * factorial(n-2))
}
func factorial(a int) int {
	if a <= 1 {
		return 1
	}
	result := 1
	for i := 2; i <= a; i++ {
		result *= i
	}
	return result
}
