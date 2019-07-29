// https://leetcode.com/problems/expressive-words/
/*
Sometimes people repeat letters to represent extra feeling,
such as "hello" -> "heeellooo", "hi" -> "hiiii".
In these strings like "heeellooo",
we have groups of adjacent letters that are all the same:  "h", "eee", "ll", "ooo".

For some given string S,
a query word is stretchy if it can be made to be equal to S
by any number of applications of the following extension operation:
choose a group consisting of characters c,
and add some number of characters c to the group
so that the size of the group is 3 or more.

For example, starting with "hello",
we could do an extension on the group "o" to get "hellooo",
but we cannot get "helloo" since the group "oo" has size less than 3.
Also, we could do another extension like "ll" -> "lllll" to get "helllllooo".
If S = "helllllooo", then the query word "hello" would be stretchy
because of these two extension operations:
query = "hello" -> "hellooo" -> "helllllooo" = S.

Given a list of query words, return the number of words that are stretchy.



Example:
Input:
S = "heeellooo"
words = ["hello", "hi", "helo"]
Output: 1
Explanation:
We can extend "e" and "o" in the word "hello" to get "heeellooo".
We can't extend "helo" to get "heeellooo"
because the group "ll" is not size 3 or more.


Notes:

0 <= len(S) <= 100.
0 <= len(words) <= 100.
0 <= len(words[i]) <= 100.
S and all words in words consist only of lowercase letters
*/

// test cases:
/*
	"heeellooo"
	["axxxrrzzz"]
*/
/*
	"zzzzzyyyyy"
	["zzyy","zy","zyy"]
*/
/*
	"heeellooo"
	"hello", "hi", "helo"
*/
/*
		"dddiiiinnssssssoooo"
	  ["dinnssoo","ddinso","ddiinnso","ddiinnssoo","ddiinso","dinsoo","ddiinsso","dinssoo","dinso"]
*/
// res := expressiveWords("heeellooo", []string{"axxxrrzzz"})

func expressiveWords(S string, words []string) int {
	if len(words) == 0 {
		return 0
	}
	res := 0
	SByteArray, SByteNum := generatePair(S)
	for i := 0; i < len(words); i++ {
		word := words[i]
		wordByteArray, wordByteNum := generatePair(word)

		if sliceEqual(SByteArray, wordByteArray) == true {
			for j := 0; j < len(wordByteNum); j++ {
				if (SByteNum[j] < 3 && SByteNum[j] != wordByteNum[j]) || SByteNum[j] < wordByteNum[j] {
					break
				}

				if j == len(wordByteNum)-1 {
					res++
				}

			}
		}
	}
	return res
}

func generatePair(s string) ([]byte, []int) {
	var byteArray []byte
	var byteNum []int
	byteArray = append(byteArray, s[0])
	num := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			num++
		} else {
			byteArray = append(byteArray, s[i])
			byteNum = append(byteNum, num)
			num = 1
		}
		if i == len(s)-1 {
			byteNum = append(byteNum, num)
		}
	}
	return byteArray, byteNum
}

func sliceEqual(s1, s2 []byte) bool {
	if len(s1) != len(s2) {
		return false
	}

	if (s1 == nil) != (s2 == nil) {
		return false
	}

	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}

	return true
}
