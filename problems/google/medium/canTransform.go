// https://leetcode.com/problems/swap-adjacent-in-lr-string/
/*
In a string composed of 'L', 'R', and 'X' characters, like "RXXLRXRXL", a move consists of either replacing one occurrence of "XL" with "LX", or replacing one occurrence of "RX" with "XR". Given the starting string start and the ending string end, return True if and only if there exists a sequence of moves to transform one string to the other.

Example:

Input: start = "RXXLRXRXL", end = "XRLXXRRLX"
Output: True
Explanation:
We can transform start to end following these steps:
RXXLRXRXL ->
XRXLRXRXL ->
XRLXRXRXL ->
XRLXXRRXL ->
XRLXXRRLX
Note:

1 <= len(start) = len(end) <= 10000.
Both start and end will only consist of characters in {'L', 'R', 'X'}.
*/
/*
test case:
	start = "RXXLRXRXL", end = "XRLXXRRLX" true
	"XLLR"  "LXLX"  false
	"XXXXXLXXXX" "LXXXXXXXXX" true
*/
func canTransform(start string, end string) bool {
	if len(start) == 1 {
		if start == end {
			return true
		} else {
			return false
		}
	}
	m := map[string]string{"XL": "LX", "RX": "XR"}
	for i := 0; i < len(start)-1; i++ {
		if start[i] == end[i] {
			continue
		} else {
			s := string(start[i]) + string(start[i+1])
			e := string(end[i]) + string(end[i+1])
			if v, _ := m[s]; v == e {
				i++
			} else {
				return false
			}
		}
	}
	return true
}
