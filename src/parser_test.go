package main

import (
	"testing"
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

        file := "./callgraph.dot"

        // Parse the input file
        l,l_ := Parse(file)
        // Map names to ids
        caller,callee := Map(l_, "main", "buzz")

	for n := 0; n < b.N; n++ {
		Search(l, caller, callee)
	}
}
