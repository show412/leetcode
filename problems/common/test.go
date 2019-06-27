package main

import (
	"fmt"
	"math"
)

func main() {
	// reg := regexp.MustCompile("([a-z]*)([0-9]+)(\\[\1)(.)*(\\]\1)")
	// str := "aaaa3[a2[cc]]2[bc]"
	// data := reg.FindAllStringSubmatch(str, -1)
	// fmt.Println(data)
	// 1, 2, 3, 0, 2
	res := maxProfit([]int{3, 3})
	fmt.Println(res)
}
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	profit := 0
	for i := 0; i < len(prices); i++ {
		lastPrice := math.MaxInt32
		nextPrice := math.MinInt32
		price := prices[i]

		if i != 0 {
			lastPrice = prices[i-1]
		}

		if i != len(prices)-1 {
			nextPrice = prices[i+1]
		}
		// fmt.Println("********")
		// fmt.Println(lastPrice)
		// fmt.Println(nextPrice)
		// fmt.Println(price)
		if price < lastPrice && price <= nextPrice {
			profit -= price
		}
		if price >= lastPrice && price > nextPrice {
			profit += price
		}

	}
	return profit
}
