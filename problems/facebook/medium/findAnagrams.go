// https://leetcode.com/problems/find-all-anagrams-in-a-string/
/*
Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.

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

// for big data casef, it's LTE
func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	if len(s) == 0 || len(p) == 0 {
		return res
	}
	pl := len(p) - 1
	mp := make(map[byte]int, 0)
	for k := 0; k < len(p); k++ {
		if _, ok := mp[p[k]]; !ok {
			mp[p[k]] = 1
		} else {
			mp[p[k]]++
		}
	}
	for i := 0; i < len(s)-pl; i++ {
		sub := s[i:(i + pl + 1)]
		flag := true
		// fmt.Println(sub)
		ms := make(map[byte]int, 0)

		for j := 0; j < len(sub); j++ {
			if _, ok := ms[sub[j]]; !ok {
				ms[sub[j]] = 1
			} else {
				ms[sub[j]]++
			}
			if ms[sub[j]] > mp[sub[j]] {
				flag = false
				break
			}
		}

		if flag == true {
			for k, v := range ms {
				if mp[byte(k)] != v {
					flag = false
					break
				}
			}
		}

		if flag == true {
			for k, v := range mp {
				if ms[byte(k)] != v {
					flag = false
				}
			}
		}
		if flag == true {
			res = append(res, i)
		}
	}
	return res
}

// best solution split window
class Solution {
	public:
			vector<int> findAnagrams(string s, string p) {
					if (s.empty()) return {};
					vector<int> res, m(256, 0);
					int left = 0, right = 0, cnt = p.size(), n = s.size();
					for (char c : p) ++m[c];
					while (right < n) {
							if (m[s[right++]]-- >= 1) --cnt;
							if (cnt == 0) res.push_back(left);
							if (right - left == p.size() && m[s[left++]]++ >= 0) ++cnt;
					}
					return res;
			}
	};