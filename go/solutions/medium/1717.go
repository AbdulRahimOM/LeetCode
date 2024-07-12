package medium

//used int32 to reduce memory usage and improve performance
func MaximumGain(s string, x int, y int) int {
	var result,countA,countB int32
	xx,yy := int32(x),int32(y)
	if xx < yy {
		for _, v := range s {
			switch v {
			case 'a':
				if countB > 0 {
					countB--
					result += yy
				} else {
					countA++
				}
			case 'b':
				countB++
			default:
				result += min(countA, countB) * xx
				countA = 0
				countB = 0
			}
		}
		result += min(countA, countB) * xx
	} else if xx > yy {
		for _, v := range s {
			switch v {
			case 'b':
				if countA > 0 {
					countA--
					result += xx
				} else {
					countB++
				}
			case 'a':
				countA++
			default:
				result += min(countA, countB) * yy
				countA = 0
				countB = 0
			}
		}
		result += min(countA, countB) * yy
	} else {
		for _, v := range s {
			switch v {
			case 'a':
				countA++
			case 'b':
				countB++
			default:
				result += min(countA, countB) * xx
				countA = 0
				countB = 0
			}
		}
		result += min(countA, countB) * xx
	}
	return int(result)
}

// Earlier attempt, same algorithm but didn't use less sized variables like int32
func MaximumGain_1stAttempt(s string, x int, y int) int {
	result := 0
	countA := 0
	countB := 0
	if x < y {
		for _, v := range s {
			switch v {
			case 'a':
				if countB > 0 {
					countB--
					result += y
				} else {
					countA++
				}
			case 'b':
				countB++
			default:
				result += min(countA, countB) * x
				countA = 0
				countB = 0
			}
		}
		result += min(countA, countB) * x
	} else if x > y {
		for _, v := range s {
			switch v {
			case 'b':
				if countA > 0 {
					countA--
					result += x
				} else {
					countB++
				}
			case 'a':
				countA++
			default:
				result += min(countA, countB) * y
				countA = 0
				countB = 0
			}
		}
		result += min(countA, countB) * y
	} else {
		for _, v := range s {
			switch v {
			case 'a':
				countA++
			case 'b':
				countB++
			default:
				result += min(countA, countB) * x
				countA = 0
				countB = 0
			}
		}
		result += min(countA, countB) * x
	}
	return result
}