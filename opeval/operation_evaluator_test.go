package opeval

import "testing"

func TestEvaluate(t *testing.T) {
	var eq string
	var res float64
	var expRes float64
	var err error

	t.Log("Simple operations")
	eq = "1 + 2"
	expRes = float64(1) + float64(2)
	res, _ = Evaluate(eq)
	if res != expRes {
		t.Error("Expected ", expRes, ", got ", res)
	}
	eq = "3 * 4"
	expRes = float64(3) * float64(4)
	res, _ = Evaluate(eq)
	if res != expRes {
		t.Error("Expected ", expRes, ", got ", res)
	}
	eq = "5 - 6"
	expRes = float64(5) - float64(6)
	res, _ = Evaluate(eq)
	if res != expRes {
		t.Error("Expected ", expRes, ", got ", res)
	}

	eq = "7 / 8"
	expRes = float64(7) / float64(8)
	res, _ = Evaluate(eq)
	if res != expRes {
		t.Error("Expected ", expRes, ", got ", res)
	}

	t.Log("Simple operations with operands of length greater than 2")
	eq = "10/11"
	expRes = float64(10) / float64(11)
	res, _ = Evaluate(eq)
	if res != expRes {
		t.Error("Expected ", expRes, ", got ", res)
	}
	eq = "81 + 129999 - 0900"
	expRes = float64(81) + float64(129999) - float64(900)
	res, _ = Evaluate(eq)
	if res != expRes {
		t.Error("Expected ", expRes, ", got ", res)
	}

	t.Log("Simple operations with decimal operands")
	eq = "1.1*0.8"
	expRes = float64(1.1) * float64(0.8)
	res, _ = Evaluate(eq)
	if res != expRes {
		t.Error("Expected ", expRes, ", got ", res)
	}
	eq = "1/3 - 0.333333333333"
	expRes = float64(1)/float64(3) - float64(0.333333333333)
	res, _ = Evaluate(eq)
	if res != expRes {
		t.Error("Expected ", expRes, ", got ", res)
	}
	eq = "2. + 8."
	expRes = float64(2.) + float64(8.)
	res, _ = Evaluate(eq)
	if res != expRes {
		t.Error("Expected ", expRes, ", got ", res)
	}

	t.Log("Complex operation")
	eq = "(10.8 * (1.2 - 0.87) / 14) * 10"
	expRes = (float64(10.8) * (float64(1.2) - float64(0.87)) / float64(14)) * float64(10)
	res, _ = Evaluate(eq)
	if res != expRes {
		t.Error("Expected ", expRes, ", got ", res)
	}
	eq = "((((1500)+127)*900)-1)"
	expRes = ((((float64(1500)) + float64(127)) * float64(900)) - float64(1))
	res, _ = Evaluate(eq)
	if res != expRes {
		t.Error("Expected ", expRes, ", got ", res)
	}
	eq = "(1500+(127*(900-(1))))"
	expRes = (float64(1500) + (float64(127) * (float64(900) - (float64(1)))))
	res, _ = Evaluate(eq)
	if res != expRes {
		t.Error("Expected ", expRes, ", got ", res)
	}

	t.Log("Wrong parenthesis configuration")
	eq = "(1+2"
	res, err = Evaluate(eq)
	if err == nil {
		t.Error("Expected error, got ", res)
	}
	eq = "(1+2))"
	res, err = Evaluate(eq)
	if err == nil {
		t.Error("Expected error, got ", res)
	}
	eq = ")(1+2)("
	res, err = Evaluate(eq)
	if err == nil {
		t.Error("Expected error, got ", res)
	}

	t.Log("Division by 0")
	eq = "1 / 0"
	res, err = Evaluate(eq)
	if err == nil {
		t.Error("Expected error, got ", res)
	}
}
