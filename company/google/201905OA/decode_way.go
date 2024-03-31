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

