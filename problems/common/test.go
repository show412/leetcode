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
	res := addBinary("0", "0")
	fmt.Println(res)
}

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	l := max(i+1, j+1) + 1
	res := make([]byte, l)
	// carry := byte(0) - '0'
	for i >= 0 && j >= 0 {
		a := a[i] - '0'
		b := b[j] - '0'
		res[l-1] = a + b
		i--
		j--
		l--
	}
	for i >= 0 {
		a := a[i] - '0'
		res[l-1] = a
		i--
		l--
	}
	for j >= 0 {
		b := b[j] - '0'
		res[l-1] = b
		j--
		l--
	}
	// deal with the carry
	for k := len(res) - 1; k >= 0; k-- {
		if k > 0 {
			res[k-1] = res[k]/2 + res[k-1]
		}

		res[k] = res[k]%2 + '0'
	}
	// when the first index is 0, no carry, remove the first index 0
	if res[0] == '0' {
		res = res[1:]
	}
	if len(res) == 0 {
		return "0"
	}

	return string(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
