/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:52
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-03 10:32:35
 * @Description: file content
 */
// https://leetcode.com/problems/daily-temperatures/
/*
Given a list of daily temperatures T, return a list such that,
for each day in the input,
tells you how many days you would have to wait until a warmer temperature.
If there is no future day for which this is possible, put 0 instead.

For example, given the list of temperatures
T = [73, 74, 75, 71, 69, 72, 76, 73],
your output should be [1, 1, 4, 2, 1, 1, 0, 0].

Note: The length of temperatures will be in the range [1, 30000].
Each temperature will be an integer in the range [30, 100].

*/

/*
维护一个最大栈. 存的是index, 当后面来的有比栈最后一个大的时候，说明是最近的大于它的值。
根据栈里的index 更新ret即可
*/
func dailyTemperatures(temperatures []int) []int {

	l := len(temperatures)

	ret := make([]int, l)

	stack := make([]int, 0)

	for i := 0; i < l; i++ {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			ret[stack[len(stack)-1]] = (i - stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ret
}



/*
Solutions:
	The brute force could solve the problem but it has low performance with TC and SC
	TC is O(n^2)
*/
func dailyTemperatures(T []int) []int {
	if len(T) == 1 {
		return []int{0}
	}
	var res []int
	for i := 0; i < len(T)-1; i++ {
		for j := i + 1; j < len(T); j++ {
			if T[j] > T[i] {
				res = append(res, j-i)
				break
			}
			if j == len(T)-1 {
				res = append(res, 0)
			}
		}
	}
	res = append(res, 0)
	return res
}

/*
Use stack and a struct make the TC better
TC is O(n) SC is O(n)
*/
func dailyTemperatures(T []int) []int {
	if len(T) == 1 {
		return []int{0}
	}
	var res []int
	type node struct {
		value  int
		index  int
		bigger int
	}
	var stack []*node
	var nodeList []*node
	for i := 0; i < len(T); i++ {
		item := node{value: T[i], index: i, bigger: 0}
		nodeList = append(nodeList, &item)
	}
	stack = append(stack, nodeList[0])
	for i := 1; i < len(nodeList); i++ {
		for len(stack) != 0 && nodeList[i].value > stack[len(stack)-1].value {
			last := stack[len(stack)-1]
			last.bigger = i - last.index
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nodeList[i])
	}
	for i := 0; i < len(nodeList); i++ {
		res = append(res, nodeList[i].bigger)
	}
	return res
}


