package main

import (
	"fmt"
	"regexp"
	"strconv"
	// "math"
)

func main() {
	// reg := regexp.MustCompile("([a-z]*)([0-9]+)(\\[\1)(.)*(\\]\1)")
	// str := "aaaa3[a2[cc]]2[bc]"
	// data := reg.FindAllStringSubmatch(str, -1)
	// fmt.Println(data)
	// 1, 2, 3, 0, 2
	// output [[1,1,0],[1,0,1],[0,0,0]]
	// expected [[1,0,0],[0,1,0],[1,1,1]]
	res := isPowerOfThree(12)
	fmt.Println(res)
}

func isPowerOfThree(n int) bool {
	// 转为3进制 return is string
	s3 := strconv.FormatInt(int64(n), 3)
	// reg := regexp.MustCompile("^10*$")
	m, _ := regexp.MatchString("^10*$", s3)
	return m
}
