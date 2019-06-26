package main

import (
	"fmt"
)

func main() {
	// reg := regexp.MustCompile("([a-z]*)([0-9]+)(\\[\1)(.)*(\\]\1)")
	// str := "aaaa3[a2[cc]]2[bc]"
	// data := reg.FindAllStringSubmatch(str, -1)
	// fmt.Println(data)
	res := maxProfit([]int{1, 2, 4})
	fmt.Println(res)
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	res := 0
	// var benift []int
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] > prices[i] {
				res = max(res, res+prices[j]-prices[i])
				j++
				i = j + 1
			}
		}
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
