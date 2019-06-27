import "math"

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
/*
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).

Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

Example 1:

Input: [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
             Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
Example 2:

Input: [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
             Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are
             engaging multiple transactions at the same time. You must sell before buying again.
Example 3:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
*/

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	profit := 0
	for i := 0; i < len(prices); i++ {
		// it should be here not outside of loop
		// because nextPrice will be prices[len(prices)-1] if it is outside
		// otherwise we did
		/*
			if i == 0
			lastPrice = math.MaxInt32
			if i == len(prices) - 1
			nextPrice := math.MinInt32
		*/
		// it should be max for lastPrice, because it must to buy at first time unless the price[0] > price[1]
		// it should be min for nextPrice, becasue it must to sell at last time unless the price[len-1] < price[len-2]
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

		// in order to prevent this case: only sell no buy
		// notice there must be <=, because for this case [2,2,5]
		// it must to buy unless price > nextPrice
		if price < lastPrice && price <= nextPrice {
			profit -= price
		}
		// in order to prevent this case: only buy no sell
		// notice there must be <=, because for this case [3,3]
		// it must to sell unless price < lastPrice
		if price >= lastPrice && price > nextPrice {
			profit += price
		}

	}
	return profit
}

// greedy
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		// it looks like it means it's buy and sell at the same time
		// actually, suppose that from a to b is increased
		/*
			it means prices[b+1] - prices[b] + prices[b] - prices[b-1] ....prices[a+1] - prices[a]
			finally prices[b] - prices[a]
			whatever how many days from a to b, it's fine for it
			p[2]-p[1] + p[1] - p[0] => p[2]-p[0]
			其实就是波峰减波谷
		*/
		diff := prices[i+1] - prices[i]
		if diff > 0 {
			profit += diff
		}
	}
	return profit
}
