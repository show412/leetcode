/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:51
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-03 15:31:57
 * @Description: file content
 */
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
/*
map的key是一个定长数组，这样就能比较了
value是[]string 
*/

func groupAnagrams(strs []string) [][]string {
    m := make(map[[26]int][]string, 0)
	res := make([][]string, 0)
	for _, str := range strs {
		array := getAnagrams(str)
		
			m[array] = append(m[array], str)
		
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
 }

 func getAnagrams(str string) [26]int {
	var a [26]int
	for _, s := range str {
		a[s - 'a']++
	}
	return a
 }



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
