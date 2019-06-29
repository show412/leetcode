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
	res := judgeCircle("UD")
	fmt.Println(res)
}

func judgeCircle(moves string) bool {
	x, y := 0, 0
	for i := 0; i < len(moves); i++ {
		if moves[i] == 'U' {
			y++
		}
		if moves[i] == 'D' {
			y--
		}
		if moves[i] == 'R' {
			x++
		}
		if moves[i] == 'L' {
			x--
		}
	}
	return x == 0 && y == 0
}
