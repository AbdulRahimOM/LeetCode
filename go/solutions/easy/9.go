package easy

import "fmt"

func IsPalindrome(x int) bool {
	txt := fmt.Sprintf("%d", x)
	for i := 0; i < len(txt)/2; i++ {
		if txt[i] != txt[len(txt)-1-i] {
			return false
		}
	}
	return true
}
