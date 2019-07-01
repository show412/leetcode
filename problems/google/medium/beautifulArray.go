// https://leetcode.com/problems/beautiful-array/
/*
For some fixed N, an array A is beautiful if it is a permutation of the integers 1, 2, ..., N, such that:

For every i < j, there is no k with i < k < j such that A[k] * 2 = A[i] + A[j].

Given N, return any beautiful array A.  (It is guaranteed that one exists.)



Example 1:

Input: 4
Output: [2,1,4,3]
Example 2:

Input: 5
Output: [3,1,2,5,4]


Note:

1 <= N <= 1000
*/
/*
first thought, DFS get all possible permutation for N, then check it
but the TC is O(n!) so DFS is not the answer

2*A[k] != A[i] + A[j]
if A[i] is odd, A[j] is even, the condition must work
漂亮数组满足以下规律：
奇数和偶数的和肯定满足漂亮数组的条件
漂亮数组每个元素同时加减乘除之后仍然是漂亮数组
两个漂亮数组合并组成的数组仍然是一个漂亮数组
how to recursion?
from 1 start to recursion
首先左面放奇数 右边放偶数 满足漂亮数组条件
然后再次递归res 因为上面的规律下一次循环生成的数组仍然是漂亮数组

那怎么知道tmp最后的个数一定是N呢？
因为tmp是由所有小于等于N的奇数和偶数组成的 所以最后一定是N个
refer to https://leetcode.com/problems/beautiful-array/discuss/186679/Odd-%2B-Even-Pattern-O(N)
*/
func beautifulArray(N int) []int {
	res := make([]int, 1)
	res[0] = 1
	//
	for len(res) < N {
		var tmp []int
		for i := 0; i < len(res); i++ {
			if (2*res[i] - 1) <= N {
				tmp = append(tmp, 2*res[i]-1)
			}
		}
		for i := 0; i < len(res); i++ {
			if 2*res[i] <= N {
				tmp = append(tmp, 2*res[i])
			}
		}
		// 注意这里不用copy也行 因为最后的tmp就是最后的结果
		// 所以不需要copy， 但最好还是copy 因为tmp在运算过程中是变化的
		// cpy := make([]int, len(tmp))
		// copy(cpy, tmp)
		res = tmp
	}
	return res
}
