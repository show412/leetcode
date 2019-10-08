import "unicode"

// https://leetcode.com/problems/basic-calculator-iii/
/*
Implement a basic calculator to evaluate a simple expression string.

The expression string may contain open ( and closing parentheses ),
the plus + or minus sign -, non-negative integers and empty spaces .

The expression string contains only non-negative integers, +, -, *, / operators ,
open ( and closing parentheses ) and empty spaces .
The integer division should truncate toward zero.

You may assume that the given expression is always valid.
All intermediate results will be in the range of [-2147483648, 2147483647].

Some examples:

"1 + 1" = 2
" 6-4 / 2 " = 4
"2*(5+5*2)/3+(6/2+8)" = 21
"(2+6* 3+5- (3*14/7+2)*5)+3"=-12


Note: Do not use the eval built-in library function.
*/
/*
this common solution for basic calculator 1,2,3,4
https://leetcode.com/problems/basic-calculator-iii/discuss/113592/Development-of-a-generic-solution-for-the-series-of-the-calculator-problems

You need to understand the meaning of (l1, o1) and (l2, o2).
For the latter, l2 and o2 only participate in computations of the level two.
When a + or - is encountered,
the higher level computation is finished so the result is demoted to level one
(you see we're updating l1 and o1 here).
Then we need to prepare l2 and o2 for the following computations by reinitializing them to 1.
This is to say that computations of level two for different parts in the expression are independent from each other.
For example, in the expression 2 * 3 + 3 * 4 + 4 * 5, the computation of 2 * 3 is independent from 3 * 4 and 4 * 5, etc.
So when carrying out computations for each of them, l2 and o2 need to be properly initialized.

And why don't we reset l1 and o1?
Because they participate in computations of level one,
and there is no other levels that are even lower than level one.
Conceptually, l1 will hold the partial results accumulated so far.
Again using the above example, after computing 2 * 3, we have l1 = 6,
which is the partial result of 2 * 3. After computing 3 * 4,
we have l1 = 18, which is the partial result of 2 * 3 + 3 * 4, and so.
We surely do not want to reset l1 in the middle of the computation,
otherwise we lose some of the partial results and the final answer will be wrong.

Hope this clear it up!

这个思路是把运算式分为两个level 一个是加减的一个是乘除， 用l1 o1 l2 o2来记应该的数和操作
任何运算公式最后的结果其实都是l1+o1*l2
*/
func calculate(s string) int {
	l1, o1, l2, o2 := 0, 1, 1, 1
	// 如果第一个是 - 按这种思路-之后l1应该是l1=o1*l2 这个时候是1 但实际应该是0 所以应该在前面加0
	if s[0] == '-' {
		s = "0" + s
	}
	for i := 0; i < len(s); i++ {
		c := rune(s[i])
		if unicode.IsDigit(c) {
			num := int(c - '0')
			for (i+1) < len(s) && unicode.IsDigit(rune(s[i+1])) {
				num = num*10 + int(rune(s[i+1])-'0')
				i++
			}

			if o2 == 1 {
				l2 = l2 * num
			} else {
				l2 = l2 / num
			}
		} else if c == '(' {
			j := i
			for cnt := 0; i < len(s); i++ {
				if s[i] == '(' {
					cnt++
				}
				if s[i] == ')' {
					cnt--
				}
				if cnt == 0 {
					break
				}
			}
			num := calculate(s[(j + 1):i])

			if o2 == 1 {
				l2 = l2 * num
			} else {
				l2 = l2 / num
			}
		} else if c == '*' || c == '/' {
			if c == '*' {
				o2 = 1
			} else {
				o2 = -1
			}
		} else if c == '+' || c == '-' {
			// 到这了说明已经到了level 1的运算 会把l2的结果下沉到l1 所以要把l1和o1 reset
			l1 = l1 + o1*l2
			if c == '+' {
				o1 = 1
			} else {
				o1 = -1
			}
			l2 = 1
			o2 = 1
		}
	}
	return (l1 + o1*l2)
}
