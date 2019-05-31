// https://leetcode.com/problems/odd-even-jump/
// it should be DP to solve it
// but it's a very hard DP， from end to start to traverse array
func oddEvenJumps(A []int) int {
	// from i to j

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
// 因为i 要跳到的下一个点事high.getValue， 如果这个value能跳到终点 i这个点就也能跳到
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
