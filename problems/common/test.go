package main

import (
	"fmt"
	"strconv"
)

func main() {
	// reg := regexp.MustCompile("([a-z]*)([0-9]+)(\\[\1)(.)*(\\]\1)")
	// str := "aaaa3[a2[cc]]2[bc]"
	// data := reg.FindAllStringSubmatch(str, -1)
	// fmt.Println(data)
	res := decodeString("3[z]2[2[y]pq4[2[jk]e1[f]]]ef")
	fmt.Println(res)
}

func decodeString(s string) string {
	sByte := []byte(s)
	start := 0
	decodeS := decode(sByte, &start)
	return string(decodeS)
}

func decode(s []byte, i *int) []byte {
	var res []byte
	var numByte []byte

	for *i < len(s) && s[*i] != ']' {
		if is_digit(s[*i]) == true {
			numByte = append(numByte, s[*i])
			(*i)++
		} else if s[*i] == '[' {
			k, _ := strconv.Atoi(string(numByte))
			(*i)++
			sub := decode(s, i)
			for j := 0; j < k; j++ {
				res = append(res, sub...)
			}
			(*i)++
			// notice we need to clear the numByte because there will add last number
			numByte = []byte{}
		} else {
			res = append(res, s[*i])
			(*i)++
		}
	}
	return res
}

func is_digit(c byte) bool {
	return c >= '0' && c <= '9'
}
