import (
	"math"
	"strconv"
)

// https://leetcode.com/problems/generalized-abbreviation/
/*
Write a function to generate the generalized abbreviations of a word.

Note: The order of the output does not matter.

Example:

Input: "word"
Output:
["word", "1ord", "w1rd", "wo1d", "wor1", "2rd", "w2d", "wo2", "1o1d", "1or1", "w1r1", "1o2", "2r1", "3d", "w3", "4"]
*/
/*
数了一下题目中给的例子的所有情况的个数，
是16个，而word有4个字母，刚好是2的4次方，这是巧合吗，当然不是，
后来我又发现如果把0到15的二进制写出来，每一个可以对应一种情况，如下所示：

0000 word
0001 wor1
0010 wo1d
0011 wo2
0100 w1rd
0101 w1r1
0110 w2d
0111 w3
1000 1ord
1001 1or1
1010 1o1d
1011 1o2
1100 2rd
1101 2r1
1110 3d
1111 4

那么我们就可以观察出规律，凡是0的地方都是原来的字母，
单独的1还是1，如果是若干个1连在一起的话，就要求出1的个数，用这个数字来替换对应的字母
*/
func generateAbbreviations(word string) []string {
	res := make([]string, 0)
	for i := 0; i < int(math.Pow(2.0, float64(len(word)))); i++ {
		out := ""
		cnt := 0
		for j := 0; j < len(word); j++ {
			if (i>>uint(j))&1 == 1 {
				cnt++
			} else {
				if cnt > 0 {
					out += strconv.Itoa(cnt)
					cnt = 0
				}
				out += string(word[j])
			}
		}
		if cnt > 0 {
			out += strconv.Itoa(cnt)
		}
		res = append(res, out)
	}
	return res
}

// DFS 方法 不理解 helper(t, i + 1 + to_string(j).size(), res)？
class Solution {
public:
    vector<string> generateAbbreviations(string word) {
        vector<string> res{word};
        helper(word, 0, res);
        return res;
    }
    void helper(string word, int pos, vector<string> &res) {
        for (int i = pos; i < word.size(); ++i) {
            for (int j = 1; i + j <= word.size(); ++j) {
                string t = word.substr(0, i);
                t += to_string(j) + word.substr(i + j);
                res.push_back(t);
                helper(t, i + 1 + to_string(j).size(), res);
            }
        }
    }
};
