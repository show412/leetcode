/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:52
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-03 12:01:19
 * @Description: file content
 */
// https://leetcode.com/problems/car-fleet/
/*
N cars are going to the same destination along a one lane road.  The destination is target miles away.

Each car i has a constant speed speed[i] (in miles per hour), and initial position position[i] miles towards the target along the road.

A car can never pass another car ahead of it, but it can catch up to it, and drive bumper to bumper at the same speed.

The distance between these two cars is ignored - they are assumed to have the same position.

A car fleet is some non-empty set of cars driving at the same position and same speed.  Note that a single car is also a car fleet.

If a car catches up to a car fleet right at the destination point, it will still be considered as one car fleet.


How many car fleets will arrive at the destination?



Example 1:

Input: target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]
Output: 3
Explanation:
The cars starting at 10 and 8 become a fleet, meeting each other at 12.
The car starting at 0 doesn't catch up to any other car, so it is a fleet by itself.
The cars starting at 5 and 3 become a fleet, meeting each other at 6.
Note that no other cars meet these fleets before the destination, so the answer is 3.

Note:

0 <= N <= 10 ^ 4
0 < target <= 10 ^ 6
0 < speed[i] <= 10 ^ 6
0 <= position[i] < target
All initial positions are different.
*/
/*
准备一个二维数组{position, speed}按位置排序， 
如果前面的车(index在前)到target的时间比后面的车到target的少，说明前面的车在某个位置能赶上后面的车合成一个车队
然后以后面的速度继续。否则说明永远也赶不上后面的车.
暗藏的条件是后面的车速度肯定小于前面，否则不可能赶上。
所以要从后往前遍历，因为不管前面的车合没合成一个车队，都会以后面的车速度前进(因为暗藏的条件)。
用最大stack去存到达各个车target的time，如果后面到target的time大于stack最后一个值 说明永远也合不成车队
结果就是stack的length
*/
func carFleet(target int, position []int, speed []int) int {
	pair := make([][2]int, 0)
	// push position and speed into one pair
	for i := 0; i < len(position); i++ {
		pair = append(pair, [2]int{position[i], speed[i]})
	}
	// sort pair by position
	sort.Slice(pair, func(i, j int) bool {
		return pair[i][0] < pair[j][0]
	})
	stack := make([]float64, 0)
	// range pair from end
	for i := len(pair) - 1; i >=0; i-- {
		pos := pair[i][0]
		sp := pair[i][1]
		time := float64(target-pos)/float64(sp)
		// 因为是从后往前遍历，所以栈的最后一个值是后面的到target的时间
		// 前面的time大于它说明永远也追不上后面的，所以是单独的一个值
		if len(stack) == 0 || time > stack[len(stack)-1] {
			stack = append(stack, time)
		}
	}
	return len(stack)
}
