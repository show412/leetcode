package main

import (
	"fmt"
	"unicode"
)

/*
test case:
[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3
10

[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2
6
*/
func main() {
	res := calculate("3+5 / 2")
	// 998001
	fmt.Println(res)
}

func calculate(s string) int {
	l1, o1, l2, o2 := 0, 1, 1, 1
	if s[0] == '-' {
		s = "0" + s
	}
	for i := 0; i < len(s); i++ {
		c := rune(s[i])
		if unicode.IsDigit(c) {
			num := int(c - '0')
			for (i+1) < len(s) && unicode.IsDigit(rune(s[i+1])) {
				num = num*10 + int(s[i+1]-'0')
				i++
			}
			if o2 == 1 {
				l2 = l2 * num
			} else {
				l2 = l2 / num
			}
		} else if c == '*' || c == '/' {
			if c == '*' {
				o2 = 1
			} else {
				o2 = -1
			}
		} else if c == '+' || c == '-' {
			l1 = l1 + o1*l2
			if c == '+' {
				o1 = 1
			} else {
				o1 = -1
			}
			o2 = 1
			l2 = 1
		}
	}
	return l1 + o1*l2
}
