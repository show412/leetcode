/*
 * @Author: hongwei.sun
 * @Date: 2024-02-14 15:15:32
 * @LastEditors: your name
 * @LastEditTime: 2024-02-14 15:40:41
 * @Description: file content
 */
//  A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

//  Given a string s, return true if it is a palindrome, or false otherwise.

//  Example 1:

//  Input: s = "A man, a plan, a canal: Panama"
//  Output: true
//  Explanation: "amanaplanacanalpanama" is a palindrome.
//  Example 2:

//  Input: s = "race a car"
//  Output: false
//  Explanation: "raceacar" is not a palindrome.
//  Example 3:

//  Input: s = " "
//  Output: true
//  Explanation: s is an empty string "" after removing non-alphanumeric characters.
//  Since an empty string reads the same forward and backward, it is a palindrome.

//  Constraints:

//  1 <= s.length <= 2 * 105
//  s consists only of printable ASCII characters.
import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	slower := strings.ToLower(s)
	re := regexp.MustCompile(`[^0-9a-z]`)
	newS := re.ReplaceAllString(slower, "")
	first := 0
	last := len(newS) - 1
	for first < last {
		if newS[first] != newS[last] {
			return false
		}
		first++
		last--
	}
	return true
}