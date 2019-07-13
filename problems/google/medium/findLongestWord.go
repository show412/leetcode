// https://leetcode.com/problems/longest-word-in-dictionary-through-deleting/
/*
Given a string and a string dictionary, find the longest string in the dictionary that can be formed by deleting some characters of the given string. If there are more than one possible results, return the longest word with the smallest lexicographical order. If there is no possible result, return the empty string.

Example 1:
Input:
s = "abpcplea", d = ["ale","apple","monkey","plea"]

Output:
"apple"
Example 2:
Input:
s = "abpcplea", d = ["a","b","c"]

Output:
"a"
Note:
All the strings in the input will only contain lower-case letters.
The size of the dictionary won't exceed 1,000.
The length of all the strings in the input won't exceed 1,000.
*/
/*
use case:
	Input:
s = "abpcplea", d = ["ale","apple","monkey","plea"]

Output:
"apple"

Input:
s = "abpcplea", d = ["a","b","c"]

Output:
"a"

"mjmqqjrmzkvhxlyruonekhhofpzzslupzojfuoztvzmmqvmlhgqxehojfowtrinbatjujaxekbcydldglkbxsqbbnrkhfdnpfbuaktupfftiljwpgglkjqunvithzlzpgikixqeuimmtbiskemplcvljqgvlzvnqxgedxqnznddkiujwhdefziydtquoudzxstpjjitmiimbjfgfjikkjycwgnpdxpeppsturjwkgnifinccvqzwlbmgpdaodzptyrjjkbqmgdrftfbwgimsmjpknuqtijrsnwvtytqqvookinzmkkkrkgwafohflvuedssukjgipgmypakhlckvizmqvycvbxhlljzejcaijqnfgobuhuiahtmxfzoplmmjfxtggwwxliplntkfuxjcnzcqsaagahbbneugiocexcfpszzomumfqpaiydssmihdoewahoswhlnpctjmkyufsvjlrflfiktndubnymenlmpyrhjxfdcq"




output: "ntgcykxhdfescxxypyew"
*/
func findLongestWord(s string, d []string) string {
	if len(d) == 0 || len(s) == 0 {
		return ""
	}
	candidateList := make([]string, 0)
	// find which word is in s
	for i := 0; i < len(d); i++ {
		word := d[i]
		pointS := 0
		pointW := 0
		for pointW < len(word) && pointS < len(s) {
			if pointW == len(word)-1 && s[pointS] == word[pointW] {
				candidateList = append(candidateList, word)
				break
			}
			if s[pointS] == word[pointW] {
				pointS++
				pointW++
				continue
			} else {
				if (len(s) - pointS) < (len(word) - pointW) {
					break
				}
				pointS++
			}
		}
	}
	// find out the result
	if len(candidateList) == 0 {
		return ""
	}
	res := ""
	for i := 0; i < len(candidateList); i++ {
		word := candidateList[i]
		if len(word) > len(res) {
			res = word
			continue
		}
		if len(word) == len(res) {
			res = checkLex(word, res)
		}
	}
	return res
}

// 字典排序
func checkLex(word1 string, word2 string) string {
	s := 0
	for s < len(word1) {
		if word1[s] < word2[s] {
			return word1
		} else if word1[s] > word2[s] {
			return word2
		}
		s++
	}
	return word1
}
