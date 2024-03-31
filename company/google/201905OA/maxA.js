// Imagine you have a special keyboard with the following keys:

// Key 1: (A): Print one 'A' on screen.
//   Key 2: (Ctrl - A): Select the whole screen.
//     Key 3: (Ctrl - C): Copy selection to buffer.
//       Key 4: (Ctrl - V): Print buffer on screen appending it after what has already been printed.
//         Now, you can only press the keyboard for N times(with the above four keys), find out the maximum numbers of 'A' you can print on screen.

//           Example
// Example 1:

// Input: 3
// Output: 3
// Explanation: A, A, A
// Example 2:

// Input: 7
// Output: 9
// Explanation: A, A, A, Ctrl A, Ctrl C, Ctrl V, Ctrl V
// Notice
// 1 <= N <= 50
// Answers will be in the range of 32 - bit signed integer.

// https://www.lintcode.com/problem/4-keys-keyboard/description

// 获得最优解的按键序列一定可以用以下两种子序列拼接而成:

// 数个连续的 A
// Ctrl - A + Ctrl - C + 数个连续的 Ctrl - V
// 设定状态: f[i] 表示i次按键可以获得的最多数量的A

/**
 * @param N: an integer
 * @return: return an integer
 */
const maxA = function (N) {
  // write your code here
  // f[i] means that we could get the max A on the ith times
  var f = [];
  f[0] = 0;
  f[1] = 1;
  for(var i=2; i<N+1; i++){
    f[i] = f[i-1]+1;
    for(var j=1; j<i -2; j++){
      // 根据数学推导  2^((i-j-1)/3) < 2(i-j-1) 就是两边同时log 然后可以约掉
      //从j到i ctr+A ctr+C ctr+V 肯定不如一直ctr+V 获得的最后的f[i]大
      f[i] = Math.max(f[i], f[j] * (i - j - 1));
    }
  }
  return f[N];
}
