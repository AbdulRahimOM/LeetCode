package easy

// When submitted: Runtime: 77 ms, memory usage: 7.5 MB (beats 95.67%, 92.6%)
func LemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, bill := range bills {
		switch bill {
		case 5:
			five++
		case 10:
			if five == 0 {
				return false
			}
			five--
			ten++
		default:
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}