/**
 * @param s: the long string
 * @param word: the secrect word
 * @return: whether a substring exists in the string can be transformed by the above encoding rule
 */
//  https://www.lintcode.com/problem/secret-word/description

/* Give `s="abcabcd"`, `word="xyzxyz"`, return `yes`
Input:
abcabcd
xyzxyz
Output:
yes

Explaination:
"x" can transfer to "a", "y" can transfer to "b" and "z" can transfer to "c".
Examlpe 2:

Give `s="abca"`, `word="xyzd"`, return `no`
Input:
abca
xyzd
Output:
no

Explaination:
the word "xyzd" has no way to transfer to "abca" */

func getAns(s string, word string) string {
	// Write a code here
	if len(s) == 0 || len(word) == 0 || len(s) < len(word) {
		return "no"
	}
	slen := len(s)
	wlen := len(word)
	// Give `s="abcabcd"`, `word="xyzxyz"`, return `yes`
	// "jhdsgadhgsahdjgsahjdgashjdgshaj", "aaaaa"
	for i := 0; i < (slen - wlen + 1); i++ {
		j := 0
		m1 := make(map[string]string)
		m2 := make(map[string]string)
		for j < wlen {
			v1, ok1 := m1[string(s[i+j])]
			v2, ok2 := m2[string(word[j])]
			if !ok1 && !ok2 {
				m1[string(s[i+j])] = string(word[j])
				m2[string(word[j])] = string(s[i+j])
				j++
			} else if ok1 && ok2 && v1 == string(word[j]) && v2 == string(s[i+j]) {
				j++
			} else {
				break
			}
			if j == wlen {
				// fmt.Println(m1)
				// fmt.Println(m2)
				return "yes"
			}
		}
		i++
	}
	return "no"
}
