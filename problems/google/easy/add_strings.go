// https://leetcode.com/problems/add-strings/
/*
Given two non-negative integers num1 and num2 represented as string, return the sum of num1 and num2.

Note:

The length of both num1 and num2 is < 5100.
Both num1 and num2 contains only digits 0-9.
Both num1 and num2 does not contain any leading zero.
You must not use any built-in BigInteger library or convert the inputs to integer directly.
*/
// the ideas is we could use byte and '0' to convert to int
// and + '0' at last to convert to string
// and carry we could use divide 10 and mod 10
func addStrings(num1 string, num2 string) string {
	// the result length should be the longest between num1 and num2 to add 1
	// because it maybe have a carry
	res := make([]byte, len(num1)+1)
	if len(num1) < len(num2) {
		res = make([]byte, len(num2)+1)
	}
	// from end of num1 and num2 to add by byte
	// if one of num1 or num2 has no digits, add the left disgits into res
	j, z := len(num1)-1, len(num2)-1
	for i := len(res) - 1; i >= 0; i-- {
		if j >= 0 && z < 0 {
			res[i] = num1[j] - '0'
		} else if j < 0 && z >= 0 {
			res[i] = num2[z] - '0'
		} else if j >= 0 && z >= 0 {
			res[i] = (num1[j] - '0') + (num2[z] - '0')
		}
		j--
		z--
	}

	// Deal with the carry
	for i := len(res) - 1; i > 0; i-- {
		res[i-1] = res[i]/10 + res[i-1]
		res[i] = res[i] % 10
	}
	// work out the result
	var str []byte
	for i := 0; i < len(res); i++ {
		if i == 0 && res[i] == 0 {
			continue
		}
		str = append(str, res[i]+'0')
	}
	if len(str) == 0 {
		return "0"
	}
	return string(str)
}
