import (
	"fmt"
	"strings"
)

// https://www.lintcode.com/problem/query-string/description
/*
Give a binary string str and an integer n to check if the substring of the string contains all binary representations of non-negative integers less than or equal to the given integer.

Example
Example 1:

Input："0110"，n=3
Output："yes"
Explanation：
The substring of str contains "0", "1", "10", "11".
Example 2:

Input："0110"，n=4
Output："no"
Explanation：
The substring of str does not contain "100",
Notice
String length does not exceed 100,000
n does not exceed 100,000
Binary starts at 0 and does not require leading zeros
*/
/**
 * @param str: the string
 * @param n: the n
 * @return: yes or no
 */
func queryString(str string, n int) string {
	// Write your code here.
	if str == "" {
		return "no"
	}
	for i := 0; i <= n; i++ {
		bi := fmt.Sprintf("%b", i)
		// because the Sprintf return a string
		// s := strconv.Itoa(bi)
		if strings.Contains(str, bi) == false {
			return "no"
		}
	}
	return "yes"
}
