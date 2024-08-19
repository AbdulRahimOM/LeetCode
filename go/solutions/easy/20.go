package easy

var oppMap map[byte]byte = map[byte]byte{
	'{': '}',
	'[': ']',
	'(': ')',
}

// When submitted: Runtime: 1ms, Memory: 2.26MB (Beats 73%, 72%)
func IsValid(s string) bool {
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '{' || c == '[' || c == '(' {
			stack = append(stack, c)
		} else if len(stack) == 0 || c != oppMap[stack[len(stack)-1]] {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}