package main

import (
	"fmt"
	"math"
)

func main() {
	res := titleToNumber("ZY")
	fmt.Println(res)
}

func titleToNumber(s string) int {
	if s == "" {
		return 0
	}
	arr := []rune(s)
	length := len(arr)
	res := 0
	base := int('A') - 1
	for i := 0; i < length; i++ {
		charNum := int(arr[i]) - base
		fmt.Println(charNum)
		if length >= 2 {
			res += charNum * int(math.Pow(float64(26), float64(length-1-i)))
		} else {
			res = charNum
		}
	}
	return res
}
