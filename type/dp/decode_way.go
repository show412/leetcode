/*
 * @Author: hongwei.sun
 * @Date: 2024-04-01 16:52:58
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-01 16:52:59
 * @Description: file content
 */
import (
	"strconv"
	"strings"
)

// A message containing letters from A-Z is being encoded to numbers using the following mapping:

// 'A' -> 1
// 'B' -> 2
// ...
// 'Z' -> 26
// Given an encoded message containing digits, determine the total number of ways to decode it.

// Example
// Example 1:

// Input: "12"
// Output: 2
// Explanation: It could be decoded as AB (1 2) or L (12).
// Example 2:

// Input: "10"
// Output: 1

// https://www.lintcode.com/problem/decode-ways/description
/**
 * @param s: a string,  encoded message
 * @return: an integer, the number of ways decoding
 */
func numDecodings(s string) int {
	// write your code here
	numArray := strings.Split(s, "")
	// when the first string is 0, it's only one combinations
	if numArray[0] == "0" {
		return 0
	}
	// when only 1 numbers, it's only one combinations
	if len(numArray) == 1 {
		return 1
	}
	// f[i] means how many combinations for the first i numbers
	f := make([]int, len(numArray)+1)
	// 前0个有1种组合方式 比较费解。。。。。
	f[0] = 1
	f[1] = 1
	for i := 2; i < len(numArray)+1; i++ {
		if a, _ := strconv.Atoi(numArray[i-1]); a != 0 {
			// 只看最后一个元素组合 前i个组合和前i-1个的组合是一样多的
			f[i] = f[i-1]
		}
		// 当i和i-1组成数字大于10小于26时，以这两个数做为一个元素，前i个元素的组合数和前i-2个是一样的
		// 再考虑到如果有上面1个数和前i-1的组合情况， 所以总共就是 f[i] = f[i] + f[i-1]
		// 注意这里不是f[i] = f[i-1] + f[i-2] 因为 f[i] 不一定等于 f[i-1] f[i]默认是0
		if b, _ := strconv.Atoi(numArray[i-2] + numArray[i-1]); b >= 10 && b <= 26 {
			f[i] += f[i-2]
		}
	}
	return f[len(numArray)]
}
/*
下面那个方法有overlap， 会算的多 因为+1 or +2 will impact other solution
就是说通过判断加1或者加2  f[i] 会影响f[i+1]的结果 进而导致算法不准确

正确的：
abc => a+(bc) or (ab) + c 
正常情况下 带上a的结果 是和bc数量是一样的  但如果ab<26 意味着 ab可以看成一个整体 它又和c的数量是一样的
这样就不会有overlap了
*/
func numDecodings(s string) int {
    f := make([]int, len(s)+1)
	if s[0] == '0' {
		return 0
	}
    if len(s) == 1 {
		return 1
	}
	if s[len(s)-1] > '0' {
		f[len(s)-1] = 1
	}
    f[len(s)] = 1
	for i := len(s)-2; i >= 0; i-- {

        if s[i] == '0' {
            f[i] = 0
            continue
        } else {
            f[i] += f[i+1]
        }
		if a, _ := strconv.Atoi(s[i:i+2]); a <= 26 {
            f[i] += f[i+2]
        } 
	}
	return f[0]
}

// this is one wrong solution 方法有overlap， 会算的多 因为+1 or +2 will impact other solution
func numDecodings(s string) int {
    f := make([]int, len(s))
	if s[0] == '0' {
		return 0
	}
    if len(s) == 1 {
		return 1
	}
	if s[len(s)-1] > '0' {
		f[len(s)-1] = 1
	}
	for i := len(s)-2; i >= 0; i-- {
		a := s[i]
		b := s[i+1]
		if b >= '0' && b <= '6' {
            if a >= '1' && a <= '2' {
                f[i] = 2 + f[i+1]
            } else if a >= '3' && a <= '9' {
                f[i] = 1 + f[i+1]
            } else {
                f[i] = f[i+1]
            }
        } else {
            if a == '0' {
                f[i] = f[i+1]
            } else if a == '1' {
                f[i] = 2 + f[i+1]
            } else {
                f[i] = 1 + f[i+1]
            }
        }
	}
	return f[0]
}