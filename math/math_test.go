package math

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Add(2, 3) = %d; expected 5", result)
	}
}

func Add(firstNum int, secondNum int) int {
	return firstNum + secondNum
}
