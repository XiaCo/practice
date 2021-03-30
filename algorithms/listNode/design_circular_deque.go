package listNode

import "fmt"

// https://leetcode-cn.com/problems/design-circular-deque/

type PPListNode struct {
	Val  int
	Next *PPListNode
	Pre  *PPListNode
}

func (l *PPListNode) String() string {
	var result string
	node := l
	for node != nil {
		result += fmt.Sprintf("%d-", node.Val)
		node = node.Next
	}
	return result
}

type MyCircularDeque struct {
	head, tail *PPListNode
	length     int
	maxLength  int
}

func (this MyCircularDeque) String() string {
	return fmt.Sprint(this.head)
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	if k <= 0 {
		panic("k must gt than 0")
	}
	q := MyCircularDeque{
		nil,
		nil,
		0,
		k,
	}
	return q
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.length == this.maxLength {
		return false
	}
	newHead := &PPListNode{Val: value}
	if this.head == nil {
		this.head = newHead
		this.tail = newHead
	} else {
		// 连接新旧head
		this.head.Pre = newHead
		newHead.Next = this.head
		// 替换新head
		this.head = newHead
	}
	this.length++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.length == this.maxLength {
		return false
	}
	newTail := &PPListNode{Val: value}
	if this.tail == nil {
		this.head = newTail
		this.tail = newTail
	} else {
		// 连接新旧tail
		this.tail.Next = newTail
		newTail.Pre = this.tail
		// 替换新tail
		this.tail = newTail
	}
	this.length++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.head == nil {
		return false
	}
	if this.head == this.tail {
		this.head = nil
		this.tail = nil
	} else {
		temp := this.head.Next
		// 断开连接
		this.head.Next.Pre = nil
		this.head.Next = nil
		// 替换新head
		this.head = temp
	}
	this.length--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.tail == nil {
		return false
	}
	if this.head == this.tail {
		this.head = nil
		this.tail = nil
	} else {
		temp := this.tail.Pre
		// 断开连接
		this.tail.Pre.Next = nil
		this.tail.Pre = nil
		// 替换新tail
		this.tail = temp
	}
	this.length--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.head == nil {
		return -1
	}
	return this.head.Val
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.tail == nil {
		return -1
	}
	return this.tail.Val
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.head == nil
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.length == this.maxLength
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
