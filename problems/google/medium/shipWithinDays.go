// https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/
/*
A conveyor belt has packages that must be shipped
from one port to another within D days.

The i-th package on the conveyor belt has a weight of weights[i].
Each day, we load the ship with packages on the conveyor belt
(in the order given by weights).
We may not load more weight than the maximum weight capacity of the ship.

Return the least weight capacity of the ship
that will result in all the packages on the conveyor belt being shipped within D days.



Example 1:

Input: weights = [1,2,3,4,5,6,7,8,9,10], D = 5
Output: 15
Explanation:
A ship capacity of 15 is the minimum to ship all the packages in 5 days like this:
1st day: 1, 2, 3, 4, 5
2nd day: 6, 7
3rd day: 8
4th day: 9
5th day: 10

Note that the cargo must be shipped in the order given,
so using a ship of capacity 14
and splitting the packages into parts
like (2, 3, 4, 5), (1, 6, 7), (8), (9), (10) is not allowed.
Example 2:

Input: weights = [3,2,2,4,1,4], D = 3
Output: 6
Explanation:
A ship capacity of 6 is the minimum to ship all the packages in 3 days like this:
1st day: 3, 2
2nd day: 2, 4
3rd day: 1, 4
Example 3:

Input: weights = [1,2,3,1,1], D = 4
Output: 3
Explanation:
1st day: 1
2nd day: 2
3rd day: 3
4th day: 1, 1


Note:

1 <= D <= weights.length <= 50000
1 <= weights[i] <= 500
*/
// 主要是这个二分的思路 中间这个加1减1的是难点
func shipWithinDays(weights []int, D int) int {
	start := 0
	end := 0
	// 这个中间weight 不一定就是weights里的前缀和的值 它只是一个值
	for i := 0; i < len(weights); i++ {
		// 减少二分的次数 提高start的值
		start = max(start, weights[i])
		end += weights[i]
	}
	if D == 1 {
		return end
	}
	// it needs to loop until start >= end , because we need to find out the min capacity
	for start < end {
		mid := start + (end-start)/2
		// bigger than 0 means the mid is small
		// so we need to make start = mid+1
		// in order to no LTE, start = mid+1 因为当前的mid不行 那start可能至少比mid大1
		if checkCapacity(weights, mid, D) > 0 {
			start = mid + 1
		} else {
			// if checkCapacity <=0, it means the mid is big
			// we could try to narrow the range
			end = mid
		}
	}
	// 最后取的较小的那个值 所以是start 不是mid
	return start
}

func checkCapacity(weights []int, capacity int, D int) int {
	// the count should be 1 because the last weight should be included
	count := 1
	sum := 0
	for i := 0; i < len(weights); i++ {
		if (sum + weights[i]) > capacity {
			count++
			// 意味着weights[i]装不下了 只能装到下一个package里 所以sum=weights[i]
			sum = weights[i]
		} else {
			sum += weights[i]
		}
	}
	if count > D {
		return 1
	} else if count < D {
		return -1
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
