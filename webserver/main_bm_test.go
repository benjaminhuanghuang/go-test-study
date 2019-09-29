package main

import (
	"testing"
)

func BenchmarkAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Pring1to20()
	}
}
