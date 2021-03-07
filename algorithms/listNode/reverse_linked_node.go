package listNode

import "fmt"

// https://leetcode-cn.com/problems/reverse-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var result string
	node := l
	for node != nil {
		result += fmt.Sprintf("%d-", node.Val)
		node = node.Next
	}
	return result
}

func ReverseList(head *ListNode) *ListNode {
	var curNode, preNode *ListNode
	curNode = head
	preNode = nil
	for curNode != nil {
		nextNode := curNode.Next
		curNode.Next = preNode
		preNode = curNode
		curNode = nextNode
	}
	return preNode
}
