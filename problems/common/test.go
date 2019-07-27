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
	res := shortestWay("aaaaa", "aaaaaaaaaaaaa")
	fmt.Println(res)
}

func shortestWay(source string, target string) int {
	for i := 0; i < len(target); i++ {
		char := target[i]
		flag := false
		for j := 0; j < len(source); j++ {
			if source[j] == char {
				flag = true
				continue
			}
		}
		if flag == false {
			return -1
		}
	}
	pt := 0
	res := 0
	for pt < len(target) {
		ps := 0
		for ps < len(source) {
			if pt > len(target)-1 {
				break
			}
			if source[ps] != target[pt] {
				ps++
			} else {
				pt++
				ps++
			}
		}
		res++
	}
	return res
}
