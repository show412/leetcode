import "math"

// https://leetcode.com/problems/bulb-switcher/
/*
There are n bulbs that are initially off.
You first turn on all the bulbs.
Then, you turn off every second bulb.
On the third round, you toggle every third bulb (turning on if it's off or turning off if it's on).
For the i-th round, you toggle every i bulb.
For the n-th round, you only toggle the last bulb.
Find how many bulbs are on after n rounds.

Example:

Input: 3
Output: 1
Explanation:
At first, the three bulbs are [off, off, off].
After first round, the three bulbs are [on, on, on].
After second round, the three bulbs are [on, off, on].
After third round, the three bulbs are [on, off, off].

So you should return 1, because there is only one bulb is on.
*/
// https://www.cnblogs.com/grandyang/p/5100098.html
/*
对于第n个灯泡，只有当次数是n的因子的之后，才能改变灯泡的状态，即n能被当前次数整除，
比如当n为36时，它的因数有(1,36), (2,18), (3,12), (4,9), (6,6),
可以看到前四个括号里成对出现的因数各不相同，括号中前面的数改变了灯泡状态，后面的数又变回去了，
等于灯泡的状态没有发生变化，只有最后那个(6,6)，在次数6的时候改变了一次状态，
没有对应其它的状态能将其变回去了，所以灯泡就一直是点亮状态的。
所以所有平方数都有这么一个相等的因数对，即所有平方数的灯泡都将会是点亮的状态。

那么问题就简化为了求1到n之间完全平方数的个数，我们可以用force brute来比较从1开始的完全平方数和n的大小
*/
//关键是 对于第n个灯泡，只有当次数是n的因子的之后，才能改变灯泡的状态
/*
现在假如我们执行第i次操作，即从编号i开始对编号每次+i进行switch操作，对于这些灯来说，
如果其编号j（j=1,2,3,⋯,n）能够整除i，则编号j的灯需要执switch操作。
具备这样性质的i是成对出现的，比如：
j=12时，编号为12的灯，在第1次，第12次；第2次，第6次；第3次，第4次一定会被执行Switch操作，
这样的话，编号为12的等肯定为灭。
但是当完全平方数36就不一样了，因为他有一个特殊的因数6，
这样当i=6时，只能被执行一次Switch操作，
这样推出，完全平方数一定是亮着的，所以本题的关键在于找完全平方数的个数。

*/
func bulbSwitch(n int) int {
	res := 1
	for res*res <= n {
		res++
	}
	return res - 1
}

// 还有一种方法更简单，我们直接对n开方，这个整型数的平方最接近于n，即为n包含的所有完全平方数的个数，
func bulbSwitch(n int) int {
	return int(math.Floor(math.Sqrt(float64(n))))
}
