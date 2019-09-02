import "sort"

// https://leetcode.com/problems/group-anagrams/
/*
Given an array of strings, group anagrams together.

Example:

Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
Output:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
Note:

All inputs will be in lowercase.
The order of your output does not matter.
*/

// 定义一个type是[]byte的alias
type strArray []byte

// 都是大写 并且参数i j都是int
func (this strArray) Len() int {
	return len(this)
}

func (this strArray) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this strArray) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	if len(strs) == 0 {
		return res
	}
	m := make(map[string][]string, 0)
	for _, s := range strs {
		arr := []byte(s)
		// 注意这里一定要做强制转换
		sort.Sort(strArray(arr))
		key := string(arr)
		if _, ok := m[key]; !ok {
			m[key] = []string{s}
		} else {
			m[key] = append(m[key], s)
		}
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
