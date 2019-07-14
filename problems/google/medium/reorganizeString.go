// https://leetcode.com/problems/reorganize-string/
/*
Given a string S, check if the letters can be rearranged so that two characters that are adjacent to each other are not the same.

If possible, output any possible result.  If not possible, return the empty string.

Example 1:

Input: S = "aab"
Output: "aba"
Example 2:

Input: S = "aaab"
Output: ""
Note:

S will consist of lowercase letters and have length in range [1, 500].
*/
// 把出现最多的字符放在偶数index， 然后剩下的跳跃着放到奇数位 肯定就是不重的了
// 重要的是重复的数字如果超过 (len(S)+1)/2 就肯定组成不了no duplicate adjacent的结果了
// TC is O(N+P)  SC is O(P) P is types number of character in S
func reorganizeString(S string) string {
	if len(S) == 1 {
		return S
	}
	res := make([]byte, len(S))
	m := make(map[byte]int)
	max := 0
	var maxLetter byte
	for i := 0; i < len(S); i++ {
		m[S[i]] += 1
		if m[S[i]] > max {
			max = m[S[i]]
			maxLetter = S[i]
		}
	}
	if max > (len(S)+1)/2 {
		return ""
	}
	index := 0
	for m[maxLetter] > 0 {
		res[index] = maxLetter
		m[maxLetter]--
		index += 2
	}
	for k, _ := range m {
		for m[k] > 0 {
			if index >= len(S) {
				index = 1
			}
			res[index] = k
			m[k]--
			index += 2
		}
	}
	return string(res)
}

// DFS don't work for this problem, because maybe it cause a TLE
func reorganizeString(S string) string {
	if len(S) == 1 {
		return S
	}
	ary := []byte(S)
	visit := make(map[int]bool, 0)
	entry := make([]byte, 0)
	res := make([][]byte, 0)
	dfs(ary, visit, &entry, &res)
	if len(res) == 0 {
		return ""
	}
	return string(res[0])
}

func dfs(ary []byte, visit map[int]bool, entry *[]byte, res *[][]byte) bool {
	if len(*entry) == len(ary) {
		cpy := make([]byte, len(*entry))
		copy(cpy, *entry)
		*res = append(*res, cpy)
		return true
	}
	for i := 0; i < len(ary); i++ {
		if visit[i] != true {
			if len(*entry) == 0 {
				visit[i] = true
				*entry = append(*entry, ary[i])
				b := dfs(ary, visit, entry, res)
				if b == true {
					break
				}
				visit[i] = false
				*entry = (*entry)[:len(*entry)-1]
				continue
			}
			if len(*entry) > 0 && (*entry)[len(*entry)-1] != ary[i] {
				visit[i] = true
				*entry = append(*entry, ary[i])
				b := dfs(ary, visit, entry, res)
				if b == true {
					break
				}
				visit[i] = false
				*entry = (*entry)[:len(*entry)-1]
			}
		}
	}
	return false
}
