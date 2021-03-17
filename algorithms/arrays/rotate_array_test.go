package arrays

import (
	"reflect"
	"testing"
)

func TestRotateByCopy(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
		k    int
	}{
		{
			"case1",
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{6, 7, 1, 2, 3, 4, 5},
			2,
		},
		{
			"case2",
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{5, 6, 7, 1, 2, 3, 4},
			3,
		},
		{
			"case3",
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{4, 5, 6, 7, 1, 2, 3},
			4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if RotateByCopy(test.nums, test.k); !reflect.DeepEqual(test.want, test.nums) {
				t.Fatal(test)
			}
		})
	}
}

func TestRotateBySwap(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
		k    int
	}{
		{
			"case1",
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{6, 7, 1, 2, 3, 4, 5},
			2,
		},
		{
			"case2",
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{5, 6, 7, 1, 2, 3, 4},
			3,
		},
		{
			"case3",
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{4, 5, 6, 7, 1, 2, 3},
			4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if RotateBySwap(test.nums, test.k); !reflect.DeepEqual(test.want, test.nums) {
				t.Fatal(test)
			}
		})
	}
}
