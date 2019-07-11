// https://leetcode.com/problems/add-bold-tag-in-string/
/*
Given a string s and a list of strings dict, you need to add a closed pair of bold tag <b> and </b> to wrap the substrings in s that exist in dict. If two such substrings overlap, you need to wrap them together by only one pair of closed bold tag. Also, if two substrings wrapped by bold tags are consecutive, you need to combine them.
Example 1:
Input:
s = "abcxyz123"
dict = ["abc","123"]
Output:
"<b>abc</b>xyz<b>123</b>"
Example 2:
Input:
s = "aaabbcc"
dict = ["aaa","aab","bc"]
Output:
"<b>aaabbc</b>c"
Note:
The given dict won't contain duplicates, and its length won't exceed 100.
All the strings in input have length in range [1, 1000].
*/
func addBoldTag(s string, dict []string) string {
	if len(dict) == 0 {
		return s
	}
	res := make([]string, 3*len(s))
	closeSlice := make([][2]int, 0)
	tagS := ""
	first := [2]{strings.Index(s, dict[0]), strings.Index(s, dict[0])+len(dict[0])}
	closeSlice = append(closeSlice, first)
	// generate close position according to dict
	for i :=1; i < len(dict); i++ {
		for j :=0; j<len(closeSlice); j++ {
			curStart := closeSlice[j][0]
			curEnd := closeSlice[j][1]
			start := strings.Index(s, dict[i])
			end := strings.Index(s, dict[i])+len(dict[i])
			//check [curStart, curEnd] overlap or sequence to [start, end]
			if (start <= curEnd && end >= curStart  ) || start == curEnd+1 || end == curStart+1 {
				closeSlice[j] = [2]int{min(start, curStart), max(end, curEnd)}
			} else {
				closeSlice = append(closeSlice, [2]int{start, end})
			}
		}
	}
	for i :=0; i<len(closeSlice); i++ {
		start := closeSlice[i][0]
		end := closeSlice[i][1]
		res[start] = "<b>"
		res[end] = "</b>"
	}
	for i :=0; i<len(s); i++ {
		if res[i] == "" {
			res[i] = string(s[i])
		} else {
			res[i+1] = string(s[i])
		}
	}
	for i :=0; i<len(res); i++ {
		if res[i] == "" {
			break
		}
		tagS += res[i]
	}
	return tagS
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}