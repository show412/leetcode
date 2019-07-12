package main

import (
	"fmt"
	"strings"
	// "math"
)

func main() {
	/*
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

	// res := addBoldTag("zxioozhirtosxhlkaeociwdsgyqnvxuxizhvfrkestehiiwpmayhlyybhnpgfcxfjltnhlwppzfhxyozlaorrcfdhpbjxwalxtxutnperhobhatxwnvowtolsbsxfwzpyttpksrhhkvtfycpuvztjeeabraqqwustqubzzmjkuzafwcuixdbiuwqgexmlqzlyjxdwjzwjuxopecpcbojkrrceysozvkabzirctyroynqbyomecatnjupiqjexevltvtspnkpkzajzqcagsqnkagrludmpsgczbtkvbbcbjenacgrtowsjxazboxcjdhvjnhcaqsjokztkskntlwpnrecnrfuhuseuhshqoddhutcocgrbvnxgnliiinbjcoydtdlyvfjlmpzegscxcfoacxhwaqusneoptvmfrbljkpsbixsjlbvatkyoillejllsvlqfuvxqaareawgfwwultacnjgepmzqhykqalbqfhxohzfntgiatttqupukfhcvghqjzrutxlwidbfqrssrrnbtbjnsggdvrxrbotvwpofpxrwgcecbcqczuuxlrfpnskznmjdcbqktilxijkilpjwywplxdnirjhgoisfwamuauljoqzbmxlvtzdqocbvrusvxyvslufvbcigggddwlubnjjlxmmmcrcefssuqwvtvmwkhovflsxxraneluwdraknumssfluujdlayviukvaqgumpdefzivsqpjkymhlkltojsbyzmvyusmnuxytexhjcszblvimywfoqwsihtchhqnotxvibjzqhthhojreuuknhbfhzmivqnbofrwktfwlcbajoygobbesqeeiemullkrqwplsmsuvskuifbojrdvtffpqm", []string{"zx", "oo", "hi", "to", "xh", "ka", "oc", "wd", "gy", "nv", "ux", "zh", "fr", "es", "eh", "iw", "ma", "hl", "yb", "np", "fc", "fj", "tn", "lw", "pz", "hx", "oz", "ao", "rc", "dh", "bj"})
	res := addBoldTag("abcxyz123", []string{"abc", "123"})
	fmt.Println(res)
}

func addBoldTag(s string, dict []string) string {
	if len(dict) == 0 {
		return s
	}
	res := make([]string, 0)
	bold := make(map[int]bool)
	end := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(dict); j++ {
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

/*
public class Solution {
    public String addBoldTag(String s, String[] dict) {
        boolean[] bold = new boolean[s.length()];
        for (int i = 0, end = 0; i < s.length(); i++) {
            for (String word : dict) {
                if (s.startsWith(word, i)) {
                    end = Math.max(end, i + word.length());
                }
            }
            bold[i] = end > i;
        }

        StringBuilder result = new StringBuilder();
        for (int i = 0; i < s.length(); i++) {
            if (!bold[i]) {
                result.append(s.charAt(i));
                continue;
            }
            int j = i;
            while (j < s.length() && bold[j]) j++;
            result.append("<b>" + s.substring(i, j) + "</b>");
            i = j - 1;
        }

        return result.toString();
    }
}
*/
