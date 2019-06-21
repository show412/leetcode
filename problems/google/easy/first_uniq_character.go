// https://leetcode.com/problems/first-unique-character-in-a-string/
/*
Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.

Examples:

s = "leetcode"
return 0.

s = "loveleetcode",
return 2.
Note: You may assume the string contain only lowercase letters.
*/
func firstUniqChar(s string) int {
	if s == "" {
		return -1
	}
	m := make(map[byte]int, 0)
	a := make(map[byte]bool, len(s))
	for i := 0; i < len(s); i++ {
		_, ok1 := m[s[i]]
		_, ok2 := a[s[i]]
		if ok1 {
			delete(m, s[i])
			a[s[i]] = true
		} else if !ok1 && !ok2 {
			m[s[i]] = i
		} else if !ok1 && ok2 {
			continue
		}
	}
	if len(m) == 0 {
		return -1
	}
	res := len(s)
	for _, v := range m {
		if v < res {
			res = v
		}
	}
	return res
}

// slow and fast pointer
func firstUniqChar(s string) int {
	if len(s) == 0 {
		return -1
	}
	lenny := len(s)
	if lenny == 1 {
		return 0
	}

	cc := []byte(s)
	slow := 0
	fast := 1
	count := make([]int, 256)
	count[cc[slow]]++
	for fast < lenny {
		count[cc[fast]]++
		for slow < lenny && count[cc[slow]] > 1 {
			slow++
		}
		if slow >= lenny {
			return -1
		}
		if count[cc[slow]] == 0 {
			count[cc[slow]]++
			fast = slow
		}
		fast++
	}
	return slow
}