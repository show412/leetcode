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
	// "2-4A0r7-4k", 3
	res := licenseKeyFormatting("--a-a-a-a--", 2)
	fmt.Println(res)
}

func licenseKeyFormatting(S string, K int) string {
	sByte := []byte(S)
	l := len(sByte)
	// res := make([]byte, 2*l)
	// index := len(res) - 1
	var res []byte
	count := 0
	for i := l - 1; i >= 0; i-- {
		if sByte[i] >= 'a' && sByte[i] <= 'z' {
			sByte[i] = sByte[i] - ('a' - 'A')
		}

		if sByte[i] != '-' && count < K {
			// res[index] = sByte[i]
			res = append(res, sByte[i])

			// index--
			count++
		}
		if count == K && i != 0 {
			// res[index] = '-'
			res = append(res, '-')

			// index--
			count = 0
		}
	}
	// if index < 0 {
	// 	index = 0
	// }
	// formate := make([]byte, len(res))
	for i := 0; i < (len(res)+1)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	if len(res) == 0 {
		return ""
	}
	if res[0] == '-' {
		res = res[1:]
	}
	return string(res)
}
