// https://leetcode.com/problems/strobogrammatic-number-ii/
/*
A strobogrammatic number is a number that looks the same when rotated 180 degrees (looked at upside down).

Find all strobogrammatic numbers that are of length = n.

Example:

Input:  n = 2
Output: ["11","69","88","96"]
*/
func findStrobogrammatic(n int) []string {
	if n == 1 {
		return []string{"0", "1", "8"}
	}
	if n == 2 {
		return []string{"11", "69", "88", "96"}
	}
	res := make([]string, 0)
	for _, v := range helper(n-2, res) {
		res = append(res, "1"+v+"1")
		res = append(res, "8"+v+"8")
		res = append(res, "6"+v+"9")
		res = append(res, "9"+v+"6")
	}
	return res
}
func helper(n int, res []string) []string {
	if n == 0 {
		return []string{""}
	}
	if n == 1 {
		return []string{"0", "1", "8"}
	}
	// res := make([]string, 0)
	for _, v := range helper(n-2, res) {
		res = append(res, "0"+v+"0")
		res = append(res, "1"+v+"1")
		res = append(res, "8"+v+"8")
		res = append(res, "6"+v+"9")
		res = append(res, "9"+v+"6")
	}
	return res
}
