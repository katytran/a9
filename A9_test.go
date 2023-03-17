package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Expected Add(2, 3) to be 5, but got %d", result)
	}
}