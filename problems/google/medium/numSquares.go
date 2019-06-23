import "math"

// https://leetcode.com/problems/perfect-squares/
/*
Given a positive integer n, find the least number of perfect square numbers
(for example, 1, 4, 9, 16, ...) which sum to n.

Example 1:

Input: n = 12
Output: 3
Explanation: 12 = 4 + 4 + 4.
Example 2:

Input: n = 13
Output: 2
Explanation: 13 = 4 + 9.

the best solution is here
https://leetcode.com/problems/perfect-squares/discuss/71488/Summary-of-4-different-solutions-(BFS-DP-static-DP-and-mathematics)
*/

// DP solution
func numSquares(n int) int {
	if n <= 0 {
		return 0
	}
	// f means i need the square sum at least
	f := make([]int, n+1)
	f[0] = 0
	for i := 1; i <= n; i++ {
		// i could be the first i-1 square sum + 1
		f[i] = f[i-1] + 1
		for j := 1; j*j <= i; j++ {
			// For each i, it must be the sum of some number (i - j*j) and
			// a perfect square number (j*j).
			f[i] = min(f[i], f[i-j*j]+1)
		}
	}
	return f[n]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// LTE need to improve for 7168
func numSquares(n int) int {
	if n == 1 || n == 4 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 3
	}
	var list []int
	for i := 1; i < n; i++ {
		if i*i > n {
			break
		}
		list = append(list, i)
	}
	if list[len(list)-1]*list[len(list)-1] == n {
		return 1
	}
	var sub []int
	cnt := 0
	var res int
	dfs(list, 0, sub, cnt, &res, n)

	return res
}

func dfs(list []int, start int, sub []int, cnt int, res *int, max int) {

	for i := start; i < len(list); i++ {
		sub = append(sub, list[i])
		cnt += list[i] * list[i]
		if cnt == max {
			if *res == 0 || *res > len(sub) {
				*res = len(sub)
				sub = sub[:len(sub)-1]
				cnt -= list[i] * list[i]
				break
			}
		}
		if *res != 0 && *res < len(sub) {
			sub = sub[:len(sub)-1]
			cnt -= list[i] * list[i]
			break
		}
		dfs(list, i, sub, cnt, res, max)
		sub = sub[:len(sub)-1]
		cnt -= list[i] * list[i]
	}

}

// good performance with DFS
func numSquares(n int) int {
	cache := make(map[int]int)
	cache[0] = 0
	return doNumSquares(n, cache)
}

func doNumSquares(n int, cache map[int]int) int {
	// 如果算过了就不需要再算了 提高性能
	if found, ok := cache[n]; ok {
		return found
	}

	res := math.MaxInt32
	for i := 1; i*i <= n; i++ {
		target := n - i*i
		res = min(res, doNumSquares(target, cache)+1)
	}
	cache[n] = res
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
the Mathematical Solution is best
class Solution
{
private:
    int is_square(int n)
    {
        int sqrt_n = (int)(sqrt(n));
        return (sqrt_n*sqrt_n == n);
    }

public:
    // Based on Lagrange's Four Square theorem, there
    // are only 4 possible results: 1, 2, 3, 4.
    int numSquares(int n)
    {
        // If n is a perfect square, return 1.
        if(is_square(n))
        {
            return 1;
        }

        // The result is 4 if and only if n can be written in the
        // form of 4^k*(8*m + 7). Please refer to
        // Legendre's three-square theorem.
        while ((n & 3) == 0) // n%4 == 0
        {
            n >>= 2;
        }
        if ((n & 7) == 7) // n%8 == 7
        {
            return 4;
        }

        // Check whether 2 is the result.
        int sqrt_n = (int)(sqrt(n));
        for(int i = 1; i <= sqrt_n; i++)
        {
            if (is_square(n - i*i))
            {
                return 2;
            }
        }

        return 3;
    }
};
*/
