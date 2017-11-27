// Package queue implements a basic queue.
//
// This queue is based on the two-stack implementation.
package queue

import (
	"errors"
	"fmt"

	"github.com/mmiranda96/algorithms-go/stack"
)

// Queue a two-stack queue
type Queue struct {
	r *stack.Stack
	s *stack.Stack
}

// Init inits an empty queue
func (q *Queue) Init() {
	r := &stack.Stack{}
	s := &stack.Stack{}
	r.Init()
	s.Init()
	q.r = r
	q.s = s
}

// IsEmpty returns if the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.r.IsEmpty() && q.s.IsEmpty()
}

// Size returns the number of elements in the queue
func (q *Queue) Size() int {
	return q.r.Size() + q.s.Size()
}

// Enqueue adds an element in the back of the queue
func (q *Queue) Enqueue(value interface{}) {
	for !q.s.IsEmpty() {
		v, _ := q.s.Pop()
		q.r.Push(v)
	}
	q.r.Push(value)
}

// Dequeue removes and returns the front element in the queue
func (q *Queue) Dequeue() (interface{}, error) {
	for !q.r.IsEmpty() {
		v, _ := q.r.Pop()
		q.s.Push(v)
	}
	return q.s.Pop()
}

// Peek returns the front element in the queue
func (q *Queue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue - empty queue")
	}

	for !q.r.IsEmpty() {
		v, _ := q.r.Pop()
		q.s.Push(v)
	}
	return q.s.Peek()
}

func (q *Queue) String() string {
	if q.IsEmpty() {
		return "[]"
	}
	if q.r.IsEmpty() {
		return q.s.String()
	} else if q.s.IsEmpty() {
		t := &stack.Stack{}
		t.Init()
		for !q.r.IsEmpty() {
			v, _ := q.r.Pop()
			t.Push(v)
		}

		res := "["
		for !t.IsEmpty() {
			v, _ := t.Pop()
			q.r.Push(v)
			res += fmt.Sprintf("%v, ", v)
		}
		return res[0:len(res)-2] + "]"
	} else {
		t := &stack.Stack{}
		t.Init()
		for !q.r.IsEmpty() {
			v, _ := q.r.Pop()
			t.Push(v)
		}

		res := "["
		for !t.IsEmpty() {
			v, _ := t.Pop()
			q.r.Push(v)
			res += fmt.Sprintf("%v, ", v)
		}

		sString := q.s.String()
		return res + sString[1:len(sString)-1]
	}
}
