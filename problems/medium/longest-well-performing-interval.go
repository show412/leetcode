// https://leetcode.com/problems/longest-well-performing-interval/
/*
We are given hours, a list of the number of hours worked per day
for a given employee.

A day is considered to be a tiring day if and only
if the number of hours worked is (strictly) greater than 8.

A well-performing interval is an interval of days
for which the number of tiring days is strictly larger than
the number of non-tiring days.

Return the length of the longest well-performing interval.



Example 1:

Input: hours = [9,9,6,0,6,6,9]
Output: 3
Explanation: The longest well-performing interval is [9,9,6].


Constraints:

1 <= hours.length <= 10000
0 <= hours[i] <= 16


[6,6,9] 1
[9,9,6] 3
[6,9,9] 3

*/
// common question:
// length of a longest subarray arr[i, j) with sum(arr[i], ... , arr[j-1]) >= K.
// it could have stack and hash to solve
// stack refer to https://leetcode.com/problems/longest-well-performing-interval/discuss/335163/O(N)-Without-Hashmap.-Generalized-ProblemandSolution%3A-Find-Longest-Subarray-With-Sum-greater-K.
func longestWPI(hours []int) int {
	prefixSum := make([]int, len(hours)+1)
	// 前缀和的第0个是0 这里的前缀和的意思是不包括当前的数的前缀和
	prefixSum[0] = 0
	res := 0
	for i := 1; i <= len(hours); i++ {
		if hours[i-1] > 8 {
			prefixSum[i] = prefixSum[i-1] + 1
		} else {
			prefixSum[i] = prefixSum[i-1] - 1
		}
	}
	// so this question is to find out the longest subarray which sum is equal or greater than than 0
	// generate monotonic descresing stack
	stack := make([]int, 0)
	stack = append(stack, prefixSum[0])
	for i := 1; i < len(prefixSum); i++ {
		// there should be bigger not equal, because we need to find max
		if prefixSum[stack[len(stack)-1]] > prefixSum[i] {
			stack = append(stack, i)
		}
	}
	// find the biggest distance
	for i := len(prefixSum) - 1; i >= 0; i-- {
		for len(stack) > 0 && prefixSum[i] > prefixSum[stack[len(stack)-1]] {
			res = max(res, i-stack[len(stack)-1])
			stack = stack[:(len(stack) - 1)]
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

// hash solution
/*
本质: 1. 若到当前位置连续和sum是正数:整个数组就满足条件; 2.
若sum<=0, 那么从头开始整个数组不满足条件;
此时希望找到最前面前缀和为sum-1对应的index i,
那么从i+1开始到当前位置就是longest subarray;此时子数组的元素和也就是1.
sum-x并不需要关心, 因为如果是最长子数组, 子数组的和一定是1, 所以只用关心sum-1.
*/
func longestWPI(hours []int) int {
	res := 0
	// record the score is negative and index
	m := make(map[int]int, 0)
	score := 0
	for i := 0; i < len(hours); i++ {
		if hours[i] > 8 {
			score += 1
		} else {
			score -= 1
		}
		if score > 0 {
			// 这里是 i+1
			res = i + 1
		} else {
			if v, ok := m[score-1]; ok {
				res = max(res, i-v)
			}
			if _, ok := m[score]; !ok {
				m[score] = i
			}
		}
	}
	return res
}