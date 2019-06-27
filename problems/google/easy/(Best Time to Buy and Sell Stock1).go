import "math"

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
/*
Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

Example 1:

Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.
Example 2:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
*/

// it's the best solution with TC O(n), SC O(1)
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	min := math.MaxInt32
	profit := 0
	for i := 0; i < len(prices); i++ {
		price := prices[i]
		// notice this kind of advanced write for min
		if price < min {
			min = price
		}
		// notice this kind of advanced write for max
		if (price - min) > profit {
			profit = price - min
		}
	}
	return profit
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	// here cap is 0, because we will append the slice below
	prefixMin := make([]int, 0)
	suffixMax := make([]int, len(prices))
	res := 0
	for i := 0; i < len(prices); i++ {
		if i == 0 {
			prefixMin = append(prefixMin, prices[i])
		} else {
			if prices[i] <= prefixMin[len(prefixMin)-1] {
				prefixMin = append(prefixMin, prices[i])
			} else {
				prefixMin = append(prefixMin, prefixMin[len(prefixMin)-1])
			}
		}
	}

	for i := len(prices) - 1; i >= 0; i-- {
		if i == len(prices)-1 {
			suffixMax[i] = prices[i]
		} else {
			if prices[i] > suffixMax[i+1] {
				suffixMax[i] = prices[i]
			} else {
				suffixMax[i] = suffixMax[i+1]
			}
		}
	}
	for i := 0; i < len(prices); i++ {
		res = max(res, suffixMax[i]-prefixMin[i])
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

