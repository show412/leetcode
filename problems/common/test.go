package main

import (
	"fmt"
	// "math"
)

func main() {

	res := removeStones([][]int{[]int{0, 0}, []int{0, 1}, []int{1, 0}, []int{1, 2}, []int{2, 1}, []int{2, 2}})
	fmt.Println(res)
}

func removeStones(stones [][]int) int {
	fa := make([]int, 20000)
	for i := 0; i < len(stones); i++ {
		fa[i] = i
	}

	for i := 0; i < len(stones); i++ {
		unity(stones[i][0], stones[i][1]+10000, fa)
	}
	set := make(map[int]bool)
	for i := 0; i < len(stones); i++ {
		stone := find(stones[i][0], fa)
		set[stone] = true
	}
	return len(stones) - len(set)
}

func unity(x int, y int, fa []int) []int {
	x = find(x, fa)
	y = find(y, fa)
	fa[x] = y
	return fa
}

func find(x int, fa []int) int {
	if fa[x] == x {
		return x
	} else {
		fa[x] = find(fa[x], fa)
	}
	return fa[x]
}
