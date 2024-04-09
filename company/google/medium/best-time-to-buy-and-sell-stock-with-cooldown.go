// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
/*
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit.
You may complete as many transactions as you like
(ie, buy one and sell one share of the stock multiple times)
with the following restrictions:

You may not engage in multiple transactions at the same time
(ie, you must sell the stock before you buy again).
After you sell your stock, you cannot buy stock on next day.
(ie, cooldown 1 day)
Example:

Input: [1,2,3,0,2]
Output: 3
Explanation: transactions = [buy, sell, cooldown, buy, sell]
*/
// recusive or dp?
// DP solution
/*
定义两个数组buy 和 sell 代表在i天是buy or cooldown 和 sell or cooldown的max profit
*/
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	buy := make([]int, len(prices))
	sell := make([]int, len(prices))
	buy[0] = -prices[0]
	buy[1] = max(buy[0], -prices[1])
	sell[0] = 0
	sell[1] = max(sell[0], prices[1]+buy[0])
	for i := 2; i < len(prices); i++ {
		// one day is cooldown after sell
		buy[i] = max(buy[i-1], sell[i-2]-prices[i])
		sell[i] = max(sell[i-1], buy[i-1]+prices[i])
	}
	return sell[len(prices)-1]
}

/*
The series of problems are typical dp. The key for dp is to find the variables to represent the states and deduce the transition function.
Of course one may come up with a O(1) space solution directly, but I think it is better to be generous when you think and be greedy when you implement.
The natural states for this problem is the 3 possible transactions : buy, sell, rest. Here rest means no transaction on that day (aka cooldown).
Then the transaction sequences can end with any of these three states.
For each of them we make an array, buy[n], sell[n] and rest[n].
buy[i] means before day i what is the maxProfit for any sequence end with buy.
sell[i] means before day i what is the maxProfit for any sequence end with sell.
rest[i] means before day i what is the maxProfit for any sequence end with rest.
Then we want to deduce the transition functions for buy sell and rest. By definition we have:
buy[i]  = max(rest[i-1]-price, buy[i-1])
sell[i] = max(buy[i-1]+price, sell[i-1])
rest[i] = max(sell[i-1], buy[i-1], rest[i-1])
Where price is the price of day i. All of these are very straightforward. They simply represents :
(1) We have to `rest` before we `buy` and
(2) we have to `buy` before we `sell`
One tricky point is how do you make sure you sell before you buy, since from the equations it seems that [buy, rest, buy] is entirely possible.
Well, the answer lies within the fact that buy[i] <= rest[i] which means rest[i] = max(sell[i-1], rest[i-1]). That made sure [buy, rest, buy] is never occurred.
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
