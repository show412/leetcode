// https://leetcode.com/problems/minimum-cost-for-tickets/
/*
In a country popular for train travel,
you have planned some train travelling one year in advance.
The days of the year that you will travel is given as an array days.
Each day is an integer from 1 to 365.

Train tickets are sold in 3 different ways:

a 1-day pass is sold for costs[0] dollars;
a 7-day pass is sold for costs[1] dollars;
a 30-day pass is sold for costs[2] dollars.
The passes allow that many days of consecutive travel.
For example, if we get a 7-day pass on day 2,
then we can travel for 7 days: day 2, 3, 4, 5, 6, 7, and 8.

Return the minimum number of dollars
you need to travel every day in the given list of days.



Example 1:

Input: days = [1,4,6,7,8,20], costs = [2,7,15]
Output: 11
Explanation:
For example, here is one way to buy passes that lets you travel your travel plan:
On day 1, you bought a 1-day pass for costs[0] = $2, which covered day 1.
On day 3, you bought a 7-day pass for costs[1] = $7, which covered days 3, 4, ..., 9.
On day 20, you bought a 1-day pass for costs[0] = $2, which covered day 20.
In total you spent $11 and covered all the days of your travel.
Example 2:

Input: days = [1,2,3,4,5,6,7,8,9,10,30,31], costs = [2,7,15]
Output: 17
Explanation:
For example, here is one way to buy passes that lets you travel your travel plan:
On day 1, you bought a 30-day pass for costs[2] = $15 which covered days 1, 2, ..., 30.
On day 31, you bought a 1-day pass for costs[0] = $2 which covered day 31.
In total you spent $17 and covered all the days of your travel.


Note:

1 <= days.length <= 365
1 <= days[i] <= 365
days is in strictly increasing order.
costs.length == 3
1 <= costs[i] <= 1000
*/
// refer to https://leetcode.com/problems/minimum-cost-for-tickets/discuss/226659/Two-DP-solutions-with-pictures
func mincostTickets(days []int, costs []int) int {
	if len(days) == 1 {
		return costs[0]
	}
	dp := make([]int, 366)
	dayset := make(map[int]bool, 0)
	for _, v := range days {
		dayset[v] = true
	}
	for d := 1; d < 366; d++ {
		// d不是travel day 和前一天的cost相等 这点比较难想
		if _, ok := dayset[d]; !ok {
			dp[d] = dp[d-1]
		} else {
			// d是travel day
			dp[d] = min(dp[d-1]+costs[0], dp[max(0, d-7)]+costs[1])
			dp[d] = min(dp[d], dp[max(0, d-30)]+costs[2])
		}
	}
	return dp[365]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

