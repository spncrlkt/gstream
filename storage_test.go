package main

import (
	"testing"
)

func TestStorage(t *testing.T) {
	s := NewMemoryStore()
	for i := 0; i < 100; i++ {
		latestOffset, err := s.Push([]byte("foobarbaz"))
		if err != nil {
			t.Error(err)
		}
		_, err2 := s.Fetch(latestOffset)
		if err2 != nil {
			t.Error(err2)
		}
	}
}
