import "strconv"

// https://leetcode.com/problems/add-to-array-form-of-integer/solution/
/*
For a non-negative integer X,
the array-form of X is an array of its digits in left to right order.
For example, if X = 1231, then the array form is [1,2,3,1].

Given the array-form A of a non-negative integer X,
return the array-form of the integer X+K.



Example 1:

Input: A = [1,2,0,0], K = 34
Output: [1,2,3,4]
Explanation: 1200 + 34 = 1234
Example 2:

Input: A = [2,7,4], K = 181
Output: [4,5,5]
Explanation: 274 + 181 = 455
Example 3:

Input: A = [2,1,5], K = 806
Output: [1,0,2,1]
Explanation: 215 + 806 = 1021
Example 4:

Input: A = [9,9,9,9,9,9,9,9,9,9], K = 1
Output: [1,0,0,0,0,0,0,0,0,0,0]
Explanation: 9999999999 + 1 = 10000000000


Note：

1 <= A.length <= 10000
0 <= A[i] <= 9
0 <= K <= 10000
If A.length > 1, then A[0] != 0
*/
func addToArrayForm(A []int, K int) []int {
	if K == 0 {
		return A
	}
	kStr := strconv.Itoa(K)
	res := make([]int, len(A)+1)
	if len(kStr) > len(A) {
		res = make([]int, len(kStr)+1)
	}
	j := len(A) - 1
	z := len(kStr) - 1
	for i := len(res) - 1; i >= 0; i-- {
		if j >= 0 && z < 0 {
			res[i] = A[j]
		} else if j < 0 && z >= 0 {
			res[i] = int(kStr[z] - '0')
		} else if j >= 0 && z >= 0 {
			res[i] = A[j] + int(kStr[z]-'0')
		}
		j--
		z--
	}
	for i := len(res) - 1; i > 0; i-- {
		res[i-1] = res[i]/10 + res[i-1]
		res[i] = res[i] % 10
	}
	if res[0] == 0 {
		return res[1:]
	}
	return res
}