package main

import (
	"fmt"
	// "math"
)

func main() {
	// "e##e#o##oyof##q"
	// "e##e#fq##o##oyof##q"
	// "xywrrmp", "xywrrmu#p"
	// "bxj##tw"
	// "bxj###tw"
	// 	"bxj##tw"
	// "bxo#j##tw"
	res := backspaceCompare("bxj##tw", "bxo#j##tw")
	fmt.Println(res)
}

func backspaceCompare(S string, T string) bool {
	if len(S) == 0 && len(T) == 0 {
		return true
	}
	start1 := len(S) - 1
	start2 := len(T) - 1
	jump1 := 0
	jump2 := 0
	for start1 >= 0 || start2 >= 0 {
		for start1 >= 0 && S[start1] == '#' {
			jump1++
			start1--
		}
		for start2 >= 0 && T[start2] == '#' {
			jump2++
			start2--
		}
		for jump1 > 0 && start1 >= 0 {

			if start1 >= 0 && S[start1] != '#' {
				start1--
				jump1--
			}
			if start1 >= 0 && S[start1] == '#' {
				start1--
				jump1++
			}
		}
		for jump2 > 0 && start2 >= 0 {

			if start2 >= 0 && T[start2] != '#' {
				start2--
				jump2--
			}

			if start2 >= 0 && T[start2] == '#' {
				start2--
				jump2++
			}
		}
		// fmt.Println("********")
		// fmt.Println(start1)
		// fmt.Println(start2)
		if start1 < 0 && start2 < 0 {
			return true
		}
		if start1 >= 0 && start2 >= 0 && S[start1] != T[start2] {
			fmt.Println(1)
			fmt.Println(start1)
			fmt.Println(start2)
			return false
		}
		if (start1 < 0 && start2 >= 0) || (start1 >= 0 && start2 < 0) {
			fmt.Println(2)
			fmt.Println(start1)
			fmt.Println(start2)
			return false
		}
		start1--
		start2--
	}

	// if start1 >= 0 || start2 >= 0 {
	// 	return false
	// }
	return true
}
