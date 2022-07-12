package math

import "testing"

// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'
type addTest struct {
	arg1, arg2, expected int
}

var addTests = []addTest{
	addTest{1, 2, 3},
	addTest{3, 4, 7},
	addTest{5, 6, 11},
}

type subTest struct {
	arg1, arg2, expected int
}

var subTests = []subTest{
	subTest{2, 1, 1},
	subTest{4, 3, 1},
	subTest{6, 5, 1},
}

func TestAdd(t *testing.T) {

	for _, test := range addTests {
		if output := Add(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func TestSubtract(t *testing.T) {

	for _, test := range subTests {
		if output := Subtract(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
