package list

import (
	"testing"
)

func TestInit(t *testing.T) {
	l := &List{}
	t.Log("Initializing list...")
	l.Init()
}

func TestAdd(t *testing.T) {
	l := &List{}

	l.Init()
	for i := 0; i < 10; i++ {
		e := l.Add(0, i)
		if e != nil {
			t.Error("Expected to add item at position 0, got error: ", e)
		}
	}
	if l.String() != "[9, 8, 7, 6, 5, 4, 3, 2, 1, 0]" {
		t.Error("Expected list to be [9, 8, 7, 6, 5, 4, 3, 2, 1, 0], got ", l.String())
	}

	l.Init()
	t.Log("Adding elements at end...")
	for i := 0; i < 10; i++ {
		e := l.Add(i, i)
		if e != nil {
			t.Error("Expected to add item at position 0, got error: ", e)
		}
	}
	if l.String() != "[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]" {
		t.Error("Expected list to be [0, 1, 2, 3, 4, 5, 6, 7, 8, 9], got ", l.String())
	}

	l.Init()
	t.Log("Adding elements at invalid index...")
	e := l.Add(-1, -1)
	if e == nil {
		t.Error("Expected error")
	}
	e = l.Add(10, 10)
	if e == nil {
		t.Error("Expected error")
	}
}

func TestGet(t *testing.T) {
	l := &List{}

	l.Init()
	for i := 0; i < 10; i++ {
		l.Add(i, i)
	}
	t.Log("Getting elements in valid indexes...")
	for i := 0; i < 10; i++ {
		v, e := l.Get(i)
		if v != i {
			t.Error("Expected ", i, ", got ", v)
		}
		if e != nil {
			t.Error("Expected no error, got ", e)
		}
	}

	t.Log("Getting elements in invalid indexes...")
	v, e := l.Get(10)
	if e == nil {
		t.Error("Expected error, got ", v)
	} 
	v, e = l.Get(-1)
	if e == nil {
		t.Error("Expected error, got ", v)
	}
}

func TestDelete(t *testing.T) {
	l := &List{}

	l.Init()
	for i := 0; i < 10; i++ {
		l.Add(i, i)
	}
	t.Log("Deleting elements in valid indexes...")
	for i := 0; i < 5; i++ {
		v, e := l.Delete(9 - i)
		if v != 9 - i {
			t.Error("Expected ", i, ", got ", v)
		}
		if e != nil {
			t.Error("Expected no error, got ", e)
		}
	}
	if l.String() != "[0, 1, 2, 3, 4]" {
		t.Error("Expected list to be [0, 1, 2, 3, 4], got ", l.String())
	}

	t.Log("Deleting elements in invalid indexes...")
	v, e := l.Delete(10)
	if e == nil {
		t.Error("Expected error, got ", v)
	} 
	v, e = l.Delete(-1)
	if e == nil {
		t.Error("Expected error, got ", v)
	}
	
}

func TestString(t *testing.T) {
	l := &List{}
	l.Init()
	for i := 0; i < 10; i++ {
		l.Add(i, i)
	}
	lString := l.String()
	if lString != "[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]" {
		t.Error("Expected [0, 1, 2, 3, 4, 5, 6, 7, 8, 9], got ", lString)
	}
}