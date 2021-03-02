package data_structures

import (
	"fmt"
	"sync"
)

type elem struct {
	index int // 当前数组中的位置
	hit   int // 被访问次数
	key   interface{}
	value interface{}
}

func (e *elem) swap(li *lfuArray) {
	if e.index == 0 {
		return
	}
	if li.array[e.index].hit > li.array[e.index-1].hit { // 与前一个元素换位
		li.array[e.index], li.array[e.index-1] = li.array[e.index-1], li.array[e.index]
		e.index--
		li.array[e.index+1].index++
		e.swap(li)
	}
}

type lfuArray struct {
	len   int
	array []*elem
}

func newlfuArray(c int) *lfuArray {
	la := &lfuArray{0, make([]*elem, c)}
	for i := 0; i < cap(la.array); i++ {
		la.array[i] = &elem{i, 0, nil, nil}
	}
	return la
}

type lfuMap map[interface{}]*elem

type LFUv1 struct {
	*lfuArray
	lfuMap
	lock *sync.Mutex
}

func (L *LFUv1) Set(key, value interface{}) {
	L.lock.Lock()
	defer L.lock.Unlock()
	// 替换数组中的元素
	if L.len == cap(L.array) { // 数组已满，替换map与array尾部
		delete(L.lfuMap, L.array[L.len-1].key)
		L.array[L.len-1].hit = 0
		L.array[L.len-1].key = key
		L.array[L.len-1].value = value
		L.lfuMap[key] = L.array[L.len-1]
	} else {
		L.array[L.len].key = key
		L.array[L.len].value = value
		L.lfuMap[key] = L.array[L.len]
		L.len++
	}
}

func (L *LFUv1) Get(key interface{}) interface{} {
	L.lock.Lock()
	defer L.lock.Unlock()
	if val, exist := L.lfuMap[key]; exist {
		val.hit++
		val.swap(L.lfuArray)
		return val.value
	} else {
		return nil
	}
}

func (L *LFUv1) String() string {
	var result string
	if L.len > 4 {
		result = fmt.Sprintf("(%v: %v) -- (%v: %v) ... (%v: %v) -- (%v: %v)",
			L.array[0].key, L.array[0].value, L.array[1].key, L.array[1].value,
			L.array[L.len-2].key, L.array[L.len-2].value, L.array[L.len-1].key, L.array[L.len-1].value)
	} else {
		for i := 0; i < L.len; i++ {
			result += fmt.Sprintf("(%v: %v)", L.array[i].key, L.array[i].value)
			if i != L.len-1 {
				result += " -- "
			}
		}
	}
	return result
}

func NewLFUv1(cap int) *LFUv1 {
	return &LFUv1{
		newlfuArray(cap),
		make(map[interface{}]*elem, cap),
		new(sync.Mutex),
	}
}
