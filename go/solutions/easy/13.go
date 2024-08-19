package easy

// When submitted: Runtime: 0ms, Memory: 3.01MB (Beats 100%, 20%)
func RomanToInt(s string) int {
	var result int16
	var lastLetter byte
	length := len(s)
	for i := length - 1; i >= 0; i-- {
		if lastLetter == s[i] {
			result += romanMap[s[i]]
		} else if romanMap[s[i]] > romanMap[lastLetter] {
			result += romanMap[s[i]]
			lastLetter = s[i]
		}else{ 
			result -= romanMap[s[i]]
		}
	}
	return int(result)
}

var romanMap = map[byte]int16{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}