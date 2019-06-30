package main

import (
	"fmt"
	// "math"
)

func main() {
	// reg := regexp.MustCompile("([a-z]*)([0-9]+)(\\[\1)(.)*(\\]\1)")
	// str := "aaaa3[a2[cc]]2[bc]"
	// data := reg.FindAllStringSubmatch(str, -1)
	// fmt.Println(data)
	// 1, 2, 3, 0, 2
	// output [[1,1,0],[1,0,1],[0,0,0]]
	// expected [[1,0,0],[0,1,0],[1,1,1]]
	res := flipAndInvertImage([][]int{[]int{1, 1, 0, 0}, []int{1, 0, 0, 1}, []int{0, 1, 1, 1}, []int{1, 0, 1, 0}})
	fmt.Println(res)
}

func flipAndInvertImage(A [][]int) [][]int {
	if len(A) <= 0 {
		return A
	}
	r := len(A)
	c := len(A[0])
	cl := c/2 + 1
	if len(A[0])%2 == 0 {
		cl = c / 2
	}
	for i := 0; i < r; i++ {
		for j := 0; j < cl; j++ {
			A[i][j] ^= 1
			if j < c/2 {
				A[i][c-j-1] ^= 1
				A[i][j], A[i][c-1-j] = A[i][c-1-j], A[i][j]
			}
		}
	}
	// fmt.Println(A)
	return A
}
