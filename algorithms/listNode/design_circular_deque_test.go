package listNode

import (
	"fmt"
	"testing"
)

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

func TestMyCircularDeque(t *testing.T) {
	dq := Constructor(3)
	if result := dq.InsertFront(3); !result {
		t.Fatal(dq)
	}
	if result := dq.DeleteLast(); !result {
		t.Fatal(dq)
	}
	if val := dq.GetRear(); val != -1 {
		fmt.Println(dq)
		t.Fatal(dq)
	}

	if result := dq.InsertFront(3); !result {
		t.Fatal(dq)
	}
	if result := dq.InsertLast(2); !result {
		t.Fatal(dq)
	}

	if result := dq.InsertLast(4); !result {
		t.Fatal(dq)
	}
	if str := fmt.Sprint(dq); str != "3-2-4-" {
		fmt.Println(dq)
		t.Fatal(dq)
	}

	if result := dq.DeleteLast(); !result {
		t.Fatal(dq)
	}
	if str := fmt.Sprint(dq); str != "3-2-" {
		fmt.Println(dq)
		t.Fatal(dq)
	}

	if result := dq.DeleteFront(); !result {
		t.Fatal(dq)
	}
	if str := fmt.Sprint(dq); str != "2-" {
		fmt.Println(dq)
		t.Fatal(dq)
	}
}
