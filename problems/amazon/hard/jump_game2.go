// https://leetcode.com/problems/jump-game-ii/
/*
Given an array of non-negative integers,
you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Your goal is to reach the last index in the minimum number of jumps.

Example:

Input: [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2.
    Jump 1 step from index 0 to 1, then 3 steps to the last index.
*/
// 贪心算法
/*
我们这里贪的是一个能到达的最远范围，
我们遍历当前跳跃能到的所有位置，
然后根据该位置上的跳力来预测下一步能跳到的最远距离，
贪出一个最远的范围，一旦当这个范围到达末尾时，
当前所用的步数一定是最小步数。
我们需要两个变量curScope和preScope分别来保存当前的能到达的最远位置和之前能到达的最远位置，
只要curScope未达到最后一个位置则循环继续，preScope先赋值为curScope的值，
表示上一次循环后能到达的最远位置，如果当前位置i小于等于preScope，
说明还是在上一跳能到达的范围内，我们根据当前位置加跳力来更新curScope，
更新cur的方法是比较当前的cur和i + A[i]之中的较大值，
如果题目中未说明是否能到达末尾，我们还可以判断此时pre和cur是否相等，如果相等说明cur没有更新，即无法到达末尾位置
*/
func jump(nums []int) int {
	res := 0
	curScope := 0
	preScope := 0
	for curScope < len(nums)-1 {
		res++
		preScope = curScope
		for i := 0; i <= preScope; i++ {
			curScope = max(curScope, nums[i]+i)
		}
		if curScope == preScope {
			return -1
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}