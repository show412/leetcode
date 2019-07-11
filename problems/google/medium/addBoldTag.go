import "strings"

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
	res := make([]string, 0)
	closeSlice := make([][2]int, 0)
	tagS := ""
	// boldStart := make(map[int]bool)
	// boldEnd := make(map[int]bool)
	// generate close position according to dict
	for i := 0; i < len(dict); i++ {
		if len(closeSlice) == 0 && strings.Index(s, dict[i]) >= 0 {
			first := [2]int{strings.Index(s, dict[i]), strings.Index(s, dict[i]) + len(dict[i]) - 1}
			closeSlice = append(closeSlice, first)
		}

		for j := 0; j < len(closeSlice); j++ {
			curStart := closeSlice[j][0]
			curEnd := closeSlice[j][1]
			if strings.Index(s, dict[i]) >= 0 {
				start := strings.Index(s, dict[i])
				end := strings.Index(s, dict[i]) + len(dict[i]) - 1
				//check [curStart, curEnd] overlap or sequence to [start, end]
				if (start <= curEnd && end >= curStart) || start == curEnd+1 || end == curStart+1 {
					closeSlice[j] = [2]int{min(start, curStart), max(end, curEnd)}
				} else {
					closeSlice = append(closeSlice, [2]int{start, end})
				}
			}

		}
	}
	// fmt.Println(closeSlice)
	boldStart := make(map[int]bool)
	boldEnd := make(map[int]bool)
	for i := 0; i < len(closeSlice); i++ {
		start := closeSlice[i][0]
		end := closeSlice[i][1]
		boldStart[start] = true
		boldEnd[end] = true
	}

	for i := 0; i < len(s); i++ {
		if boldStart[i] == true {
			res = append(res, "<b>")
			res = append(res, string(s[i]))
		} else if boldEnd[i] == true {
			res = append(res, string(s[i]))
			res = append(res, "</b>")
		} else {
			res = append(res, string(s[i]))
		}
	}

	for i := 0; i < len(res); i++ {
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

/*
class Solution {
    public String boldWords(String S, String[] words) {
        int N = S.length();
        boolean[] mask = new boolean[N];
        for (int i = 0; i < N; ++i)
            for (String word: words) search: {
                for (int k = 0; k < word.length(); ++k)
                    if (k+i >= S.length() || S.charAt(k+i) != word.charAt(k))
                        break search;

                for (int j = i; j < i+word.length(); ++j)
                    mask[j] = true;
            }

        StringBuilder ans = new StringBuilder();
        int anchor = 0;
        for (int i = 0; i < N; ++i) {
            if (mask[i] && (i == 0 || !mask[i-1]))
                ans.append("<b>");
            ans.append(S.charAt(i));
            if (mask[i] && (i == N-1 || !mask[i+1]))
                ans.append("</b>");
        }
        return ans.toString();
    }

    public boolean match(String S, int i, int j, String T) {
        for (int k = i; k < j; ++k)
            if (k >= S.length() || S.charAt(k) != T.charAt(k-i))
                return false;
        return true;
    }
}
*/
