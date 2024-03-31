import "bytes"

// https://leetcode.com/problems/reverse-vowels-of-a-string/
/*
Write a function that takes a string as input and reverse only the vowels of a string.

Example 1:

Input: "hello"
Output: "holle"
Example 2:

Input: "leetcode"
Output: "leotcede"
Note:
The vowels does not include the letter "y".
*/

func reverseVowels(s string) string {

	start := 0
	end := len(s) - 1
	sByte := []byte(s)
	set := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for start < end {
		f := bytes.IndexByte(set, sByte[start])
		e := bytes.IndexByte(set, sByte[end])
		// string couldn't assign or change it's value
		// it must to convert to byte
		if f >= 0 && e >= 0 {
			sByte[start], sByte[end] = sByte[end], sByte[start]
			start++
			end--
			continue
		}
		if f < 0 {
			start++
		}
		if e < 0 {
			end--
		}
	}

	return string(sByte)
}
