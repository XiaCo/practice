package arrays

import "testing"

func TestMaxArea(t *testing.T) {
	if maxArea := MaxArea([]int{1,8,6,2,5,4,8,3,7}); maxArea != 49 {
		t.Fatal(maxArea)
	}
}
