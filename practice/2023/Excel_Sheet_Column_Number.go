/*
 * @Author: hongwei.sun
 * @Date: 2023-10-04 16:13:12
 * @LastEditors: your name
 * @LastEditTime: 2023-10-04 16:59:43
 * @Description: file content
 */

//  https://leetcode.com/problems/excel-sheet-column-number/description/
//  it's one 26 system convert question
func titleToNumber(columnTitle string) int {
	res := 0
	scale := 1
	for i := len(columnTitle) - 1; i >= 0; i -= 1 {
		res += (int(columnTitle[i]) - int('A') + 1) * scale
		scale *= 26
	}
	return res
}
