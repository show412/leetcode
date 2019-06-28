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
	res := maxProfit(2, []int{3, 2, 6, 5, 0, 3})
	fmt.Println(res)
}
func maxProfit(k int, prices []int) int {
	if len(prices) <= 1 || k == 0 {
		return 0
	}
	profit := 0
	if k >= len(prices)/2 {
		for i := 1; i < len(prices); i++ {
			if prices[i] > prices[i-1] {
				profit += (prices[i] - prices[i-1])
			}
		}
		return profit
	}
	// dp[i][j] means the max profit j days finish i times transaction
	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, len(prices))
	}
	dp[0][0] = 0
	for i := 0; i < len(dp); i++ {
		dp[i][0] = 0
	}
	for j := 0; j < len(dp[0]); j++ {
		dp[0][j] = 0
	}
	// function
	for i := 1; i <= k; i++ {
		buy := -prices[0]
		for j := 1; j < len(prices); j++ {
			dp[i][j] = max(dp[i][j-1], prices[j]+buy)
			buy = max(buy, dp[i-1][j-1]-prices[j])
		}
	}
	return dp[k][len(prices)-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
