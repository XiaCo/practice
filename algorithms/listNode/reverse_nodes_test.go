package listNode_test

import (
	"fmt"
	"github.com/XiaCo/practice/algorithms/listNode"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	elem1 := listNode.ListNode{1, nil}
	elem2 := listNode.ListNode{2, nil}
	elem3 := listNode.ListNode{3, nil}
	elem4 := listNode.ListNode{4, nil}
	elem1.Next = &elem2
	elem2.Next = &elem3
	elem3.Next = &elem4
	if result := fmt.Sprintf("%s", listNode.ReverseKGroup(&elem1, 3)); result != "3-2-1-4-" {
		t.Fatal(result)
	}
}
