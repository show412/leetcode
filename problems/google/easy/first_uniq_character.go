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
/*
这的思路是因为都是小写字母 所以 byte 的大小是一致的 小写字母都在256的 accill code 之内
用快慢指针 , slow 总是保持在只出现一次的字母的位置 fast 会向前走 如果有多次出现的就会回来
重写 count 里的值 这样 slow 再继续往前走找到下一个只出现一次的地方
*/
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