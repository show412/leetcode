/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:51
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-03-27 16:13:28
 * @Description: file content
 */
// https://leetcode.com/problems/minimum-window-substring/
/*
Given a string S and a string T,
find the minimum window in S
which will contain all the characters in T in complexity O(n).

Example:

Input: S = "ADOBECODEBANC", T = "ABC"
Output: "BANC"
Note:

If there is no such window in S that covers all characters in T,
return the empty string "".
If there is such window, you are guaranteed
that there will always be only one unique minimum window in S.
*/
/*
s中用滑动窗口
1， 两个hash 去存s滑动窗口的字符串 和 t的字符串
2， 为了防止重复比较s和t 可以用两个数字去比较s和t
3， 先向右滑动 有结果了 再向左 直到r到了s末尾
*/
func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	// hashmap which keeps a count of all the unique characters in t.
	mT := make(map[byte]int, 0)
	for i := 0; i < len(t); i++ {
		mT[t[i]] += 1
	}
	// hashmap record the charcter in s
	mS := make(map[byte]int, 0)
	// two pointer for slide windows
	l, r := 0, 0
	// have matched charcter number
	have, need := 0, len(mT)
	// result here we need to define end to be one value bigger than s length, it can be len(s) or len(s)+1, even math.maxInt64
	start, end := 0, len(s)
	for r < len(s) {
		c := s[r]
		mS[c] += 1
		if _, ok := mT[c]; ok && mT[c] == mS[c] {
			have++
		}
		for have == need {
			if (r - l) < (end-start) {
				start = l
				end = r
			}
			// left pointer move forward
			mS[s[l]]--
			// check if s[l] is in mT and if count is smaller than mT, have decerement by 1
			// why it should be smaller than mT cause of mS[s[l]] possible is bigger than mT[s[l]]
			if _, ok := mT[s[l]]; ok && mS[s[l]] < mT[s[l]]{
				have--
			}
			l++
		}
		r++
	}
    if end == len(s) {
        return ""
    } else {
        return s[start:(end+1)]
    }
}
