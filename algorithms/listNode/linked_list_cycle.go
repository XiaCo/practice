package listNode

// https://leetcode-cn.com/problems/linked-list-cycle/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slowNode := head
	fastNode := head.Next
	for fastNode != nil && fastNode.Next != nil {
		if slowNode == fastNode {
			return true
		}
		slowNode = slowNode.Next
		fastNode = fastNode.Next.Next
	}
	return false
}
