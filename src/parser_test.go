package main

import (
	"testing"
	"container/list"
)

func TestParser(t *testing.T) {

}

func BenchmarkParser(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
	}
}

func BenchmarkSearch(b *testing.B) {
	// run the function b.N times

	caller := "main"
	callee := "buzz"
	l := list.New()
	for n := 0; n < b.N; n++ {
		Search(l, caller, callee)
	}
}
