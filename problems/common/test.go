package main

import (
	"fmt"
	"unicode"
)

/*
test case:
"1 + 1" = 2
" 6-4 / 2 " = 4
"2*(5+5*2)/3+(6/2+8)" = 21
"(2+6* 3+5- (3*14/7+2)*5)+3"=-12
"-1+4*3/3/3" 0
*/
func main() {
	res := calculate("-1+4*3/3/3")
	// res := canFinish(4, [][]int{[]int{0, 1}})
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
				num = num*10 + int(rune(s[i+1])-'0')
				i++
			}

			if o2 == 1 {
				l2 = l2 * num
			} else {
				l2 = l2 / num
			}
		} else if c == '(' {
			j := i
			for cnt := 0; i < len(s); i++ {
				if s[i] == '(' {
					cnt++
				}
				if s[i] == ')' {
					cnt--
				}
				if cnt == 0 {
					break
				}
			}
			num := calculate(s[(j + 1):i])

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
			l2 = 1
			o2 = 1
		}
		fmt.Println("****")
		fmt.Println(l1)
		fmt.Println(o1)
		fmt.Println(l2)
		fmt.Println(o2)
	}
	return (l1 + o1*l2)
}
