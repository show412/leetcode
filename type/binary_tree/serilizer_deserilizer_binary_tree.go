/*
 * @Author: hongwei.sun
 * @Date: 2024-04-10 16:19:20
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-10 16:19:20
 * @Description: file content
 */
// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/description/
// Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

// Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

// Example: 

// You may serialize the following tree:

//     1
//    / \
//   2   3
//      / \
//     4   5

// as "[1,2,3,null,null,4,5]"
// Clarification: The above format is the same as how LeetCode serializes a binary tree. You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.

// Note: Do not use class member/global/static variables to store states. Your serialize and deserialize algorithms should be stateless.



/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// serialize can use preorder to be in one string
// deserialize can be from this string
// key is to figure out dfs
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 type Codec struct {
    
 }
 
 func Constructor() Codec {
     return Codec{}
 }
 
 // Serializes a tree to a single string.
 func (this *Codec) serialize(root *TreeNode) string {
    str := ""
    // 函数的定义要提前定义 因为dfs里面有递归
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
    if str != "" {
        str += ","
    }
    if node == nil {
        str += "N"
        return
    }
    str += strconv.Itoa(node.Val)
    dfs(node.Left)
    dfs(node.Right)
    }
    dfs(root)
    // fmt.Println("serialize is ", str)
    return str
 }
 
 // Deserializes your encoded data to tree.
 func (this *Codec) deserialize(data string) *TreeNode {    
    array := strings.Split(data, ",")
    var dfs func() *TreeNode
    dfs = func () *TreeNode {
        cur := array[0]
        array = array[1:]
        // fmt.Println("cur is ", cur)
        // fmt.Println("-------")
        // return nil here, 也就因为着递归的结束，将会逐层return 因为不再调一样的函数了
        // 这也是递归的常用结束方式
        if cur == "N" {
            return nil
        }
        val, _ := strconv.Atoi(cur)
        return &TreeNode{val, dfs(), dfs()}
    }
    return dfs()
 }

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */



//  java implementation
// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/description/
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
public class Codec {

    // Encodes a tree to a single string.
    public String serialize(TreeNode root) {
        if (root == null) {
            return "{}";
        }

        ArrayList<TreeNode> queue = new ArrayList<TreeNode>();
        queue.add(root);

        for (int i = 0; i < queue.size(); i++) {
            TreeNode node = queue.get(i);
            if (node == null) {
                continue;
            }
            queue.add(node.left);
            queue.add(node.right);
        }

        while (queue.get(queue.size() - 1) == null) {
            queue.remove(queue.size() - 1);
        }

        StringBuilder sb = new StringBuilder();
        sb.append("{");
        sb.append(queue.get(0).val);
        for (int i = 1; i < queue.size(); i++) {
            if (queue.get(i) == null) {
                sb.append(",#");
            } else {
                sb.append(",");
                sb.append(queue.get(i).val);
            }
        }
        sb.append("}");
        return sb.toString();
    }

    // Decodes your encoded data to tree.
    public TreeNode deserialize(String data) {
        if (data.equals("{}")) {
            return null;
        }
        String[] vals = data.substring(1, data.length() - 1).split(",");
        ArrayList<TreeNode> queue = new ArrayList<TreeNode>();
        TreeNode root = new TreeNode(Integer.parseInt(vals[0]));
        queue.add(root);
        int index = 0;
        boolean isLeftChild = true;
        for (int i = 1; i < vals.length; i++) {
            if (!vals[i].equals("#")) {
                TreeNode node = new TreeNode(Integer.parseInt(vals[i]));
                if (isLeftChild) {
                    queue.get(index).left = node;
                } else {
                    queue.get(index).right = node;
                }
                queue.add(node);
            }
            if (!isLeftChild) {
                index++;
            }
            isLeftChild = !isLeftChild;
        }
        return root;
    }
}

// Your Codec object will be instantiated and called as such:
// Codec codec = new Codec();
// codec.deserialize(codec.serialize(root));