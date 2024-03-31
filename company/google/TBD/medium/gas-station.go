/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:52
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-03-26 17:36:47
 * @Description: file content
 */
// https://leetcode.com/problems/gas-station/
/*

There are N gas stations along a circular route, where the amount of gas at station i is gas[i].

You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from station i to its next station (i+1). You begin the journey with an empty tank at one of the gas stations.

Return the starting gas station's index if you can travel around the circuit once in the clockwise direction, otherwise return -1.

Note:

If there exists a solution, it is guaranteed to be unique.
Both input arrays are non-empty and have the same length.
Each element in the input arrays is a non-negative integer.
Example 1:

Input:
gas  = [1,2,3,4,5]
cost = [3,4,5,1,2]

Output: 3

Explanation:
Start at station 3 (index 3) and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
Travel to station 4. Your tank = 4 - 1 + 5 = 8
Travel to station 0. Your tank = 8 - 2 + 1 = 7
Travel to station 1. Your tank = 7 - 3 + 2 = 6
Travel to station 2. Your tank = 6 - 4 + 3 = 5
Travel to station 3. The cost is 5. Your gas is just enough to travel back to station 3.
Therefore, return 3 as the starting index.
Example 2:

Input:
gas  = [2,3,4]
cost = [3,4,3]

Output: -1

Explanation:
You can't start at station 0 or 1, as there is not enough gas to travel to the next station.
Let's start at station 2 and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
Travel to station 0. Your tank = 4 - 3 + 2 = 3
Travel to station 1. Your tank = 3 - 3 + 3 = 3
You cannot travel back to station 2, as it requires 4 unit of gas but you only have 3.
Therefore, you can't travel around the circuit once no matter where you start.
*/
/*
1，总的gas和肯定要大于等于cost的和才能run完cycle
2，gas 和 cost的diff 产生的数组 不能在某个时候和是负数 这个时候就意味着不能走完cycle
*/
func canCompleteCircuit(gas []int, cost []int) int {
	// 1，总的gas和肯定要大于等于cost的和才能run完cycle
	var gasTotal, costTotal int
	for _ , v := range gas {
		gasTotal += v
	}
	for _ , v := range cost {
		costTotal += v
	}
	if gasTotal < costTotal {
		return -1
	}
	// 2，gas 和 cost的diff 产生的数组 不能在某个时候和是负数 这个时候就意味着不能走完cycle
	// 为什么我们不用再验证前面的索引: 因为我们已经在上面保证了 肯定gas sum > cost sum
	// 从当前的索引 total已经是大于0的了，意味着往后我们gas total总是大于0的 然后本来gas sum就是大于cost sum的
	// 意味着即使从前面的索引开始 比如0， gas也一直是有盈余的 所以就不用算了
	total := 0
	res := 0
	for i := 0; i < len(gas); i++ {
		total += gas[i] - cost[i]
		if total < 0 {
			total = 0
			res = i+1
		}
	}
	return res
}
