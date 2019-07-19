package main

import (
	"fmt"
	// "math"
)

/*
test case:
"kkkkzrkatkwpkkkktrq"
"bbbbaaaaababaababab"
*/
func main() {
	res := canTransform("XLLR", "LXLX")
	fmt.Println(res)
}

func canTransform(start string, end string) bool {
	if len(start) == 1 {
		if start == end {
			return true
		} else {
			return false
		}
	}
	i := 0
	j := 0
	for i < len(start) && j < len(start) {
		for i < len(start) && start[i] == 'X' {
			i++
		}
		for j < len(start) && end[j] == 'X' {
			j++
		}
		if (i < len(start) && j >= len(start)) || (i >= len(start) && j < len(start)) {
			return false
		}
		if i < len(start) && j < len(start) {
			if start[i] != end[j] || (start[i] == 'L' && i < j) || (start[i] == 'R' && i > j) {
				return false
			}
		}
		i++
		j++
	}
	return true
}
