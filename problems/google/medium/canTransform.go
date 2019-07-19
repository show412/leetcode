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
	"XXXLXXXXXX" "XXXLXXXXXX" true
*/
// not from left to right
// solution https://leetcode.com/problems/swap-adjacent-in-lr-string/solution/
/*
L postion in start must be not smaller than L postion in end, becasue L only move to left
R postion in start must be not bigger than R postion in end, becasue R only move to right
It's the sufficent condition for this problem

Time Complexity: O(N)O(N), where NN is the length of start and end.
Space Complexity: O(1)O(1).
*/
func canTransform(start string, end string) bool {
	if len(start) == 1 {
		if start == end {
			return true
		} else {
			return false
		}
	}
	i := 0
	j := 0
	for i < len(start) && j < len(start) {
		for i < len(start) && start[i] == 'X' {
			i++
		}
		for j < len(start) && end[j] == 'X' {
			j++
		}
		// it means the count of L or count of R in start and end have not the same.
		if (i < len(start) && j >= len(start)) || (i >= len(start) && j < len(start)) {
			return false
		}
		if i < len(start) && j < len(start) {
			if start[i] != end[j] || (start[i] == 'L' && i < j) || (start[i] == 'R' && i > j) {
				return false
			}
		}
		i++
		j++
	}
	return true
}
