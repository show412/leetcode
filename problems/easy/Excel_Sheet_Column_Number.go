import "math"

// # Given a column title as appear in an Excel sheet, return its corresponding column number.
// #
// # For example:
// #
// #     A -> 1
// #     B -> 2
// #     C -> 3
// #     ...
// #     Z -> 26
// #     AA -> 27
// #     AB -> 28
// #     ...
// # Example 1:
// #
// # Input: "A"
// # Output: 1
// # Example 2:
// #
// # Input: "AB"
// # Output: 28
// # Example 3:
// #
// # Input: "ZY"
// # Output: 701
// # https://leetcode.com/problems/excel-sheet-column-number/description/

// # @param {String} s
// # @return {Integer}
func titleToNumber(s string) int {
	if s == "" {
		return 0
	}
	arr := []rune(s)
	length := len(arr)
	res := 0
	base := int('A') - 1
	for i := 0; i < length; i++ {
		charNum := int(arr[i]) - base
		if length >= 2 {
			res += charNum * int(math.Pow(float64(26), float64(length-1-i)))
		} else {
			res = charNum
		}
	}
	return res
}