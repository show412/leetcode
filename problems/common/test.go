package main

import (
	"fmt"
)

/*
test case:
101, 56, 69, 48, 30
4

[8,85,24,85,69]
4

[73,106,39,6,26,15,30,100,71,35,46,112,6,60,110]
29

[98,60,24,89,84,51,61,96,108,87,68,29,14,11,13,50,13,104,57,8,57,111,92,87,9,59,65,116,56,39,55,11,21,105,57,36,48,93,20,94,35,68,64,41,37,11,50,47,8,9]
439
*/
func main() {
	res := multiply("999", "999")
	// 998001
	fmt.Println(res)
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	num1Array := []byte(num1)
	num2Array := []byte(num2)
	l1 := len(num1Array)
	l2 := len(num2Array)
	l := l1 + l2
	res := make([]int, l)
	str := ""
	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			mul := (num1[i] - '0') * (num2[j] - '0')
			p1 := i + j
			p2 := i + j + 1
			sum := int(mul) + res[p2]
			res[p1] += sum / 10
			res[p2] = sum % 10
		}
	}
	for _, v := range res {
		if str != "" || v != 0 {
			str += string(byte(v + '0'))
		}
	}
	return str
}
