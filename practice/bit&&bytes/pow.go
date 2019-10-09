// https://leetcode.com/problems/powx-n/
/*
Implement pow(x, n), which calculates x raised to the power n (xn).

Example 1:

Input: 2.00000, 10
Output: 1024.00000
Example 2:

Input: 2.10000, 3
Output: 9.26100
Example 3:

Input: 2.00000, -2
Output: 0.25000
Explanation: 2-2 = 1/22 = 1/4 = 0.25
Note:

-100.0 < x < 100.0
n is a 32-bit signed integer, within the range [−231, 231 − 1]
*/
// TC is O(n) and SC is O(1)
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.00000
	}
	if x == 1.00000 {
		return x
	}
	if x == -1.00000 {
		if n%2 == 0 {
			return 1.00000
		} else {
			return -1.00000
		}
	}

	res := 1.00
	if n < 0 {
		x = 1 / x
		n = -n
	}
	for n > 0 {
		res *= x
		n--
	}
	return float64(res)
}

// last soultion is brute force to solve this question
// there are one better solution with better TC and SC
// TC is O(logn) and SC is O(1)
// refer to https://blog.csdn.net/Lison_Zhu/article/details/80031375
// because the x^n = x ^ (∑2^i) = 某些2进制相加总会等于n
// 这样就可以把n变为2进制 看哪位是1 根据2进制计算就会减少时间复杂度
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.00000
	}
	if n == 1 {
		return x
	}
	if x == 1.00000 {
		return x
	}
	if x == -1.00000 {
		if n%2 == 0 {
			return 1.00000
		} else {
			return -1.00000
		}
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	res := 1.00000
	for n != 0 {
		if n&1 == 1 {
			res *= x
		}
		// 因为转为了2进制
		// 这步的目的是往前挪动2的幂次 比如第一次循环是x 下一次就是x^2 再下一次是x^4
		x *= x
		n = n >> 1
	}
	return res
}
