import (
	"strings"
	"unicode"
)

// https://leetcode.com/problems/most-common-word/
/*
Given a paragraph and a list of banned words, return the most frequent word that is not in the list of banned words.  It is guaranteed there is at least one word that isn't banned, and that the answer is unique.

Words in the list of banned words are given in lowercase, and free of punctuation.  Words in the paragraph are not case sensitive.  The answer is in lowercase.



Example:

Input:
paragraph = "Bob hit a ball, the hit BALL flew far after it was hit."
banned = ["hit"]
Output: "ball"
Explanation:
"hit" occurs 3 times, but it is a banned word.
"ball" occurs twice (and no other word does), so it is the most frequent non-banned word in the paragraph.
Note that words in the paragraph are not case sensitive,
that punctuation is ignored (even if adjacent to words, such as "ball,"),
and that "hit" isn't the answer even though it occurs more because it is banned.


Note:

1 <= paragraph.length <= 1000.
0 <= banned.length <= 100.
1 <= banned[i].length <= 10.
The answer is unique, and written in lowercase (even if its occurrences in paragraph may have uppercase symbols, and even if it is a proper noun.)
paragraph only consists of letters, spaces, or the punctuation symbols !?',;.
There are no hyphens or hyphenated words.
Words only consist of letters, never apostrophes or other punctuation symbols.
*/

func mostCommonWord(paragraph string, banned []string) string {
	bannedMap := make(map[string]bool, 0)
	for _, v := range banned {
		bannedMap[v] = true
	}
	pureText := make([]byte, 0)
	for i := 0; i < len(paragraph); i++ {
		letter := paragraph[i]
		if unicode.IsLetter(rune(letter)) || unicode.IsSpace(rune(letter)) {
			pureText = append(pureText, letter)
		} else {
			pureText = append(pureText, ' ')
		}
	}
	wordArray := strings.Split(strings.ToLower(string(pureText)), " ")
	wordMap := make(map[string]int, 0)
	maxWord := ""
	maxNum := 0
	for i := 0; i < len(wordArray); i++ {
		word := strings.TrimSpace(wordArray[i])
		// 都是空格的字符串 len是0
		if bannedMap[word] == true || len(word) == 0 {
			continue
		}
		if _, ok := wordMap[word]; !ok {
			wordMap[word] = 0
		}
		wordMap[word] += 1
		if wordMap[word] > maxNum {
			maxWord = word
			maxNum = wordMap[word]
		}
	}
	return maxWord
}
