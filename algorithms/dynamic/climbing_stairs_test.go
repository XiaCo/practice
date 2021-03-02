package dynamic

import "testing"

func TestClimbStairs(t *testing.T) {
	if value := ClimbStairs(1); value != 1 {
		t.Fatal(value)
	}
	if value := ClimbStairs(2); value != 2 {
		t.Fatal(value)
	}
	if value := ClimbStairs(3); value != 3 {
		t.Fatal(value)
	}
	if value := ClimbStairs(4); value != 5 {
		t.Fatal(value)
	}
	if value := ClimbStairs(5); value != 8 {
		t.Fatal(value)
	}
}

func TestClimbStairs2(t *testing.T) {
	if value := ClimbStairs2(1); value != 1 {
		t.Fatal(value)
	}
	if value := ClimbStairs2(2); value != 2 {
		t.Fatal(value)
	}
	if value := ClimbStairs2(3); value != 3 {
		t.Fatal(value)
	}
	if value := ClimbStairs2(4); value != 5 {
		t.Fatal(value)
	}
	if value := ClimbStairs2(5); value != 8 {
		t.Fatal(value)
	}
}
