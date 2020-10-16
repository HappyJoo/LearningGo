/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    cur := head
    var pre *ListNode = nil
    for cur != nil {
        temp := cur.Next
        cur.Next = pre
        pre = cur
        cur = temp
    }
    return pre
}