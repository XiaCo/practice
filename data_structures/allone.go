package data_structures

import (
	"container/list"
	"fmt"
)

// https://leetcode-cn.com/problems/all-oone-data-structure/

type one struct {
	key string
	num int
}

type AllOne struct {
	m      map[string]*list.Element
	l      *list.List
	maxKey string
	minKey string
}

func (a AllOne) String() string {
	var elem = a.l.Front()
	var res string
	for elem != nil {
		res += fmt.Sprintf("%s,%d -> ", elem.Value.(one).key, elem.Value.(one).num)
		elem = elem.Next()
	}
	res += "\n"
	return res
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	l := list.New()
	l.Init()
	return AllOne{
		make(map[string]*list.Element),
		l,
		"",
		"",
	}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	if val, exist := this.m[key]; exist {
		newVal := val.Value.(one).num + 1
		val.Value = one{key, newVal}
		for {
			if val.Next() != nil && val.Next().Value.(one).num < newVal {
				// 右移当前元素，直到末端或者小于下一个元素
				this.l.MoveAfter(val, val.Next())
			} else {
				break
			}
		}
		if this.l.Back() == val {
			this.maxKey = key
		}
	} else {
		this.m[key] = this.l.PushFront(one{key, 1})
	}
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	if val, exist := this.m[key]; exist {
		if val.Value.(one).num == 1 {
			delete(this.m, key)
			this.l.Remove(val)
		} else {
			newVal := val.Value.(one).num - 1
			val.Value = one{key, newVal}
			for {
				if val.Prev() != nil && val.Prev().Value.(one).num > newVal {
					// 左移当前元素
					this.l.MoveBefore(val, val.Prev())
				} else {
					break
				}
			}
			if this.l.Front() == val {
				this.minKey = key
			}
		}
	}
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	if e := this.l.Back(); e != nil {
		return e.Value.(one).key
	}
	return ""
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	if e := this.l.Front(); e != nil {
		return e.Value.(one).key
	}
	return ""
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
