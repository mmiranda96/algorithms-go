package queue

import (
	"testing"
)

func TestInit(t *testing.T) {
	q := &Queue{}
	t.Log("Initializing stack...")
	q.Init()
}

func TestAdd(t *testing.T) {
	q := &Queue{}
	q.Init()
	t.Log("Enqueueing elements...")
	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}
}

func TestDequeue(t *testing.T) {
	q := &Queue{}
	q.Init()
	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}

	t.Log("Dequeueing elements...")
	for i := 0; i < 10; i++ {
		v, e := q.Dequeue()
		if v != i {
			t.Error("Expected ", i, ", got ", v)
		}
		if e != nil {
			t.Error("Expected no error, got ", e)
		}
	}

	t.Log("Dequeueing from an empty stack...")
	v, e := q.Dequeue()
	if e == nil {
		t.Error("Expected error, got ", v)
	}
}

func TestEnqueueAndDequeue(t *testing.T) {
	q := &Queue{}
	q.Init()
	t.Log("Enqueueing and dequeueing elements...")
	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}

	for i := 0; i < 3; i++ {
		q.Dequeue()
	}

	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}

	for i := 0; i < 2; i++ {
		q.Dequeue()
	}

	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}

	qString := q.String()
	if qString != "[0, 1, 2, 3, 4, 0, 1, 2, 3, 4]" {
		t.Error("Expected [0, 1, 2, 3, 4, 0, 1, 2, 3, 4], got ", qString)
	}
}

func TestPeek(t *testing.T) {
	q := &Queue{}
	q.Init()
	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}
	
	t.Log("Peeking elements...")
	v, e := q.Peek()
	if v != 0 {
		t.Error("Expected 9, got ", v)
	}
	if e != nil {
		t.Error("Expected no error, got ", e)
	}
	
	for i := 0; i < 10; i++ {
		q.Dequeue()
	}

	t.Log("Peeking from an empty stack...")
	v, e = q.Peek()
	if e == nil {
		t.Error("Expected error, got ", v)
	}
}

func TestString(t *testing.T) {
	q := &Queue{}
	q.Init()
	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}
	qString := q.String()
	if qString != "[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]" {
		t.Error("Expected [0, 1, 2, 3, 4, 5, 6, 7, 8, 9], got ", qString)
	}
}