import "math"

// https://leetcode.com/problems/new-21-game/
/*
Alice plays the following game, loosely based on the card game "21".

Alice starts with 0 points, and draws numbers while she has less than K points.
During each draw, she gains an integer number of points randomly from the range [1, W], where W is an integer.
Each draw is independent and the outcomes have equal probabilities.

Alice stops drawing numbers when she gets K or more points.  What is the probability that she has N or less points?

Example 1:

Input: N = 10, K = 1, W = 10
Output: 1.00000
Explanation:  Alice gets a single card, then stops.
Example 2:

Input: N = 6, K = 1, W = 10
Output: 0.60000
Explanation:  Alice gets a single card, then stops.
In 6 out of W = 10 possibilities, she is at or below N = 6 points.
Example 3:

Input: N = 21, K = 17, W = 10
Output: 0.73278
Note:

0 <= K <= N <= 10000
1 <= W <= 10000
Answers will be accepted as correct if they are within 10^-5 of the correct answer.
The judging time limit has been reduced for this question.
*/
// each number from W is independet, means it could be same number

/*
Intuition

It is clear that the probability that Alice wins the game is only related to how many points x
she starts the next draw with, so we can try to formulate an answer in terms of x.

Algorithm

Let f(x) be the answer when we already have x points.
When she has between K and N points, then she stops drawing and wins. If she has more than N points, then she loses.

The key recursion is f(x) = (\frac{1}{W}) * (f(x+1) + f(x+2) + ... + f(x+W))f(x)=(
W
1
​
 )∗(f(x+1)+f(x+2)+...+f(x+W)) This is because there is a probability of \frac{1}{W}
W
1
​
  to draw each card from 11 to WW.

With this recursion, we could do a naive dynamic programming solution as follows:

#psuedocode
dp[k] = 1.0 when K <= k <= N, else 0.0
# dp[x] = the answer when Alice has x points
for k from K-1 to 0:
    dp[k] = (dp[k+1] + ... + dp[k+W]) / W
return dp[0]
This leads to a solution with O(K*W + (N-K))O(K∗W+(N−K)) time complexity, which isn't efficient enough. Every calculation of dp[k] does O(W)O(W) work, but the work is quite similar.

Actually, the difference is f(x) - f(x-1) = \frac{1}{W} \big( f(x+W) - f(x) \big)f(x)−f(x−1)=
W
1
​
 (f(x+W)−f(x)). This allows us to calculate subsequent f(k)f(k)'s in O(1)O(1) time, by maintaining the numerator S = f(x+1) + f(x+2) + \cdots + f(x+W)S=f(x+1)+f(x+2)+⋯+f(x+W).

As we calculate each dp[k] = S / W, we maintain the correct value of this numerator S \Rightarrow S + f(k) - f(k+W)S⇒S+f(k)−f(k+W).

*/

//
func new21Game(N int, K int, W int) float64 {
	// dp[x] = the answer when Alice has x points
	dp := make([]float64, N+W+1)

	for k := K; k <= N; k++ {
		dp[k] = 1.0
	}
	// S = dp[k+1] + dp[k+2] + ... + dp[k+W]
	S := math.Min(float64(N-K+1), float64(W))
	for k := K - 1; k >= 0; k-- {
		dp[k] = S / float64(W)
		S += dp[k] - dp[k+W]
	}
	return dp[0]
}
