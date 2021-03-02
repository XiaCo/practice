package data_structures

import (
	"fmt"
	"testing"
)

func TestNewLFUv1(t *testing.T) {
	lfu := NewLFUv1(2)
	lfu.Set("1", 1)
	lfu.Set("2", 2)
	if val := lfu.Get("2"); val != 2 {
		t.Fatal(val)
	}
	lfu.Set("3", 3)
	if val := lfu.Get("1"); val != nil {
		t.Fatal(val)
	}
	if val := lfu.Get("3"); val != 3 {
		t.Fatal(val)
	}
	if val := lfu.Get("3"); val != 3 {
		t.Fatal(val)
	}
}

func TestString(t *testing.T) {
	lfu := NewLFUv1(5)
	lfu.Set(1, 1)
	lfu.Set(2, 2)
	lfu.Set(3, 3)
	fmt.Println(lfu)
	lfu.Set(4, 4)
	fmt.Println(lfu)
	lfu.Set(5, 5)
	fmt.Println(lfu)
	lfu.Set(6, 6)
	fmt.Println(lfu)
}

func BenchmarkSet(b *testing.B) {
	lfu := NewLFUv1(1024)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		n := 0
		for pb.Next() {
			lfu.Set(n, n)
			n++
		}
	})
}

func BenchmarkGet(b *testing.B) {
	lfu := NewLFUv1(1500)
	for i := 0; i < 1500; i++ {
		lfu.Set(i, i)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		n := 0
		for pb.Next() {
			lfu.Get(n & 1023)
			n++
		}
	})
}
