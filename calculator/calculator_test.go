package calculator_test

import (
	"golearn/calculator"
	"testing"
)

func TestCalculator(t *testing.T) {
	if calculator.Add(1, 2) != 3 {
		t.Fatal("Adding 1 and 2 should give 3")
	}

	if calculator.Sub(4, 2) != 2 {
		t.Fatal("Subtracting 2 from 4 should give 2")
	}

	if calculator.Mul(3, 2) != 6 {
		t.Fatal("Multiplying 3 to 2 should give 6")
	}

	if calculator.Div(6, 2) != 3 {
		t.Fatal("Dividing 6 by 2 should give 3")
	}
}
