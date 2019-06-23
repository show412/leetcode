package main

import (
	"fmt"
)

func main() {
	// reg := regexp.MustCompile("([a-z]*)([0-9]+)(\\[\1)(.)*(\\]\1)")
	// str := "aaaa3[a2[cc]]2[bc]"
	// data := reg.FindAllStringSubmatch(str, -1)
	// fmt.Println(data)
	res := numSquares(7168)
	// [1, 1, 4, 2, 1, 1, 0, 0]
	fmt.Println(res)
}

// Input:  [0,1,2,4,5,7]
// Output: ["0->2","4->5","7"]
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
