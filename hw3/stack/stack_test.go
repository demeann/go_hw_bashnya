package stack

import "testing"

func TestNewAndEmpty(t *testing.T) {
	s := New[int]()
	if !s.IsEmpty() {
		t.Fatalf("ожидался пустой стек")
	}
	if s.Size() != 0 {
		t.Fatalf("размер=%d, ожидался 0", s.Size())
	}
}

func TestClear(t *testing.T) {
	s := New[string]()
	s.Push("a")
	s.Push("b")
	s.Push("c")
	if s.IsEmpty() {
		t.Fatalf("ожидался непустой стек перед очисткой")
	}
	s.Clear()
	if !s.IsEmpty() || s.Size() != 0 {
		t.Fatalf("ожидался пустой стек после очистки")
	}
	_, ok := s.Pop()
	if ok {
		t.Fatalf("pop после Clear должен вернуть ok=false")
	}
}

func TestFloats(t *testing.T) {
	s := New[float64]()
	s.Push(1.5)
	s.Push(2.5)
	if s.Size() != 2 {
		t.Fatalf("размер=%d, ожидался 2", s.Size())
	}
	if v, ok := s.Pop(); !ok || v != 2.5 {
		t.Fatalf("ожидалось (2.5,true), получили (%v,%v)", v, ok)
	}
}
