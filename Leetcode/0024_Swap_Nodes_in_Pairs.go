/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func swapPairs(head *ListNode) *ListNode {
    dummy := &ListNode{Next: head}
    pre := dummy

    for head != nil && head.Next != nil {
        first_node := head
        second_node := head.Next

        pre.Next = second_node
        first_node.Next = second_node.Next
        second_node.Next = first_node

        head = first_node.Next
        pre = first_node

    }
    return dummy.Next
}