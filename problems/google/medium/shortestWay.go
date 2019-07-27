// https://leetcode.com/problems/shortest-way-to-form-string/
/*
From any string, we can form a subsequence of that string by deleting some number of characters (possibly no deletions).

Given two strings source and target, return the minimum number of subsequences of source such that their concatenation equals target. If the task is impossible, return -1.



Example 1:

Input: source = "abc", target = "abcbc"
Output: 2
Explanation: The target "abcbc" can be formed by "abc" and "bc", which are subsequences of source "abc".
Example 2:

Input: source = "abc", target = "acdbc"
Output: -1
Explanation: The target string cannot be constructed from the subsequences of source string due to the character "d" in target string.
Example 3:

Input: source = "xyz", target = "xzyxz"
Output: 3
Explanation: The target string can be constructed as follows "xz" + "y" + "xz".


Note:

Both the source and target strings consist of only lowercase English letters from "a"-"z".
The lengths of source and target string are between 1 and 1000.
*/
/*
"abc", "abcbc" 2
"abc", "acdbc" -1
"xyz", "xzyxz" 3
"aaaaa" "aaaaaaaaaaaaa"
*/

func shortestWay(source string, target string) int {
	for i := 0; i < len(target); i++ {
		char := target[i]
		flag := false
		for j := 0; j < len(source); j++ {
			if source[j] == char {
				flag = true
				continue
			}
		}
		if flag == false {
			return -1
		}
	}
	pt := 0
	res := 0
	for pt < len(target) {
		ps := 0
		for ps < len(source) {
			if pt > len(target)-1 {
				break
			}
			if source[ps] != target[pt] {
				ps++
			} else {
				pt++
				ps++
			}
		}
		res++
	}
	return res
}
