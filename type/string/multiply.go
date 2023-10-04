// https://leetcode.com/problems/multiply-strings/
/*
Given two non-negative integers num1 and num2 represented as strings,
return the product of num1 and num2, also represented as a string.

Example 1:

Input: num1 = "2", num2 = "3"
Output: "6"

Example 2:

Input: num1 = "123", num2 = "456"
Output: "56088"

Note:

The length of both num1 and num2 is < 110.
Both num1 and num2 contain only digits 0-9.
Both num1 and num2 do not contain any leading zero, except the number 0 itself.
You must not use any built-in BigInteger library or convert the inputs to integer directly.

*/
// string refer to https://www.cnblogs.com/grandyang/p/4395356.html
/*
num1 和 num2 中任意位置的两个数字相乘，得到的两位数在最终结果中的位置是确定的，
比如 num1 中位置为i的数字乘以 num2 中位置为j的数字，那么得到的两位数字的位置为 i+j 和 i+j+1
*/
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	num1Array := []byte(num1)
	num2Array := []byte(num2)
	l1 := len(num1Array)
	l2 := len(num2Array)
	l := l1 + l2
	res := make([]int, l)
	str := ""
	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			mul := (num1[i] - '0') * (num2[j] - '0')
			p1 := i + j
			p2 := i + j + 1
			// next loop, current p2 is p1 of next loop
			sum := int(mul) + res[p2]
			res[p1] += sum / 10
			res[p2] = sum % 10
		}
	}
	for _, v := range res {
		if str != "" || v != 0 {
			str += string(byte(v + '0'))
		}
	}
	return str
}