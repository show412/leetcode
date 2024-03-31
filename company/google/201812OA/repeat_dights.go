// https://leetcode.com/problems/numbers-with-repeated-digits/

/*
Given a positive integer N, return the number of positive integers less than or equal to N that have at least 1 repeated digit.
Example 1:

Input: 20
Output: 1
Explanation: The only positive number (<= 20) with at least 1 repeated digit is 11.
Example 2:

Input: 100
Output: 10
Explanation: The positive numbers (<= 100) with atleast 1 repeated digit are 11, 22, 33, 44, 55, 66, 77, 88, 99, and 100.
Example 3:

Input: 1000
Output: 262

Note:

1 <= N <= 10^9
*/

//  the answer is right but it's TLE (time limit exceeded)
func numDupDigitsAtMostN(N int) int {
	if N < 11 {
		return 0
	}
	count := 0
	for i := 11; i <= N; i++ {
		m := make(map[int]bool)
		num := i
		// last := num % 10
		for num > 0 {
			last := num % 10
			num = num / 10
			if v, ok := m[last]; ok && v == true {
				count++
				break
			} else {
				m[last] = true
			}
		}
	}
	return count
}

// java accept answer
// https://leetcode.com/problems/numbers-with-repeated-digits/discuss/256725/JavaPython-Count-the-Number-Without-Repeated-Digit
// https://www.jianshu.com/p/e71ba9354447

class Solution {
  public int numDupDigitsAtMostN(int N) {
        List<Integer> digits = new ArrayList<>();
    //这个N + 1是解法的精髓，想想如果N=20 （末尾为0的情况）
        for (int tmp = N + 1; tmp > 0; tmp /= 10) {
            digits.add(0, tmp % 10);
        }

        int len = digits.size();

        int noRepeatNum = 0;
        for (int i = 0; i < len - 1; i++) {
            noRepeatNum += calNoRepeatWithinDigitNum(i);
        }
        // 用set来记录经过的prefix
        // 如果出现N=220, 当我们发现十位上的数已经出现过了
        // break掉loop，也就是说不用计算22X
        Set<Integer> set = new HashSet<>();
        for (int i = 0; i < len; i++) {
            int start = 1;
            if (i != 0) {
                start = 0;
            }
            for (int j = start; j < digits.get(i); j++) {
                if (!set.contains(j)) {
                    noRepeatNum += calNoRepeatWithinDigitNum(9 - i, len - i - 1);
                }
            }
            if (set.contains(digits.get(i))) break;

            set.add(digits.get(i));
        }
        return N - noRepeatNum;
    }

    private int calNoRepeatWithinDigitNum(int num) {
        int res = 9;
        for (int i = 0; i < num; i++) {
            res *= (9 - i);
        }
        return res;
    }

    private int calNoRepeatWithinDigitNum(int largestNum, int num) {
        if (num == 0) return 1;
        int res = largestNum;
        for (int i = 1; i < num; i++) {
            res *= (largestNum - i);
        }
        return res;
    }
}