package main

import (
	"bytes"
	"fmt"
	// "math"
)

func main() {
	// reg := regexp.MustCompile("([a-z]*)([0-9]+)(\\[\1)(.)*(\\]\1)")
	// str := "aaaa3[a2[cc]]2[bc]"
	// data := reg.FindAllStringSubmatch(str, -1)
	// fmt.Println(data)
	// 1, 2, 3, 0, 2
	res := reverseVowels("OE")
	fmt.Println(res)
}

func reverseVowels(s string) string {

	start := 0
	end := len(s) - 1
	sByte := []byte(s)
	set := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for start < end {
		f := bytes.IndexByte(set, sByte[start])
		e := bytes.IndexByte(set, sByte[end])
		// string couldn't assign or change it's value
		// it must to convert to byte
		if f >= 0 && e >= 0 {
			sByte[start], sByte[end] = sByte[end], sByte[start]
			start++
			end--
			continue
		}
		if f < 0 {
			start++
		}
		if e < 0 {
			end--
		}
	}

	return string(sByte)
}
