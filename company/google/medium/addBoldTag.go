import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/add-bold-tag-in-string/
/*
Given a string s and a list of strings dict,
you need to add a closed pair of bold tag <b> and </b> to wrap the substrings in s that exist in dict.
If two such substrings overlap,
you need to wrap them together by only one pair of closed bold tag.
Also, if two substrings wrapped by bold tags are consecutive,
you need to combine them.
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

/*
use case:
									Input:
							s = "aaabbcc"
							dict = ["aaa","aab","bc"]
							Output:
							"<b>aaabbc</b>c"

									Input:
							s = "abcxyz123"
							dict = ["abc","123"]
							Output:
							"<b>abc</b>xyz<b>123</b>"

							"aaabbcc"
					["d"]

			"zxioozhirtosxhlkaeociwdsgyqnvxuxizhvfrkestehiiwpmayhlyybhnpgfcxfjltnhlwppzfhxyozlaorrcfdhpbjxwalxtxutnperhobhatxwnvowtolsbsxfwzpyttpksrhhkvtfycpuvztjeeabraqqwustqubzzmjkuzafwcuixdbiuwqgexmlqzlyjxdwjzwjuxopecpcbojkrrceysozvkabzirctyroynqbyomecatnjupiqjexevltvtspnkpkzajzqcagsqnkagrludmpsgczbtkvbbcbjenacgrtowsjxazboxcjdhvjnhcaqsjokztkskntlwpnrecnrfuhuseuhshqoddhutcocgrbvnxgnliiinbjcoydtdlyvfjlmpzegscxcfoacxhwaqusneoptvmfrbljkpsbixsjlbvatkyoillejllsvlqfuvxqaareawgfwwultacnjgepmzqhykqalbqfhxohzfntgiatttqupukfhcvghqjzrutxlwidbfqrssrrnbtbjnsggdvrxrbotvwpofpxrwgcecbcqczuuxlrfpnskznmjdcbqktilxijkilpjwywplxdnirjhgoisfwamuauljoqzbmxlvtzdqocbvrusvxyvslufvbcigggddwlubnjjlxmmmcrcefssuqwvtvmwkhovflsxxraneluwdraknumssfluujdlayviukvaqgumpdefzivsqpjkymhlkltojsbyzmvyusmnuxytexhjcszblvimywfoqwsihtchhqnotxvibjzqhthhojreuuknhbfhzmivqnbofrwktfwlcbajoygobbesqeeiemullkrqwplsmsuvskuifbojrdvtffpqm"
			["zx","oo","hi","to","xh","ka","oc","wd","gy","nv","ux","zh","fr","es","eh","iw","ma","hl","yb","np","fc","fj","tn","lw","pz","hx","oz","ao","rc","dh","bj"]

	"aaabbcc"
	["a","b","c"]

*/
// 思路是哪些字会被包在tag里 哪些就会被标上bold[i]=true

func addBoldTag(s string, dict []string) string {
	if len(dict) == 0 {
		return s
	}
	res := make([]string, 0)
	bold := make(map[int]bool)
	end := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(dict); j++ {
			// this is the key, that means if or not there is a dict[j] is start from i in s
			if (i+len(dict[j]) <= len(s)) && s[i:(i+len(dict[j]))] == dict[j] {
				end = max(end, i+len(dict[j]))
			}
		}
		if end > i {
			bold[i] = true
		}

	}

	for i := 0; i < len(s); i++ {
		if _, ok := bold[i]; !ok {
			res = append(res, string(s[i]))
			continue
		}
		j := i
		// find out which s will be included in s
		for j < len(s) && bold[j] == true {
			j++
		}

		res = append(res, "<b>"+s[i:j]+"</b>")
		i = j - 1
	}
	return strings.Join(res, "")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func addBoldTag(s string, dict []string) string {
	if len(dict) == 0 {
		return s
	}
	res := make([]string, 0)
	closeSlice := make([][]int, 0)
	// tagS := ""
	// boldStart := make(map[int]bool)
	// boldEnd := make(map[int]bool)
	// generate close position according to dict
	for i := 0; i < len(dict); i++ {
		if len(closeSlice) == 0 && strings.Index(s, dict[i]) >= 0 {
			first := []int{strings.Index(s, dict[i]), strings.Index(s, dict[i]) + len(dict[i]) - 1}
			closeSlice = append(closeSlice, first)
			continue
		}
		last := 0
		fmt.Println(len(closeSlice))
		for j := last; j < len(closeSlice); j++ {
			curStart := closeSlice[j][0]
			curEnd := closeSlice[j][1]
			if strings.Index(s, dict[i]) >= 0 {
				start := strings.Index(s, dict[i])
				end := strings.Index(s, dict[i]) + len(dict[i]) - 1
				//check [curStart, curEnd] overlap or sequence to [start, end]
				if (start <= curEnd && end >= curStart) || start == curEnd+1 || end == curStart+1 {
					closeSlice[j] = []int{min(start, curStart), max(end, curEnd)}
					last = j
					break
				} else {
					closeSlice = append(closeSlice, []int{start, end})
					last++
					continue
				}
			}

		}
	}
	// fmt.Println(closeSlice)
	// boldStart := make(map[int]bool)
	// boldEnd := make(map[int]bool)
	// for i := 0; i < len(closeSlice); i++ {
	// 	start := closeSlice[i][0]
	// 	end := closeSlice[i][1]
	// 	boldStart[start] = true
	// 	boldEnd[end] = true
	// }

	for i := 0; i < len(s); i++ {
		flag := 0
		for j := 0; j < len(closeSlice); j++ {
			bold := closeSlice[j]
			if bold[0] == i {
				res = append(res, "<b>")
				res = append(res, string(s[i]))
				flag = 1
				break
			}
			if bold[1] == i {
				res = append(res, string(s[i]))
				res = append(res, "</b>")
				flag = 1
				break
			}
		}
		if flag == 0 {
			res = append(res, string(s[i]))
		}

		// if boldStart[i] == true {
		// 	res = append(res, "<b>")
		// 	res = append(res, string(s[i]))
		// } else if boldEnd[i] == true {
		// 	res = append(res, string(s[i]))
		// 	res = append(res, "</b>")
		// } else {
		// 	res = append(res, string(s[i]))
		// }
	}

	// for i := 0; i < len(res); i++ {
	// 	if res[i] == "" {
	// 		break
	// 	}
	// 	tagS += res[i]
	// }
	return strings.Join(res, "")
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

