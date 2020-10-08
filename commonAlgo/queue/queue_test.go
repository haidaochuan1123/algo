package queue

import (
	"testing"
)

var s ItemQueue

func TestEnQuene(t *testing.T) {
	s := NewItemQueue()
	s.EnQuene(1)
	s.EnQuene(2)
	s.EnQuene(3)

	if size := s.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}
}

func TestDeQuene(t *testing.T) {
	s.EnQuene(1)
	s.EnQuene(1)
	s.EnQuene(1)
	s.DeQuene()
	if size := len(s.items); size != 2 {
		t.Errorf("wrong count, expected 2 and got %d", size)
	}

	s.DeQuene()
	s.DeQuene()
	if size := len(s.items); size != 0 {
		t.Errorf("wrong count, expected 0 and got %d", size)
	}

	if !s.IsEmpty() {
		t.Errorf("IsEmpty should return true")
	}
}
