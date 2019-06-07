/*
Given a word s, and a string set str. This word removes one letter at a time until there is only one letter in the word. Is there a sort of deletion in which all words are in str. For example, the word is 'abc', the string set is {'a', 'ab', 'abc'}, if the order of deletion is 'c', 'b', and 'abc', 'ab', 'a' are all in the collection, so it is eligible.
Output whether this combination meets the condition.

Example
Example 1:

Input：s="abc"，str=["abc","ac","c"]
Output：true
Explanation:
First "abc" is in `str`
Delete 'b', "ac" is  in `str`
Delete 'a', "c" is in `str`
Example 2:

Input：s="abc",str=["abc","ab","c"]
Output：false
Explanation:
"abc" in `str`
Next you can only delete 'c', "ab" in `str`
Since "a" and "b" are not in `str`, return false

Notice
1<=|str[i]|,|s|<=30
1<=the number of strings in str<=100
*/
/**
 * @param s:
 * @param str:
 * @return: Output whether this combination meets the condition
 */
//  it has better performance
func checkWord(s string, str []string) bool {
	// Write your code here
	slen := len(s)
	if slen > len(str) {
		return false
	}
	if slen == 0 {
		return false
	}
	// result := make([]string, 0)
	m := make(map[string]bool)
	for i := 0; i < len(str); i++ {
		m[str[i]] = true
	}
	// strMap := make(map[string]bool, len(s))
	return dfs(s, m)
}
func dfs(s string, m map[string]bool) bool {

	if exist, ok := m[s]; !ok || exist == false {
		return false
	}
	// 只要到了1个字母的时候 因为前面的check过了 就一定是true
	if len(s) == 1 {
		return true
	}
	// if v, ok := strMap[s]; ok && v == true {
	// 	return false
	// }
	for i := 0; i < len(s); i++ {
		sub := s
		if i == 0 {
			sub = s[1:]
		} else {
			sub = s[0:i] + s[(i+1):len(s)]
		}
		if dfs(sub, m) == true {
			return true
		}
		// result = result[0 : len(result)-1]
	}
	// 不需要判断曾经没曾经选过 因为dfs前面的枚举就保证后面的不会重
	// strMap[s] = true
	return false
}
