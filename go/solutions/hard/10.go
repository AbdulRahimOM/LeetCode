package hard

import (
	"fmt"
	"strings"
	"time"
)

func Ouuu() {
	// fmt.Println("clear=", clearStars("d*o*"))
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println("s=", s)
	
	fmt.Println("s=", s)

}

type linkedList struct {
	root *node
}
type node struct {
	data int
	next *node
}

func clearStars(s string) string {
	var charCount [26]int32
	var minIndex int32
	bytes := []byte(s)
	for i, v := range s {
		if v != '*' {
			charCount[v-97]++
			if minIndex > v-97 {
				minIndex = v - 97
			}
		} else {
			for minIndex < 26 {
				if charCount[minIndex] > 0 {
					charCount[minIndex]--
					for j := i - 1; j >= 0; j-- {
						if bytes[j] == byte(minIndex)+97 {
							bytes[j] = '*'
							break
						}
					}
					break
				} else {
					minIndex++
				}
			}
		}
	}
	var builder strings.Builder
	for _, v := range bytes {
		if v != '*' {
			builder.WriteByte(v)
		}
	}
	return builder.String()
}

type meetChain struct {
	start int
	end   int
	next  *meetChain
}

// Regular Expression Matching
func countDays(days int, meetings [][]int) int {
	calendar := &meetChain{
		start: meetings[0][0],
		end:   meetings[0][1],
	}
	for _, v := range meetings {
		calendar.add(v[0], v[1])
	}
	return calendar.getGaps(days)
}
func (c *meetChain) getGaps(days int) int {
	current := c
	for current != nil {
		days -= current.end - current.start + 1
		current = current.next
	}
	return days

}
func (c *meetChain) add(start, end int) {
	current := c
	for current != nil {
		if start > current.end {
			current = current.next
			continue
		} else if end < current.start {
			current.start, current.end, current.next = c.start, c.end, current
			break
		} else if end <= current.end {
			if start < current.start {
				current.start = start
			}
			break
		} else if start >= current.start {
			if end > current.end {
				current.end = end
			}
			break
		} else {
			current.start = start
			current.end = end
			for current.next != nil {
				if current.next.end >= end {
					current.next = current.next.next
				} else if current.next.start > end {
					break
				} else {
					current.end = current.next.end
					break
				}
			}
		}
		current = current.next
	}
}

// Regular Expression Matching
func LongcountDays(days int, meetings [][]int) int {
	var collective []bool
	absence := days
	collective = make([]bool, days+1)
	for _, v := range meetings {
		for start := v[0]; start <= v[1]; start++ {
			if !collective[start] {
				collective[start] = true
				absence--
			}
		}
	}
	return absence
}
func isMatch(s string, p string) bool {
	// for t := 1; t < 10; t++ {
	fmt.Println("======================")
	fmt.Println("reciedved s=", s, ", p=", p)
	if s == "" && p == "" {
		return true
	}
	if s == "" || p == "" {
		return false
	}
	// }
	pLength := len(p)
	sLength := len(s)
	j := 0
	var prevLetter rune
	for i, v := range s {
		fmt.Println("======================")
		fmt.Printf("received s[i]=%c, p[j]=%c\n", s[i], p[j])

		time.Sleep(1 * time.Second)
		switch p[j] {
		case '*':
			if prevLetter == '.' {
				var obeyed bool
				for k := range s[i:] {
					fmt.Println("----------------------------------")
					if isMatch(s[i+k:], p[j+1:]) {
						obeyed = true
						break
					}
				}
				if !obeyed {
					return false
				}
			} else {
				for j+1 < pLength { //is this required???
					if p[j+1] == byte(prevLetter) {
						j++
					} else {
						break
					}
				}
				if j+2 < pLength {
					if p[j+1] == '.' && p[j+2] == '*' {
						//to weherever the letter is repeated in s, skip until that
					}
				}

				fmt.Println("else played")
				time.Sleep(2 * time.Second)
				var obeyed bool
				for k := i; k < sLength; k++ {
					if s[k] != byte(prevLetter) {
						break
					}
					fmt.Println("--")
					time.Sleep(2 * time.Second)

					if isMatch(s[i+k:], p[j+1:]) {
						fmt.Println("++")
						obeyed = true
						break
					} else {
						fmt.Println("returned a false")
						time.Sleep(2 * time.Second)

					}
				}
				if !obeyed {
					return false
				}
			}
		case '.':
			prevLetter = '.'
			if j+1 != pLength {
				if p[j+1] == '*' {
					i--
				}
			}
		default:
			prevLetter = v
			if j+1 < pLength {
				if p[j+1] != '*' && p[j] != s[i] {
					return false
				}
			} else {
				if p[j] != s[i] {
					return false
				}
			}
		}
		j++
		if j == pLength {
			return false
		}
	}
	return true

}
