// Package opeval implements a basic operation evaluator.
//
// The evaluator takes a numeric infix operation and returns its result.
// It uses Shunting-yard algorithm to pass from infix to postfix, and then
// the postfix operation is evaluated.
package opeval

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/mmiranda96/algorithms-go/queue"
	"github.com/mmiranda96/algorithms-go/stack"
)

func checkParenthesis(eq string, status int) bool {
	if eq == "" {
		return status == 0
	} else if eq[0] == '(' {
		return checkParenthesis(eq[1:len(eq)], status+1)
	} else if eq[0] == ')' {
		if status == 0 {
			return false
		}
		return checkParenthesis(eq[1:len(eq)], status-1)
	}
	return checkParenthesis(eq[1:len(eq)], status)
}

func isNumber(x string) bool {
	_, res := strconv.Atoi(x)
	return res == nil
}

func isOperator(x string) bool {
	return x == "+" || x == "-" || x == "*" || x == "/" || x == "(" || x == ")"
}

func isPrecedent(x string, y string) bool {
	if x == "+" || x == "-" {
		return y == "+" || y == "-"
	}
	return true
}

func toPostfix(eq string) *queue.Queue {
	q := &queue.Queue{}
	s := &stack.Stack{}

	q.Init()
	s.Init()

	var v string
	for i := 0; i < len(eq); i++ {
		t := string(eq[i])
		if isOperator(t) {
			if v != "" {
				q.Enqueue(v)
				v = ""
			}

			if t == "(" {
				s.Push(t)
			} else if t == ")" {
				for op, _ := s.Peek(); op != "("; op, _ = s.Peek() {
					op, _ = s.Pop()
					q.Enqueue(op)
				}
				s.Pop()
			} else {
				for op, _ := s.Peek(); !s.IsEmpty() && op != "(" && isPrecedent(fmt.Sprintf("%v", op), t); op, _ = s.Peek() {
					op, _ = s.Pop()
					q.Enqueue(op)
				}
				s.Push(t)
			}
		} else if isNumber(t) || t == "." {
			v += t
		}
	}
	if v != "" {
		q.Enqueue(v)
	}
	for !s.IsEmpty() {
		op, _ := s.Pop()
		q.Enqueue(op)
	}
	return q
}

func eval(a float64, b float64, op string) (float64, error) {
	if op == "+" {
		return a + b, nil
	} else if op == "-" {
		return a - b, nil
	} else if op == "*" {
		return a * b, nil
	}
	if b == 0 {
		return 0.0, errors.New("opeval - division by 0")
	}
	return a / b, nil
}

func toValue(x interface{}) (float64, error) {
	return strconv.ParseFloat(fmt.Sprintf("%v", x), 64)
}

// Evaluate receives an infix operation and returns its evaluation.
// It ignores non-operand/operators (anything else than numbers, dots,
// parenthesis or basic operators). Valid operators are '+', '-', '*' and '/'.
func Evaluate(eq string) (float64, error) {
	if !checkParenthesis(eq, 0) {
		return 0.0, errors.New("opeval - invalid parenthesis configuration")
	}

	q := toPostfix(eq)
	s := &stack.Stack{}
	s.Init()
	for !q.IsEmpty() {
		x, _ := q.Dequeue()
		op := fmt.Sprintf("%v", x)
		if isOperator(op) {
			b, err := s.Pop()
			if err != nil {
				return 0.0, errors.New("opeval - invalid operation")
			}
			a, err := s.Pop()
			if err != nil {
				return 0.0, errors.New("opeval - invalid operation")
			}
			bnum, _ := toValue(b)
			anum, _ := toValue(a)
			t, err := eval(anum, bnum, op)
			if err != nil {
				return 0.0, errors.New("opeval - division by 0")
			}
			s.Push(t)
		} else {
			s.Push(op)
		}
	}
	if s.IsEmpty() {
		return 0.0, errors.New("opeval - invalid operation")
	}
	res, _ := s.Pop()
	return toValue(res)
}
