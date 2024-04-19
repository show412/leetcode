/*
 * @Author: hongwei.sun
 * @Date: 2024-04-19 11:41:02
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-19 11:41:02
 * @Description: file content
 */
/*
 * @Author: hongwei.sun
 * @Date: 2024-04-19 11:40:26
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-19 11:40:55
 * @Description: file content
 */
/*
使用两个指针N1,N2，一个从链表1的头节点开始遍历，我们记为N1，一个从链表2的头节点开始遍历，我们记为N2。

让N1和N2一起遍历，当N1先走完链表1的尽头（为null）的时候，则从链表2的头节点继续遍历，同样，如果N2先走完了链表2的尽头，则从链表1的头节点继续遍历，也就是说，N1和N2都会遍历链表1和链表2。

因为两个指针，同样的速度，走完同样长度（链表1+链表2），不管两条链表有无相同节点，都能够到达同时到达终点。

（N1最后肯定能到达链表2的终点，N2肯定能到达链表1的终点）。

所以，如何得到公共节点：

有公共节点的时候，N1和N2必会相遇，因为长度一样嘛，速度也一定，必会走到相同的地方的，所以当两者相等的时候，则会第一个公共的节点
无公共节点的时候，此时N1和N2则都会走到终点，那么他们此时都是null，所以也算是相等了。
*/
 //双向奔赴？有点意思的题目，A+B = B+A，找最后相遇的点就是共同走的路
func FindFirstCommonNode( pHead1 *ListNode ,  pHead2 *ListNode ) *ListNode {
    // write code here
    if pHead1 == nil || pHead2 == nil {
        return nil
    }
 
    cur1, cur2 := pHead1, pHead2
    for cur1 != cur2 {
        if cur1 == nil {
            cur1 = pHead2
        } else {
            cur1 = cur1.Next
        }
 
        if cur2 == nil {
            cur2 = pHead1
        } else {
            cur2 = cur2.Next
        }
    }
    return cur1
}
