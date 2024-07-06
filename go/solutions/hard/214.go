package hard

func ShortestPalindrome(s string) string {
	length := len(s)
	for i := 0; i < length; i++ {
		if isPalindrome(s[:length-i]) {
			return reverseString(s)[:i] + s
		}
	}
	return reverseString(s) + s
}
func reverseString(s string) string {
	ss := []byte(s)
	length := len(s)
	for i := 0; i < length/2; i++ {
		ss[i], ss[length-1-i] = ss[length-1-i], ss[i]
	}
	return string(ss)
}

func isPalindrome(s string) bool {
	length:=len(s)
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}
