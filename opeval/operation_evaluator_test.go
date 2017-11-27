package opeval

import "testing"

type testpair struct {
	eq     string
	expRes float64
}

func TestSimpleEvaluation(t *testing.T) {
	var tests = []testpair{
		{"1 + 2",
			float64(1) + float64(2)},
		{"3 * 4",
			float64(3) * float64(4)},
		{"5 - 6",
			float64(5) - float64(6)},
		{"7 / 8",
			float64(7) / float64(8)},
	}

	for _, test := range tests {
		res, _ := Evaluate(test.eq)
		if res != test.expRes {
			t.Error("Expected ", test.expRes, ", got ", res)
		}
	}
}

func TestLongEvaluation(t *testing.T) {
	var tests = []testpair{
		{"10/11",
			float64(10) / float64(11)},
		{"81 + 129999 - 0900",
			float64(81) + float64(129999) - float64(900)},
	}

	for _, test := range tests {
		res, _ := Evaluate(test.eq)
		if res != test.expRes {
			t.Error("Expected ", test.expRes, ", got ", res)
		}
	}
}

func TestDecimalEvaluation(t *testing.T) {
	var tests = []testpair{
		{"1.1*0.8",
			float64(1.1) * float64(0.8)},
		{"1/3 - 0.333333333333",
			float64(1)/float64(3) - float64(0.333333333333)},
		{"2. + 8.",
			float64(2.) + float64(8.)},
	}

	for _, test := range tests {
		res, _ := Evaluate(test.eq)
		if res != test.expRes {
			t.Error("Expected ", test.expRes, ", got ", res)
		}
	}
}

func TestComplexEvaluation(t *testing.T) {
	var tests = []testpair{
		{"(10.8 * (1.2 - 0.87) / 14) * 10",
			(float64(10.8) * (float64(1.2) - float64(0.87)) / float64(14)) * float64(10)},
		{"((((1500)+127)*900)-1)",
			((((float64(1500)) + float64(127)) * float64(900)) - float64(1))},
		{"(1500+(127*(900-(1))))",
			(float64(1500) + (float64(127) * (float64(900) - (float64(1)))))},
	}

	for _, test := range tests {
		res, _ := Evaluate(test.eq)
		if res != test.expRes {
			t.Error("Expected ", test.expRes, ", got ", res)
		}
	}
}

func TestParenthesisError(t *testing.T) {
	var tests = []string{
		"(1+2",
		"(1+2))",
		")(1+2)(",
	}
	for _, test := range tests {
		res, err := Evaluate(test)
		if err == nil {
			t.Error("Expected error, got ", res)
		}
	}
}

func TestDivisionByZeroEvaluation(t *testing.T) {
	eq := "1 / 0"
	res, err := Evaluate(eq)
	if err == nil {
		t.Error("Expected error, got ", res)
	}
}
