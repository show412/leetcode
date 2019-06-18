package main

import (
	"fmt"
)

func main() {
	res := rotatedDigits(10000)
	fmt.Println(res)
}

func rotatedDigits(N int) int {
	if N < 0 {
		return 0
	}
	cnt := 0
	for i := 1; i <= N; i++ {
		if good(i) == true {
			// fmt.Println(i)
			cnt++
		}
	}
	return cnt
}
func good(num int) bool {
	flag := 0
	for num != 0 {
		if num%10 == 3 || num%10 == 4 || num%10 == 7 {
			return false
		}
		if num%10 == 0 || num%10 == 1 || num%10 == 8 {
			num = num / 10
			continue
		}
		if num%10 == 2 || num%10 == 5 || num%10 == 6 || num%10 == 9 {
			num = num / 10
			flag = 1
			continue
		}
	}
	if flag == 1 {
		return true
	}
	return false
}
