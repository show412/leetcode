// https://leetcode.com/problems/backspace-string-compare/
/*
Given two strings S and T, return if they are equal when both are typed into empty text editors. # means a backspace character.

Example 1:

Input: S = "ab#c", T = "ad#c"
Output: true
Explanation: Both S and T become "ac".
Example 2:

Input: S = "ab##", T = "c#d#"
Output: true
Explanation: Both S and T become "".
Example 3:

Input: S = "a##c", T = "#a#c"
Output: true
Explanation: Both S and T become "c".
Example 4:

Input: S = "a#c", T = "b"
Output: false
Explanation: S becomes "c" while T becomes "b".
Note:

1 <= S.length <= 200
1 <= T.length <= 200
S and T only contain lowercase letters and '#' characters.
Follow up:

Can you solve it in O(N) time and O(1) space?
*/

func backspaceCompare(S string, T string) bool {
	if len(S) == 0 && len(T) == 0 {
		return true
	}
	start1 := len(S) - 1
	start2 := len(T) - 1
	jump1 := 0
	jump2 := 0
	for start1 >= 0 || start2 >= 0 {
		for start1 >= 0 && S[start1] == '#' {
			jump1++
			start1--
		}
		for start2 >= 0 && T[start2] == '#' {
			jump2++
			start2--
		}
		for jump1 > 0 && start1 >= 0 {

			if start1 >= 0 && S[start1] != '#' {
				start1--
				jump1--
			}
			if start1 >= 0 && S[start1] == '#' {
				start1--
				jump1++
			}
		}
		for jump2 > 0 && start2 >= 0 {

			if start2 >= 0 && T[start2] != '#' {
				start2--
				jump2--
			}

			if start2 >= 0 && T[start2] == '#' {
				start2--
				jump2++
			}
		}
		if start1 < 0 && start2 < 0 {
			return true
		}
		// maybe in the before of S or T, there is items could be removed
		// so it can't use S[start1] == T[start2] here
		if start1 >= 0 && start2 >= 0 && S[start1] != T[start2] {
			return false
		}
		if (start1 < 0 && start2 >= 0) || (start1 >= 0 && start2 < 0) {
			return false
		}
		start1--
		start2--
	}
	return true
}
