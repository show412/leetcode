package main

import (
	"fmt"
)

func main() {
	// reg := regexp.MustCompile("([a-z]*)([0-9]+)(\\[\1)(.)*(\\]\1)")
	// str := "aaaa3[a2[cc]]2[bc]"
	// data := reg.FindAllStringSubmatch(str, -1)
	// fmt.Println(data)
	res := numSquares(5678)
	// [1, 1, 4, 2, 1, 1, 0, 0]
	fmt.Println(res)
}

// Input:  [0,1,2,4,5,7]
// Output: ["0->2","4->5","7"]
// DP solution
func numSquares(n int) int {
	if n <= 0 {
		return 0
	}
	// f means i need the square sum at least
	f := make([]int, n+1)
	f[0] = 0
	for i := 1; i <= n; i++ {
		// i could be the first i-1 square sum + 1
		f[i] = f[i-1] + 1
		for j := 1; j*j <= i; j++ {
			// For each i, it must be the sum of some number (i - j*j) and
			// a perfect square number (j*j).
			f[i] = min(f[i], f[i-j*j]+1)
		}
	}
	return f[n]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
