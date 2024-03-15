/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:52
 * @LastEditors: your name
 * @LastEditTime: 2024-03-15 23:23:31
 * @Description: file content
 */
// https://leetcode.com/problems/encode-and-decode-strings/
/*
Design an algorithm to encode a list of strings to a string. The encoded string is then sent over the network and is decoded back to the original list of strings.

Machine 1 (sender) has the function:

string encode(vector<string> strs) {
  // ... your code
  return encoded_string;
}
Machine 2 (receiver) has the function:
vector<string> decode(string s) {
  //... your code
  return strs;
}
So Machine 1 does:

string encoded_string = encode(strs);
and Machine 2 does:

vector<string> strs2 = decode(encoded_string);
strs2 in Machine 2 should be the same as strs in Machine 1.

Implement the encode and decode methods.



Note:

The string may contain any possible characters out of 256 valid ascii characters. Your algorithm should be generalized enough to work on any possible characters.
Do not use class member/global/static variables to store states. Your encode and decode algorithms should be stateless.
Do not rely on any library method such as eval or serialize methods. You should implement your own encode/decode algorithm.
*/
/*
this question is about how to make delimiter between string
we can define encode is like  {len(string)}#string
we can define decode to parse this encoded string to work out []string

this question is about int<->string<->rune convert
*/
// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := Encode([]string{"fda", "12dc#y", "pesm&", "mvn3ws"})
	fmt.Println(s)
	sd := Decode(s)
	fmt.Println(sd)
}

func Encode(strs []string) string {
	var res string
	for _, str := range strs {
		tmp := strconv.Itoa(len(str)) + "#" + str
		res += tmp
	}
	return res
}

func Decode(strs string) []string {
	res := make([]string, 0)
	runeArray := []rune(strs)
	for i := 1; i < len(runeArray); i++ {
		l := 0
		s := string(runeArray[i])
		if s == "#" {
			l = int(runeArray[i-1] - '0')
			res = append(res, string(runeArray[(i+1):(i+1+l)]))
			i += l
		}
	}
	return res
}
