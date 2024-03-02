// https://leetcode.com/problems/find-all-anagrams-in-a-string/
/*
Given a string s and a non-empty string p,
find all the start indices of p's anagrams in s.

Strings consists of lowercase English letters only and the length of both strings s and p will not be larger than 20,100.

The order of output does not matter.

Example 1:

Input:
s: "cbaebabacd" p: "abc"

Output:
[0, 6]

Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".
Example 2:

Input:
s: "abab" p: "ab"

Output:
[0, 1, 2]

Explanation:
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".
*/

/*
test case:
"cbaebabacd", "abc"
	[0, 6]

	Input:
s: "abab" p: "ab"

Output:
[0, 1, 2]

Input
"ababababab"
"aab"
Output
[0,1,2,3,4,5,6,7]
Expected
[0,2,4,6]

https://leetcode.com/submissions/detail/253837170/testcase/
*/
// 小写字母注意可以用数组来记录, 只要不一样肯定有一个会小于0
// 这是字母里一种很常用的方法 利用数组
func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	if len(s) == 0 || len(p) == 0 {
		return res
	}
	pl := len(p) - 1
	cnt := make([]int, 128)
	for k := 0; k < len(p); k++ {
		cnt[p[k]]++
	}
	for i := 0; i < len(s)-pl; i++ {
		sub := s[i:(i + pl + 1)]
		flag := true
		tmp := make([]int, 128)
		copy(tmp, cnt)
		for j := 0; j < len(sub); j++ {
			// this kind of way is one common method to check anagrams
			tmp[sub[j]]--
			if tmp[sub[j]] < 0 {
				flag = false
				break
			}
		}

		if flag == true {
			res = append(res, i)
		}
	}
	return res
}
