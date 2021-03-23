package calc

import (
	"github.com/filippocam/filonummanip/calc"
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 3
	result := calc.Add(1,2)
	if result != expected {
		t.Errorf("Expected %d got %d", expected, result) // to indicate test failed
	}
}
