package main

import "fmt"

func PrintMatrix(m [][]int) {
	for _, r := range m {
		fmt.Println(r)
	}
}
func CompareMatrixes(a, b [][]int) bool {
	// Check if outer slice lengths are the same
	if len(a) != len(b) {
		return false
	}

	// Check if inner slice lengths and elements are the same
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}