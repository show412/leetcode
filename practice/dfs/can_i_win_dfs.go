// https://leetcode.com/problems/can-i-win/
/*
In the "100 game," two players take turns adding, to a running total, any integer from 1..10. The player who first causes the running total to reach or exceed 100 wins.

What if we change the game so that players cannot re-use integers?

For example, two players might take turns drawing from a common pool of numbers of 1..15 without replacement until they reach a total >= 100.

Given an integer maxChoosableInteger and another integer desiredTotal, determine if the first player to move can force a win, assuming both players play optimally.

You can always assume that maxChoosableInteger will not be larger than 20 and desiredTotal will not be larger than 300.

Example

Input:
maxChoosableInteger = 10
desiredTotal = 11

Output:
false

Explanation:
No matter which integer the first player choose, the first player will lose.
The first player can choose an integer from 1 up to 10.
If the first player choose 1, the second player can only choose integers from 2 up to 10.
The second player will win by choosing 10 and get a total = 11, which is >= desiredTotal.
Same with other integers chosen by the first player, the second player will always win.
*/
// 6mon - 1 year  Google 3
/*
使用 HashMap 来记录已经计算过的结果。
我们首先来看如果给定的数字范围大于等于目标值的话，直接返回 true。
如果给定的数字总和小于目标值的话，说明谁也没法赢，返回 false。
然后我们进入递归函数，首先我们查找当前情况是否在 HashMap 中存在，有的话直接返回即可。
我们使用一个整型数按位来记录数组中的某个数字是否使用过，我们遍历所有数字，
将该数字对应的 mask 算出来，如果其和 used 相与为0的话，说明该数字没有使用过，
我们看如果此时的目标值小于等于当前数字，说明已经赢了，或者调用递归函数，如果返回 false，
说明也是第一个人赢了。为啥呢，因为当前已经选过数字了，此时就该对第二个人调用递归函数，
只有返回的结果是 false，我们才能赢，所以此时我们 true，并返回 true。
如果遍历完所有数字，标记 false，并返回 false，
*/
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger >= desiredTotal {
		return true
	}
	if maxChoosableInteger*(maxChoosableInteger+1)/2 < desiredTotal {
		return false
	}
	m := make(map[int]bool, 0)
	return canWin(maxChoosableInteger, desiredTotal, 0, m)
}

func canWin(length int, total int, used int, m map[int]bool) bool {
	if _, ok := m[used]; ok {
		return m[used]
	}
	for i := 0; i < length; i++ {
		cur := (1 << uint32(i))
		if (cur & used) == 0 {
			if total <= i+1 || !canWin(length, total-(i+1), cur|used, m) {
				m[used] = true
				return true
			}
		}
	}
	m[used] = false
	return false
}
