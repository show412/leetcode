package main

import (
	"fmt"
)

func main() {
	// reg := regexp.MustCompile("([a-z]*)([0-9]+)(\\[\1)(.)*(\\]\1)")
	// str := "aaaa3[a2[cc]]2[bc]"
	// data := reg.FindAllStringSubmatch(str, -1)
	// fmt.Println(data)
	res := islandPerimeter([]int{73, 74, 75, 71, 69, 72, 76, 73})
	// [1, 1, 4, 2, 1, 1, 0, 0]
	fmt.Println(res)
}

// Input:  [0,1,2,4,5,7]
// Output: ["0->2","4->5","7"]
func islandPerimeter(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				res += getLength(i, j, grid)
			}
		}
	}
	return res
}

func getLength(i int, j int, grid [][]int) int {
	length := 4
	if i-1 >= 0 && grid[i-1][j] == 1 {
		length -= 1
	}
	if j+1 < len(grid[0]) && grid[i][j+1] == 1 {
		length -= 1
	}
	if i+1 < len(grid) && grid[i+1][j] == 1 {
		length -= 1
	}
	if j-1 >= 0 && grid[i][j-1] == 1 {
		length -= 1
	}
	return length
}
