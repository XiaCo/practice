package listNode

// https://leetcode-cn.com/problems/swap-nodes-in-pairs/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tail := &ListNode{} // 尾节点变量用来记录当前位置
	var newHead, curNode, nextNode *ListNode
	curNode = head
	nextNode = head.Next
	newHead = nextNode
	for nextNode != nil {
		tail.Next = nextNode
		tempNode := nextNode.Next
		nextNode.Next = curNode
		tail = curNode
		tail.Next = tempNode

		curNode = tempNode
		if curNode == nil {
			return newHead
		} else {
			nextNode = tempNode.Next
		}
	}

	return newHead
}
