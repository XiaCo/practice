package listNode

// https://leetcode-cn.com/problems/merge-two-sorted-lists/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	tail := &ListNode{}
	node1 := l1
	node2 := l2

	for node1 != nil && node2 != nil {
		var value1, value2 int
		value1 = node1.Val
		value2 = node2.Val
		if value1 < value2 {
			tail.Next = node1
			tail = node1
			node1 = node1.Next
		} else {
			tail.Next = node2
			tail = node2
			node2 = node2.Next
		}
		if head == nil {
			head = tail
		}
	}
	if node1 == nil {
		tail.Next = node2
	} else {
		tail.Next = node1
	}
	if head == nil {
		return tail.Next
	}
	return head
}
