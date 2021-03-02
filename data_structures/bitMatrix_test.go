package data_structures

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func TestBitMatrix(t *testing.T) {
	// 131072 * 131072 cost 2GB
	length, width := int(math.Exp2(5)), int(math.Exp2(5))
	bm := NewBitMatrix(uint64(length), uint64(width))
	fmt.Printf("matrix cost %d Bytes\n", GetBitMatrixMemCostByte(bm))
	resArray := make([][2]uint64, 0, length)

	for i := 0; i < length; i++ {
		resArray = append(resArray, [2]uint64{rand.Uint64() % uint64(length-1), rand.Uint64() % uint64(width-1)})
		if ok := bm.SetBit1(resArray[i][0], resArray[i][1]); !ok {
			t.Errorf("set bit error")
		}
	}
	for i := 0; i < width; i++ {
		if ok := bm.GetBit(resArray[i][0], resArray[i][1]); !ok {
			t.Errorf("get bit error")
		}
	}
}
