package hard

import (
	"fmt"
	"sort"
	"strings"
)

func CountOfAtoms(formula string) string {
	elemCountMap := solveBracket(formula)
	type ElementNCount struct {
		elem  string
		count int
	}

	var elemCountSlice []ElementNCount
	for k, v := range elemCountMap {
		elemCountSlice = append(elemCountSlice, ElementNCount{k, v})
	}
	sort.Slice(elemCountSlice, func(i, j int) bool {
		return elemCountSlice[i].elem < elemCountSlice[j].elem
	})
	result := strings.Builder{}
	for _, elemCount := range elemCountSlice {
		result.WriteString(elemCount.elem + getCountString(elemCount.count))
	}
	return result.String()

}

func getCountString(count int) string {
	if count == 1 {
		return ""
	}
	return fmt.Sprintf("%d", count)
}
func solveBracket(formula string) map[string]int {
	elemCount := make(map[string]int)
	newElemCount := make(map[string]int)
	count := 0
	currentElem := ""
	length := len(formula)
	foundElem := false
	for i := 0; i < length; i++ {
		switch {
		case formula[i] >= 'A' && formula[i] <= 'Z':
			if foundElem {
				elemCount[currentElem] += max(count, 1)
				count = 0
			}
			currentElem = string(formula[i])
			foundElem = true
		case formula[i] >= 'a' && formula[i] <= 'z':
			currentElem += string(formula[i])
		case formula[i] >= '0' && formula[i] <= '9':
			count = count*10 + int(formula[i]-48)
		case formula[i] == '(':
			if foundElem {
				elemCount[currentElem] += max(count, 1)
				count = 0
				foundElem = false
			}

			//find the closing bracket, and solve until the closing bracket
			rank := 1
			for j := i + 1; j < len(formula); j++ {
				if formula[j] == '(' {
					rank++
				} else if formula[j] == ')' {
					rank--
					if rank == 0 {
						newElemCount = solveBracket(formula[i+1 : j])
						i = j
						break
					}
				}
			}

			//get the count after the bracket
			i++
			for i < length && formula[i] >= '0' && formula[i] <= '9' {
				count = count*10 + int(formula[i]-48)
				i++
			}
			i--
			count = max(count, 1)

			//add elements in the bracket (count times) to the current elemCount
			for k, v := range newElemCount {
				elemCount[k] += v * count
			}
			count = 0 //reset count
		}
	}
	if foundElem { //if some element remains to be added to elemCount
		elemCount[currentElem] += max(count, 1)
	}
	return elemCount
}
