package arrays

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	result := ThreeSum([]int{-4, -1, -1, 0, 1, 2})
	want := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	if !reflect.DeepEqual(result, want) {
		t.Fatal(result)
	}
}

func TestTwoSumByMap(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want [][]int
	}{
		{
			name: "test1",
			in:   []int{5, 12, 6, 3, 9, 2, 1, 7},
			want: [][]int{{1, 12}, {7, 6}}, // attend the order
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := TwoSumByMap(test.in)
			for index, item := range result {
				if !reflect.DeepEqual(item, test.want[index]) {
					t.Fatal("error", result)
				}
			}
		})
	}
}
