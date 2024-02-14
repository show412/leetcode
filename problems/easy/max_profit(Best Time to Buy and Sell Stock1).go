/*
 * @Author: hongwei.sun
 * @Date: 2024-02-14 15:42:43
 * @LastEditors: your name
 * @LastEditTime: 2024-02-14 15:57:48
 * @Description: file content
 */
import "math"

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
/*
Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction
(i.e., buy one and sell one share of the stock),
design an algorithm to find the maximum profit.

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

// slide windows
// it's the best solution with TC O(n), SC O(1)
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	minPrice := math.MaxInt32
	profit := 0
	for i := 0; i < len(prices); i++ {
		price := prices[i]
		// it's one slide windows implementation to jump to one index by comparing
		if price < minPrice {
			minPrice = price
		}
		if price-minPrice > profit {
			profit = price - minPrice
		}
	}
	return profit

	// if len(prices) <= 1 {
	// 	return 0
	// }
	// min := math.MaxInt32
	// profit := 0
	// for i := 0; i < len(prices); i++ {
	// 	price := prices[i]
	// 	// notice this kind of advanced write for min
	// 	if price < min {
	// 		min = price
	// 	}
	// 	// notice this kind of advanced write for max
	// 	if (price - min) > profit {
	// 		profit = price - min
	// 	}
	// }
	// return profit
}

