package data_structures

import (
	"testing"
)

func TestLRUv1_Get(t *testing.T) {
	l := NewLRUv1(2)
	if val := l.Get("1"); val != nil {
		t.Fatal(val)
	}
	l.Set("1", 1)
	l.Set("2", 2)
	l.Set("3", 3)
	if val := l.Get("1"); val != nil {
		t.Fatal(val)
	}
	if val := l.Get("2"); val != 2 {
		t.Fatal(val)
	}
	l.Set("4", 4)
	if val := l.Get("3"); val != nil {
		t.Fatal(val)
	}
}

func BenchmarkLRUv1_Get(b *testing.B) {
	lru := NewLRUv1(1500)
	for i := 0; i < 1500; i++ {
		lru.Set(i, i)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		n := 0
		for pb.Next() {
			lru.Get(n & 1023)
			n++
		}
	})
}

func BenchmarkLRUv1_Set(b *testing.B) {
	lru := NewLRUv1(1024)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		n := 0
		for pb.Next() {
			lru.Set(n, n)
			n++
		}
	})
}
