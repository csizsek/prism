package main

import "testing"

func TestAdd(t *testing.T) {
	if add(0, 0) != 0 {
		t.Fail()
	}
}
