package main

import (
	"testing"
)

type Test struct {
	arg1, arg2, expected int
}

func TestAdd(t *testing.T) {
	var addTests = []Test{
		Test{2, 3, 5},
		Test{4, 5, 9},
		Test{2, 5, 7},
		Test{2, 6, 8},
	}

	for _, test := range addTests {
		if output := Add(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func TestSub(t *testing.T) {
	var subTests = []Test{
		Test{5, 3, 2},
		Test{7, 5, 2},
		Test{9, 1, 8},
		Test{5, 5, 0},
	}

	for _, test := range subTests {
		if output := Sub(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Ouput %q not equal to expected %q", output, test.expected)
		}
	}
}