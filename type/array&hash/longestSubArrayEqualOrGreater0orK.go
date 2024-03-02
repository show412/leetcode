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
// You may have noticed that almost all O(n) hashmap solution posts are based on the +1, -1 property. And these solutions try to find a longest subarrary with sum > 0.

// Here I propose a more generalized problem and a solution to it.

// Problem:

// * input: arrary arr in which elements are arbitrary integers.
// * output: length of a longest subarray arr[i, j) with sum(arr[i], ... , arr[j-1]) >= K.

// Solution:

// 1. Compute prefix sum of arr as prefixSum so that prefixSum[j] - prefixSum[i] = sum(arr[i], ... arr[j-1])
// 2. Iterate through prefixSum from begin to end and build a strictly monotone decreasing stack smdStack. (smdStack.top() is the smallest)
// 3. Iterate through prefixSum from end to begin. For each prefixSum[i], while smdStack.top() is less than prefixSum[i] by at least K, pop smdStackand try to update result by subarray [index of top,i). Until top element is not less than it by K.
// 4. Return result.

// Time complexity: O(n)

// * step1, compute prefixSum O(n)
// * step2, build strictly monotone decreasing stack O(n)
// * step3, iterate prefixSum O(n). pop top elements in smdStack O(n)

// Explanation: For simplicity, call (i, j) a valid pair if the inequation prefixSum[j] - prefixSum[i] >= K holds. Our goal is to optimize j-i over all valid pair (i, j).

// * Firstly, fix j and minimize i. Consider any i1 and i2 that i1 < i2 < j and prefixSum[i1] <= prefixSum[i2]. It is obvious that (i2, j) can't be a candidate of optimal subarray because (i1, j) is more promising to be valid and longer than (i2, j). Therefor candidates are monotone decreasing, and we can use a strictly monotonic descresing stack to find all candidates. (Monotone stack are just normal stack with a constraint that the elements in them are monotonically increasing or decreasing. When processing a new element, we either pop top elements then push new elem or discard this new element to guarantee the property of monotonicity. If you are not familiar with this data structure, you'd better solve some similar leetcode problems before this one).
// * Secondly, fix i and maximize j. Consider any j1 and j2 that i < j1 < j2 and prefixSum[j2] - prefix[i] >= K. We can find that (i, j1) can't be a candidate of optimal subarrary because (i, j2) is better. This discovery tells us that we should iterate j from end to begin and if we find a valid (i, j), we don't need to keep i in smdStack any longer.

// CPP solution for problem 1124 using monotone stack. (20ms, 11.2MB, faster than 100%)
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
		// 如果用单调递增栈 那在初始化这个栈的时候 因为都是从0开始的 所以初始化的时候可能会丢掉一些 index, 这样后面就找不到正确结果了
		// 所以要用单调递减栈
		if prefixSum[stack[len(stack)-1]] > prefixSum[i] {
			stack = append(stack, i)
		}
	}
	// find the biggest distance
	for i := len(prefixSum) - 1; i >= 0; i-- {
		// here 0 could be K if it's for other question, e.g subsum >= K
		for len(stack) > 0 && prefixSum[i]-prefixSum[stack[len(stack)-1]] > 0 {
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
本质:
1.若到当前位置连续和sum是正数:整个数组就满足条件;
2.若sum<=0, 那么从头开始整个数组不满足条件;
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
			// 这里是 i+1, 因为从头开始都满足
			res = i + 1
		} else {
			// the key is score-1 not score
			// 因为 score 只有两种情况 要不是变得更小 要不就是变得大了1
			// 变得更小在 m 中就不存在, 变得大了1 说明score-1 存在在 m 中 这时候 i-v 就是要求的结果
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