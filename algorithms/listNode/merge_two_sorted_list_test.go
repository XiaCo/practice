package listNode_test

import (
	"fmt"
	"testing"

	"github.com/XiaCo/practice/algorithms/listNode"
)

func TestMergeTwoLists(t *testing.T) {
	elem1 := listNode.ListNode{-2, nil}
	elem2 := listNode.ListNode{-1, nil}
	elem3 := listNode.ListNode{3, nil}
	elem4 := listNode.ListNode{4, nil}
	elem1.Next = &elem2
	elem2.Next = &elem3
	elem3.Next = &elem4

	e1 := listNode.ListNode{1, nil}
	e2 := listNode.ListNode{2, nil}
	e3 := listNode.ListNode{3, nil}
	e1.Next = &e2
	e2.Next = &e3
	if result := fmt.Sprintf("%s", listNode.MergeTwoLists(&elem1, &e1)); result != "-2--1-1-2-3-3-4-" {
		t.Fatal(result)
	}
}
