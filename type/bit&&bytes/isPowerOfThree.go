import (
	"regexp"
	"strconv"
)

// https://leetcode.com/problems/power-of-three/
/*
Given an integer, write a function to determine if it is a power of three.

Example 1:

Input: 27
Output: true
Example 2:

Input: 0
Output: false
Example 3:

Input: 9
Output: true
Example 4:

Input: 45
Output: false
Follow up:
Could you do it without using any loop / recursion?
*/
/*
这应该是个套路对其他的方式也适用
有四种解决方案 这是最普通的一种
1， 循环
2， 转成3进制 正则看是不是1开头后面都是0
3， log(n)/log(3)是不是整数 通过%1是不是整数
4， 在计算机中一般32位的 3的幂次的最大值是3^19 = 1162261467 用这个数去%n看是否为0
https://leetcode.com/problems/power-of-three/solution/
*/
// 1 loop
func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	if n < 3 {
		return false
	}
	for n != 0 {
		if n%3 == 0 {
			n = n / 3
		} else {
			break
		}
	}
	if n == 1 {
		return true
	}
	return false
}

// 2
func isPowerOfThree(n int) bool {
	// 转为3进制 return is string
	s3 := strconv.FormatInt(int64(n), 3)
	// reg := regexp.MustCompile("^10*$")
	// return reg.Find([]byte(s3)) != nil
	m, _ := regexp.MatchString("^10*$", s3)
	return m
}

// 4,
func isPowerOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0
}
