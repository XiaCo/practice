package data_structures

import (
	"math"
	"runtime"
	"sync/atomic"
)

type BitArray struct {
	max   uint64
	array []uint32
}

func GetBitArrayMemCostByte(b *BitArray) uint32 {
	return uint32(math.Ceil(float64(b.max) / float64(32)))
}

func setBit0(i uint32, oldNum uint32) (newNum uint32) {
	// i is index in bits
	i = 31 - i
	return oldNum & ^(1 << i)
}

func setBit1(i, oldNum uint32) (newNum uint32) {
	i = 31 - i
	return (1 << i) | oldNum
}

func (b *BitArray) SetBit1(index uint64) bool {
	if index > b.max {
		return false
	}
	i := index & 31
	multiple := index - i
	for {
		oldValue := atomic.LoadUint32(&b.array[multiple/32])
		if atomic.CompareAndSwapUint32(&b.array[multiple/32], oldValue, setBit1(uint32(i), oldValue)) {
			return true
		}
		runtime.Gosched() // TODO: return false or spin ?
	}
}

func (b *BitArray) SetBit0(index uint64) bool {
	if index > b.max {
		return false
	}
	i := index & 31
	multiple := index - i
	for {
		oldValue := atomic.LoadUint32(&b.array[multiple/32])
		if atomic.CompareAndSwapUint32(&b.array[multiple/32], oldValue, setBit0(uint32(i), oldValue)) {
			return true
		}
		runtime.Gosched()
	}
}

func (b *BitArray) GetBit(index uint64) bool {
	if index > b.max {
		return false
	}
	i := index & 31
	oldValue := atomic.LoadUint32(&b.array[(index-i)/32])
	return ((1 << (31 - i)) | oldValue) == oldValue // lagging result
}

func NewBitArray(max uint64) *BitArray {
	return &BitArray{
		max,
		make([]uint32, int(math.Ceil(float64(max)/float64(32)))),
	}
}
