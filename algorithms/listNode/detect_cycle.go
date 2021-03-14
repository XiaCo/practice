package listNode

// https://leetcode-cn.com/problems/linked-list-cycle-ii/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func DetectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	m := make(map[*ListNode]struct{})
	node := head
	for node != nil {
		if _, exist := m[node]; exist {
			return node
		} else {
			m[node] = struct{}{}
		}
		node = node.Next
	}
	return nil
}
