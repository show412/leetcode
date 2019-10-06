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
滑动窗口
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
	tNum := len(mT)
	// hashmap record the t charcter in s
	mS := make(map[byte]int, 0)
	// two pointer
	l, r := 0, 0
	// have matched charcter number
	formerNum := 0
	// flag means if or not there is a matched
	flag := false
	// s, e means the matched start and end in s
	start, end := 0, len(s)-1

	for r < len(s) {
		c := s[r]
		mS[c] += 1
		if _, ok := mT[c]; ok && mT[c] == mS[c] {
			formerNum++
		}

		for l <= r && formerNum == tNum {
			// when the window is narrow, update the start and end
			if (r - l) <= (end - start) {
				flag = true
				start = l
				end = r
			}
			b := s[l]
			mS[b] -= 1
			if _, ok := mT[b]; ok && mS[b] < mT[b] {
				formerNum--
			}
			l++
		}
		r++
	}
	if flag == true {
		return s[start:(end + 1)]
	} else {
		return ""
	}
}
