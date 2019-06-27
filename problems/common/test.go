package main

import (
	"fmt"
)

func main() {
	// reg := regexp.MustCompile("([a-z]*)([0-9]+)(\\[\1)(.)*(\\]\1)")
	// str := "aaaa3[a2[cc]]2[bc]"
	// data := reg.FindAllStringSubmatch(str, -1)
	// fmt.Println(data)
	// 1, 2, 3, 0, 2
	res := maxProfit([]int{1, 2, 3, 0, 2})
	fmt.Println(res)
}

// 这个算法的问题是 sell 不一定就要在 比buy的价格大的那天 sell, 是一个动态的过程
// so it looks like a DP solution
/*
	buy[i]  = max(rest[i-1]-price, buy[i-1])
 	sell[i] = max(buy[i-1]+price, sell[i-1])
	rest[i] = max(sell[i-1], buy[i-1], rest[i-1])
*/
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	buy := make([]int, len(prices))
	sell := make([]int, len(prices))
	cooldown := make([]int, len(prices))
	buy[0] = -prices[0]
	sell[0] = 0
	cooldown[0] = 0
	for i := 1; i < len(prices); i++ {
		price := prices[i]
		buy[i] = max(cooldown[i-1]-price, buy[i-1])
		sell[i] = max(buy[i-1]+price, sell[i-1])
		cooldown[i] = maxThree(sell[i-1], buy[i-1], cooldown[i-1])
	}
	return sell[len(prices)-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxThree(a int, b int, c int) int {
	if a > b {
		if a >= c {
			return a
		} else {
			return c
		}
	} else {
		if b >= c {
			return b
		} else {
			return c
		}
	}
	return a
}
