/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverFromPreorder(S string) *TreeNode {

}

// https://leetcode.com/problems/recover-a-tree-from-preorder-traversal/submissions/
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
/* 用hash map的解法
class Solution {
    public TreeNode recoverFromPreorder(String S) {
      Map<Integer, TreeNode> levels = new HashMap<>();
        int i = 0;
        while(i < S.length()) {
            int curLevel = 0, curNum = 0;
            while(i < S.length() && S.charAt(i) == '-') {
                ++curLevel;
                ++i;
            }
            while(i < S.length() && S.charAt(i) >= '0' && S.charAt(i) <= '9') {
                curNum = curNum*10 + (S.charAt(i) - '0');
                i++;
            }
            TreeNode curNode = new TreeNode(curNum);
            levels.put(curLevel, curNode);
            if(curLevel > 0) {
                TreeNode parent = levels.get(curLevel - 1);
                if(parent.left == null) parent.left = curNode;
                else parent.right = curNode;
            }
        }
        return levels.get(0);
    }
}
*/
/*
用栈 stack的数据结构来判断level的关系
    public TreeNode recoverFromPreorder(String S) {
        int level, val;
        Stack<TreeNode> stack = new Stack<>();
        for (int i = 0; i < S.length();) {
						// 获得当前的level
            for (level = 0; S.charAt(i) == '-'; i++) {
                level++;
						}
						// 为了字符串到数字的转换
            for (val = 0; i < S.length() && S.charAt(i) != '-'; i++) {
                val = val * 10 + (S.charAt(i) - '0');
						}
						// 当stack的size大于level的时候 说明已经遍历到右子树了 就可以pop了
            while (stack.size() > level) {
                stack.pop();
            }
            TreeNode node = new TreeNode(val);
            if (!stack.isEmpty()) {
                if (stack.peek().left == null) {
                    stack.peek().left = node;
                } else {
                    stack.peek().right = node;
                }
            }
            stack.add(node);
        }
        while (stack.size() > 1) {
            stack.pop();
        }
        return stack.pop();
    }
*/
