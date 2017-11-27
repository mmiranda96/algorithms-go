package stack

import (
	"testing"
)

func TestInit(t *testing.T) {
	s := &Stack{}
	t.Log("Initializing stack...")
	s.Init()
}

func TestAdd(t *testing.T) {
	s := &Stack{}
	s.Init()
	t.Log("Pushing elements...")
	for i := 0; i < 10; i++ {
		s.Push(i)
	}
}

func TestPop(t *testing.T) {
	s := &Stack{}
	s.Init()
	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	t.Log("Popping elements...")
	for i := 0; i < 10; i++ {
		v, e := s.Pop()
		if v != 9-i {
			t.Error("Expected ", 9-i, ", got ", v)
		}
		if e != nil {
			t.Error("Expected no error, got ", e)
		}
	}

	t.Log("Popping from an empty stack...")
	v, e := s.Pop()
	if e == nil {
		t.Error("Expected error, got ", v)
	}
}

func TestPushAndPop(t *testing.T) {
	s := &Stack{}
	s.Init()
	t.Log("Enqueueing and dequeueing elements...")
	for i := 0; i < 5; i++ {
		s.Push(i)
	}

	for i := 0; i < 3; i++ {
		s.Pop()
	}

	for i := 0; i < 5; i++ {
		s.Push(i)
	}

	for i := 0; i < 2; i++ {
		s.Pop()
	}

	for i := 0; i < 5; i++ {
		s.Push(i)
	}

	sString := s.String()
	if sString != "[4, 3, 2, 1, 0, 2, 1, 0, 1, 0]" {
		t.Error("Expected [4, 3, 2, 1, 0, 2, 1, 0, 1, 0], got ", sString)
	}
}

func TestPeek(t *testing.T) {
	s := &Stack{}
	s.Init()
	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	t.Log("Peeking elements...")
	v, e := s.Peek()
	if v != 9 {
		t.Error("Expected 9, got ", v)
	}
	if e != nil {
		t.Error("Expected no error, got ", e)
	}

	for i := 0; i < 10; i++ {
		s.Pop()
	}

	t.Log("Peeking from an empty stack...")
	v, e = s.Peek()
	if e == nil {
		t.Error("Expected error, got ", v)
	}
}

func TestString(t *testing.T) {
	s := &Stack{}
	s.Init()
	for i := 0; i < 10; i++ {
		s.Push(i)
	}
	sString := s.String()
	if sString != "[9, 8, 7, 6, 5, 4, 3, 2, 1, 0]" {
		t.Error("Expected [9, 8, 7, 6, 5, 4, 3, 2, 1, 0], got ", sString)
	}
}
