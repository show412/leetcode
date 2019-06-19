package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	reg := regexp.MustCompile("([a-z]*)([0-9]+)(\\[\1)(.)*(\\]\1)")
	str := "aaaa3[a2[cc]]2[bc]"
	data := reg.FindAllStringSubmatch(str, -1)
	fmt.Println(data)
	// res := decodeString("3[a]2[bc]")
	// fmt.Println(res)
}

func decodeString(s string) string {
	reg := regexp.MustCompile("([0-9]+)(\\[)(.*)(\\])")
	data := reg.FindStringSubmatch(s)
	if len(data) == 0 {
		return s
	}
	k, _ := strconv.Atoi(data[1])
	subStr := data[3]
	res := ""
	for i := 0; i < k; i++ {
		res += decodeString(subStr)
	}
	return res
}
