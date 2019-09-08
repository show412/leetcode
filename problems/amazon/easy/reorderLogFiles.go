import (
	"sort"
	"strings"
)

// https://leetcode.com/problems/reorder-data-in-log-files/
/*
You have an array of logs.  Each log is a space delimited string of words.

For each log, the first word in each log is an alphanumeric identifier.
Then, either:

Each word after the identifier will consist only of lowercase letters, or;
Each word after the identifier will consist only of digits.
We will call these two varieties of logs letter-logs and digit-logs.
It is guaranteed that each log has at least one word after its identifier.

Reorder the logs so that all of the letter-logs come before any digit-log.
The letter-logs are ordered lexicographically ignoring identifier,
with the identifier used in case of ties.
The digit-logs should be put in their original order.

Return the final order of the logs.



Example 1:

Input: logs = ["dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"]
Output: ["let1 art can","let3 art zero","let2 own kit dig","dig1 8 1 5 1","dig2 3 6"]


Constraints:

0 <= logs.length <= 100
3 <= logs[i].length <= 100
logs[i] is guaranteed to have an identifier, and a word after the identifier.
*/
/*
test case:
["dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"]
["w 7 2", "l 1 0", "6 066", "o aay", "e yal"]
["a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act zoo"]
["1 n u", "r 527", "j 893", "6 14", "6 82"]
["8 fj dzz k", "5r 446 6 3", "63 gu psub", "5 ba iedrz", "6s 87979 5", "3r 587 01", "jc 3480612", "bb wsrd kp", "b aq cojj", "r5 6316 71"]
*/
type log []string

func (this log) Len() int {
	return len(this)
}
func (this log) Less(i, j int) bool {
	// 这里要用splitN 只把字符串分成两个数组
	a := strings.SplitN(this[i], " ", 2)
	b := strings.SplitN(this[j], " ", 2)
	if !checkDig(a[1][0]) && !checkDig(b[1][0]) {
		if a[1] != b[1] {
			return a[1] < b[1]
		}
		return a[0] < b[0]
	}
	if checkDig(a[1][0]) {
		if checkDig(b[1][0]) {
			return false
		}
		return false
	}
	return true
}

func (this log) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func checkDig(a byte) bool {
	return a >= '0' && a <= '9'
}

func reorderLogFiles(logs []string) []string {
	// 注意这里要用stable 保证顺序
	sort.Stable(log(logs))
	return logs
}
