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

// typical slide window for string
// 类似一种伸缩窗口
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	sa := make([]int, 26)
	pa := make([]int, 26)
	res := make([]int, 0)
	for i := 0; i < len(p); i++ {
		pa[p[i]-'a']++
	}
	left := 0
	right := 0
	// 这是一种典型的滑动伸缩窗口的写法
	for left < len(s) && right < len(s) {
		sa[s[right]-'a']++
		for left <= right && sa[s[right]-'a'] > pa[s[right]-'a'] {
			sa[s[left]-'a']--
			left++
		}
		if right-left+1 == len(p) {
			res = append(res, left)
			sa[s[left]-'a']--
			left++
		}
		right++
	}
	return res
}

// best solution split window for TC and SC
// 不是很好理解 但是其实和上面的是一个意思 https://www.cnblogs.com/grandyang/p/6014408.html
func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	if len(s) == 0 || len(p) == 0 {
		return res
	}
	m := make(map[byte]int, 0)
	left := 0
	right := 0
	cnt := len(p)
	n := len(s)
	for i := 0; i < len(p); i++ {
		m[p[i]]++
	}
	for right < n {
		if v, ok := m[s[right]]; ok {
			if v >= 1 {
				cnt--
			}
			m[s[right]]--
		}
		right++
		if cnt == 0 {
			res = append(res, left)
		}
		/*
			there could use this kind of code, when right minus left is len(p),
			means that the slide window should move forwards
			if (right - left) == len(p) {
		*/
		if right >= len(p) {
			if v, ok := m[s[left]]; ok {
				if v >= 0 {
					cnt++
				}
				m[s[left]]++
			}
			left++
		}
	}
	return res
}
