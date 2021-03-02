package arrays

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 2, 3, 0, 0, 4, 5}
	MoveZeroes(nums)
	if !reflect.DeepEqual(nums, []int{1, 2, 3, 4, 5, 0, 0, 0, 0}) {
		t.Fatal(nums)
	}
}
