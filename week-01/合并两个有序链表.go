//https://leetcode-cn.com/problems/merge-two-sorted-lists/

/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
示例 1：

输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
示例 2：

输入：l1 = [], l2 = []
输出：[]
示例 3：

输入：l1 = [], l2 = [0]
输出：[0]

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    //声明一个新节点用于待返回的链表头
    newLists := &ListNode{}
    //使用一个中间变量来操作节点，使新链表有序
    cur := newLists
    //如果l1和l2都不为nil那么依次比较，谁小就用谁，直到为nil
    for l1 != nil && l2 != nil {
        if l1.Val >= l2.Val {
            cur.Next = l2
            l2 = l2.Next
        }else{
            cur.Next = l1
            l1 = l1.Next
        }
        cur = cur.Next
    }
    //如果l1不为nil，把cur连接到l1中剩下的节点
    if l1 != nil {
        cur.Next = l1
    }
    //如果l2不为nil，把cur连接到l2中剩下的节点
    if l2 != nil {
        cur.Next = l2
    }

    return newLists.Next

}