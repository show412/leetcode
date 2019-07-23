package main

import (
	"fmt"
	"strconv"
	// "math"
)

/*
test case:
"kkkkzrkatkwpkkkktrq"
"bbbbaaaaababaababab"
*/
func main() {
	res := diffWaysToCompute("2*3-4*5")
	fmt.Println(res)
}

func diffWaysToCompute(input string) []int {
	res := make([]int, 0)
	for i := 0; i < len(input); i++ {
		if input[i] == '-' || input[i] == '+' || input[i] == '*' {
			part1 := ([]byte(input))[:i]
			part2 := ([]byte(input))[(i + 1):]
			// fmt.Println(string(part1))
			part1Res := diffWaysToCompute(string(part1))
			part2Res := diffWaysToCompute(string(part2))
			// fmt.Println("****")
			// fmt.Println(part1Res)
			// fmt.Println(part2Res)
			for _, p1 := range part1Res {
				for _, p2 := range part2Res {
					c := 0
					if input[i] == '+' {
						c = p1 + p2
					} else if input[i] == '-' {
						c = p1 - p2
					} else if input[i] == '*' {
						c = p1 * p2
					}
					res = append(res, c)
				}
			}
		}
	}
	if len(res) == 0 {
		// fmt.Printf("%T, %v", input, input)
		v, _ := strconv.Atoi(input)
		// fmt.Printf("%T, %v", v, v)
		res = append(res, v)
		// return res
		// fmt.Println(res)
	}
	return res
}
