import "strings"

// https://leetcode.com/problems/repeated-string-match/
/*
Given two strings A and B, find the minimum number of times A has to be repeated such that B is a substring of it. If no such solution, return -1.

For example, with A = "abcd" and B = "cdabcdab".

Return 3, because by repeating A three times (“abcdabcdabcd”), B is a substring of it; and B is not a substring of A repeated two times ("abcdabcd").

Note:
The length of A and B will be between 1 and 10000.
*/
/*
"abcd", "cdabcdab" 3
"a" "aa" 2
"aaaaaaaaaaaaaaaaaaaaaab" "ba" 2
"abcabcabcabc" "abac" -1
*/
// Time Complexity: O(N*(N+M))
// Space complexity: O(M+N)
func repeatedStringMatch(A string, B string) int {
	S := A
	res := 1
	for ; len(S) < len(B); res++ {
		S = S + A
	}
	if strings.Index(S, B) >= 0 {
		return res
	}
	if strings.Index(S+A, B) >= 0 {
		return res + 1
	}
	return -1
}
