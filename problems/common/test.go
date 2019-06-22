package main

import (
	"fmt"
)

func main() {
	// reg := regexp.MustCompile("([a-z]*)([0-9]+)(\\[\1)(.)*(\\]\1)")
	// str := "aaaa3[a2[cc]]2[bc]"
	// data := reg.FindAllStringSubmatch(str, -1)
	// fmt.Println(data)
	res := addStrings("123", "1789")
	fmt.Println(res)
}

// Input:  [0,1,2,4,5,7]
// Output: ["0->2","4->5","7"]
func addStrings(num1 string, num2 string) string {
	// the result length should be the longest between num1 and num2 to add 1
	// because it maybe have a carry
	res := make([]byte, len(num1)+1)
	if len(num1) < len(num2) {
		res = make([]byte, len(num2)+1)
	}
	// from end of num1 and num2 to add by byte
	// if one of num1 or num2 has no digits, add the left disgits into res
	j, z := len(num1)-1, len(num2)-1
	for i := len(res) - 1; i >= 0; i-- {
		if j >= 0 && z < 0 {
			res[i] = num1[j] - '0'
		} else if j < 0 && z >= 0 {
			res[i] = num2[z] - '0'
		} else if j >= 0 && z >= 0 {
			res[i] = (num1[j] - '0') + (num2[z] - '0')
		}
		j--
		z--
	}

	// Deal with the carry
	for i := len(res) - 1; i > 0; i-- {
		res[i-1] = res[i]/10 + res[i-1]
		res[i] = res[i] % 10
	}
	// work out the result
	var str []byte
	for i := 0; i < len(res); i++ {
		if i == 0 && res[i] == 0 {
			continue
		}
		str = append(str, res[i]+'0')
	}
	if len(str) == 0 {
		return "0"
	}
	return string(str)
}
