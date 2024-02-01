package main

import "testing"

func TestSum(t *testing.T) {

	result := Sum(2, 3)
	if result != 5 {
		t.Errorf("Sum(2, 3) = %d; want 5", result)
	}

	options := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{3, 4, 7},
	}

	for _, option := range options {
		result := Sum(option.a, option.b)
		if result != option.c {
			t.Errorf("Sum(%d, %d) = %d; want %d", option.a, option.b, result, option.c)
		}
	}
}

func TestGetMax(t *testing.T) {
	// Testing when a is greater than b
	if result := GetMax(5, 3); result != 5 {
		t.Errorf("GetMax(5, 3) = %d; want 5", result)
	}

	// Testing when b is greater than a
	if result := GetMax(2, 8); result != 8 {
		t.Errorf("GetMax(2, 8) = %d; want 8", result)
	}

	// Testing when a and b are equal
	if result := GetMax(4, 4); result != 4 {
		t.Errorf("GetMax(4, 4) = %d; want 4", result)
	}
}
