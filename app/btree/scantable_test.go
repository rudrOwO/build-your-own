package btree

import "testing"

// Dummy Test Function
func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}
