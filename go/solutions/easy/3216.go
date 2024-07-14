package easy

import "fmt"

func GetSmallestString(s string) string {
	ss := []byte(s)
	for i := 0; i < len(s)-1; i++ {
		if s[i] != '0' {
			if s[i+1] < s[i] {
				if s[i+1]%2 == s[i]%2 {
					fmt.Println("i: ", i)
					ss[i], ss[i+1] = s[i+1], ss[i]
					return string(ss)
				}
			}

		}
	}
	return s
}
