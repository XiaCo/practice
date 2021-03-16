package listNode

// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func ReverseKGroup(head *ListNode, k int) *ListNode {
	var newHead, tail, curNode *ListNode
	tail = &ListNode{}
	curNode = head
	for {
		stack := make([]*ListNode, k)
		for i := 0; i < k; i++ { // 入栈
			if curNode == nil { // 判断到尾部，将tail连接到尚未变换顺序的这组node
				if newHead == nil {
					return head
				}
				tail.Next = stack[0]
				return newHead
			}
			stack[i] = curNode
			curNode = curNode.Next
		}
		for i := k - 1; i >= 0; i-- { // 出栈，连接到tail
			tail.Next = stack[i]
			tail = tail.Next
		}
		if newHead == nil {
			newHead = stack[k-1]
		}
	}
}
