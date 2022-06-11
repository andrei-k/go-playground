package main

import "testing"

func TestCalc(t *testing.T) {
	if Calc(10) != 20 {
		t.Error("Expected 20, got ", Calc(10))
	}
}

func TestTableCalc(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{10, 20},
		{-1, 9},
		{0, 10},
		{990, 1000},
	}

	for _, test := range tests {
		if output := Calc(test.input); output != test.expected {
			t.Error("Test failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
