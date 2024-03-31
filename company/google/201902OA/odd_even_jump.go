// https://leetcode.com/problems/odd-even-jump/
// it should be DP to solve it
// but it's a very hard DP， from end to start to traverse array
func oddEvenJumps(A []int) int {
	// from i to j
	if len(A) <= 1 {
		return len(A)
	}
	// State in DP
	// oddJump[i] means the i could be the end by odd jump
	oddJump := make(map[int]bool)
	// evenJump[i] means the i could be the end by even jump
	evenJump := make(map[int]bool)
	// the last i could be good point
	res := 1
	// use a hash to store the key is A[i] value is i
	m := make(map[int]int)
	// init in DP
	oddJump[len(A)-1] = true
	evenJump[len(A)-1] = true
	m[A[len(A)-1]] = len(A) - 1
	for i := len(A) - 2; i >= 0; i-- {
		// find out the smallest value which is bigger than A[i]
		oddValue := getFoolrKey(m, A[i])
		// find out the smallest value which is smaller than A[i]
		evenValue := getCellingKey(m, A[i])
		/*
		 因为从偶数点能跳到的下一个点是i 如果i能奇数跳到end 取决于vaule能不能偶数跳到end
		 因为value下一跳是i，这样偶数-1 就是奇数 i就可以奇数跳到end
		 oddvalue和evenvalue肯定有一个不为空， 因为是倒序遍历i 所以找到的value一定小于i
		 同理i是偶数跳的时候
		*/
		// start to function in DP
		if oddValue >= 0 {
			oddJump[i] = evenJump[oddValue]
		}
		if evenValue >= 0 {
			evenJump[i] = oddJump[evenValue]
		}
		// start index must be a oddJump, why?
		// 因为奇数跳大于等于偶数跳 所以以奇数跳为准
		if oddJump[i] == true {
			res++
		}
		m[A[i]] = i
	}
	return res
}

func getFoolrKey(h map[int]int, k int) int {
	const MAX = int(^uint(0) >> 1)
	min := MAX
	minV := -1
	for key, value := range h {
		if key >= k {
			if key < min {
				min = key
				minV = value
			}
		}
	}
	return minV
}

func getCellingKey(h map[int]int, k int) int {
	const MAX = int(^uint(0) >> 1)
	min := MAX
	minV := -1
	for key, value := range h {
		if key <= k {
			if key < min {
				min = key
				minV = value
			}
		}
	}
	return minV
}

// It's a java solution
// https://www.jiuzhang.com/solution/odd-even-jump/#tag-highlight
// public class Solution {
//     /**
//      * @param A: An integer array A
//      * @return: Return the number of good starting indexes
//      */
//     public int oddEvenJumps(int[] A){
//         int res = 1, lenA = A.length;

//         // higher[i] = true代表从索引 i 奇数跳可以到达末尾
//         boolean []higher = new boolean[lenA];
//         // lower[i] = true代表从索引 i 偶数跳可以到达末尾
//         boolean []lower = new boolean[lenA];
//         // 初始化最后一个位置，无论哪种跳法都是已经在末尾了
//         higher[lenA - 1] = lower[lenA - 1] = true;
//         // 记录 value - index
//         TreeMap<Integer, Integer> treeMap = new TreeMap<>();
//         treeMap.put(A[lenA - 1], lenA - 1);
//         for (int i = lenA - 2; i >= 0; i--) {  // 逆推，分两种情况
//  这是java特有的数据结构的方法
//             // 找到index = i 之后大于等于A[i]的最小的数
//             Map.Entry high = treeMap.ceilingEntry(A[i]);
//             // 找到index = i 之后小于等于A[i]的最大的数
//             Map.Entry low = treeMap.floorEntry(A[i]);
// 是如何DP的呢， 因为high和low一定至少有一个存在 这样higher[i] 和 lower[i]就一定有一个会被赋值
// 因为A[lenA-1]已经在hash里了
// 然后还是i-1倒序遍历的 保证了是去取索引i之后的值
//             if (high != null){
// 当前位置进行奇数跳取决于目的地的偶数跳是否能到达终点
// 因为high.getValue 要跳到的下一个点是i， 如果这个i能跳到终点 value这个点就也能跳到
//                 higher[i] = lower[(int)high.getValue()];
//             }
//             if (low != null){  // 同理
//                 lower[i] = higher[(int)low.getValue()];
//             }
//             // 起始点一定是奇数跳 所以去奇数点里遍历
//             if (higher[i]){
//                 res ++;
//             }
//             treeMap.put(A[i], i);
//         }
//         return res;
//     }
// }
