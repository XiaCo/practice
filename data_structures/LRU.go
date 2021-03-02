package data_structures

import (
	"container/list"
	"sync"
)

type dot struct {
	key   interface{}
	value interface{}
}

type LRUv1 struct {
	len int
	cap int
	*list.List
	m map[interface{}]*list.Element
	sync.Locker
}

func (L *LRUv1) Get(key interface{}) interface{} {
	L.Lock()
	defer L.Unlock()
	if elem, exist := L.m[key]; exist {
		L.MoveToFront(elem)
		return elem.Value.(*dot).value
	}
	return nil
}

func (L *LRUv1) Set(key, value interface{}) {
	L.Lock()
	defer L.Unlock()
	if elem, exist := L.m[key]; exist { // 更新值，并放置到头部
		elem.Value.(*dot).value = value
		L.MoveToFront(elem)
		return
	}
	if L.len != L.cap {
		head := L.PushFront(&dot{key, value})
		L.m[key] = head
		L.len++
	} else {
		oldKey := L.Back().Value.(*dot).key
		delete(L.m, oldKey)
		tailDot := L.Back().Value.(*dot)
		tailDot.key = key
		tailDot.value = value
		L.MoveToFront(L.Back())
		L.m[key] = L.Front()
	}
}

func NewLRUv1(cap int) *LRUv1 {
	return &LRUv1{
		0,
		cap,
		list.New(),
		make(map[interface{}]*list.Element, cap),
		new(sync.Mutex),
	}
}
