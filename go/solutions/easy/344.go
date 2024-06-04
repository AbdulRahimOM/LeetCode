package easy

func ReverseString(s []byte) {
	n := make([]byte, len(s))
	for i, j := len(s)-1, 0; i >= 0; i-- {
		n[j] = s[i]
		j++
	}
	copy(s, n)
}
