package main

import "testing"


func TestIter(t *testing.T) {
	got := Repeat("a")
	exp := "aaaaa"

	if got != exp {
		t.Errorf("expected '%q' but got '%q'", exp, got)
	}
}

func  BenchmarkRepeat(b *testing.B){
	for i := 0; i < b.N; i++{
		Repeat("a")
	}
}
