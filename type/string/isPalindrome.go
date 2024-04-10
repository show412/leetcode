/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:51
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-10 22:14:13
 * @Description: file content
 */
import "strings"

// https://leetcode.com/problems/valid-palindrome/
/*
Given a string, determine if it is a palindrome,
considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.

Example 1:

Input: "A man, a plan, a canal: Panama"
Output: true
Example 2:

Input: "race a car"
Output: false
*/

/*
利用golang里自带的一些方法
unicode包
strings.Map
Map returns a copy of the string s with all its characters modified according to the mapping function. 
If mapping returns a negative value, the character is dropped from the string with no replacement.
*/
func isPalindrome(s string) bool {
	f := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return unicode.ToLower(r)
	}
	str := strings.Map(f, s)
	l := 0
	r := len(str)-1
	for l <= r {
		if str[l] != str[r] {
			return false 
		}
		l++
		r--
	}
	return true
 }


func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	start := 0
	end := len(s) - 1
	for start != end && start <= end {
		for start < len(s) && checkCase(s[start]) == false {
			start++
		}
		for end >= 0 && checkCase(s[end]) == false {
			end--
		}
		if start <= end && s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
func checkCase(a byte) bool {
	return (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9')
}
