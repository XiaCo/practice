package data_structures

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestBitArray(t *testing.T) {
	length := 100000
	rand.Seed(time.Now().Unix())
	ba := NewBitArray(uint64(length))
	normalArray := make([]bool, length)

	for i := 0; i < length; i++ {
		if v := rand.Uint32() & 1; v == 1 {
			normalArray[i] = true
			if res := ba.SetBit1(uint64(i)); !res {
				t.Fatal("error set bit")
			}
		} else {
			normalArray[i] = false
			if res := ba.SetBit0(uint64(i)); !res {
				t.Fatal("error set bit")
			}
		}
	}

	for i := 0; i < length; i++ {
		if ba.GetBit(uint64(i)) != normalArray[i] {
			t.Fatal("error check")
		}
	}
}

func BenchmarkBitArrayWrite(b *testing.B) {
	length := int(math.Exp2(18)) - 1
	ba := NewBitArray(uint64(length))
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			i++
			ba.SetBit1(uint64(i & length))
		}
	})
}

func BenchmarkBitArrayRead(b *testing.B) {
	length := int(math.Exp2(18)) - 1
	ba := NewBitArray(uint64(length))
	//for i := 0; i < length; i++ {
	//	ba.SetBit1(uint64(rand.Uint32()&uint32(length)))
	//}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			i++
			ba.GetBit(uint64(i & length))
		}
	})
}
