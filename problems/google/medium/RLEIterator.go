// https://leetcode.com/problems/rle-iterator/
/*
Write an iterator that iterates through a run-length encoded sequence.

The iterator is initialized by RLEIterator(int[] A), where A is a run-length encoding of some sequence.  More specifically, for all even i, A[i] tells us the number of times that the non-negative integer value A[i+1] is repeated in the sequence.

The iterator supports one function: next(int n), which exhausts the next n elements (n >= 1) and returns the last element exhausted in this way.  If there is no element left to exhaust, next returns -1 instead.

For example, we start with A = [3,8,0,9,2,5], which is a run-length encoding of the sequence [8,8,8,5,5].  This is because the sequence can be read as "three eights, zero nines, two fives".



Example 1:

Input: ["RLEIterator","next","next","next","next"], [[[3,8,0,9,2,5]],[2],[1],[1],[2]]
Output: [null,8,8,5,-1]
Explanation:
RLEIterator is initialized with RLEIterator([3,8,0,9,2,5]).
This maps to the sequence [8,8,8,5,5].
RLEIterator.next is then called 4 times:

.next(2) exhausts 2 terms of the sequence, returning 8.  The remaining sequence is now [8, 5, 5].

.next(1) exhausts 1 term of the sequence, returning 8.  The remaining sequence is now [5, 5].

.next(1) exhausts 1 term of the sequence, returning 5.  The remaining sequence is now [5].

.next(2) exhausts 2 terms, returning -1.  This is because the first term exhausted was 5,
but the second term did not exist.  Since the last term exhausted does not exist, we return -1.

Note:

0 <= A.length <= 1000
A.length is an even integer.
0 <= A[i] <= 10^9
There are at most 1000 calls to RLEIterator.next(int n) per test case.
Each call to RLEIterator.next(int n) will have 1 <= n <= 10^9.
*/
/*
use case:
	["RLEIterator","next","next","next","next"]
[[[3,8,0,9,2,5]],[2],[1],[1],[2]]

["RLEIterator","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next","next"]
[[[635,606,576,391,370,981,36,21,961,185,124,210,801,937,22,426,101,260,221,647,350,180,504,39,101,989,199,358,732,839,919,169,673,967,58,676,846,342,250,597,442,174,472,410,569,509,311,357,838,251]],[5039],[62],[3640],[671],[67],[395],[262],[839],[74],[43],[42],[77],[13],[6],[3],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1]]
*/
type RLEIterator struct {
	number []int
	times  []int
}

func Constructor(A []int) RLEIterator {
	number := make([]int, 0)
	times := make([]int, 0)
	for i := 0; i < len(A); i += 2 {
		times = append(times, A[i])
		number = append(number, A[i+1])
	}
	return RLEIterator{number: number, times: times}
}

func (this *RLEIterator) Next(n int) int {
	left := n
	for len(this.times) > 0 && this.times[0]-left < 0 {
		left -= this.times[0]
		this.times = this.times[1:]
		this.number = this.number[1:]
	}
	if len(this.times) > 0 && this.times[0]-left >= 0 {
		this.times[0] -= left
		return this.number[0]
	}
	if len(this.times) > 0 {
		return this.number[0]
	} else {
		return -1
	}
}

/**
 * Your RLEIterator object will be instantiated and called as such:
 * obj := Constructor(A);
 * param_1 := obj.Next(n);
 */
