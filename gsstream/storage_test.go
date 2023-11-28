package main

import (
	"fmt"
	"testing"
)

func TestStorage(t *testing.T) {
	s := NewMemoryStore()
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("footbar_baz[%d]", i)
		lastOffset, err := s.Push([]byte(key))
		if err != nil {
			t.Error(err)
		}
		data, err := s.Fetch(lastOffset)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(string(data))
	}
}
