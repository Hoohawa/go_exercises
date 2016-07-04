// Run: $ go test -bench=PopCount -benchmem
package main

import (
	"fmt"
	"math/rand"
	"testing"
)

var maxUint64 uint64 = (1<<64 - 1)

func TestAllPopCountFunctionReturnsTheSame(t *testing.T) {
	for i := 0; i < 1000; i++ {
		n := uint64(rand.Uint32())
		r1 := PopCount1(n)
		r2 := PopCount2(n)
		r3 := PopCount3(n)
		if r1 != r1 || r1 != r3 {
			t.Error(fmt.Sprintf("n=%v r1=%v r2=%v r3=%v", n, r1, r2, r3))
		}
	}
}

func BenchmarkPopCount1(t *testing.B) {
	for i := 0; i < t.N; i++ {
		for n := 0; n < 5000; n++ {
			PopCount1(uint64(rand.Uint32()))
		}
	}
}

func BenchmarkPopCount2(t *testing.B) {
	for i := 0; i < t.N; i++ {
		for n := 0; n < 5000; n++ {
			PopCount2(uint64(rand.Uint32()))
		}
	}
}

func BenchmarkPopCount3(t *testing.B) {
	for i := 0; i < t.N; i++ {
		for n := 0; n < 5000; n++ {
			PopCount3(uint64(rand.Uint32()))
		}
	}
}
