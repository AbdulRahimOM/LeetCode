package easy

func ScoreOfString(s string) int {
	var sum int16
	for i := len(s)-1; i >0; i-- {
		sum += absIntOfByte(s[i], s[i-1])
	}
	return int(sum)
}
func absIntOfByte(char1, char2 byte) int16 {
	if char1 < char2 {
		return int16(char2 - char1)
	} else {
		return int16(char1 - char2)
	}
}