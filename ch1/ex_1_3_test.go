package main

import (
	"testing"
)

func BenchmarkApproach1(t *testing.B) {
	args := generateTestArgList()

	for i := 0; i < t.N; i++ {
		Approach1(args)
	}
}

func BenchmarkApproach2(t *testing.B) {
	args := generateTestArgList()

	for i := 0; i < t.N; i++ {
		Approach2(args)
	}
}

func generateTestArgList() []string {
	args := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		args = append(args, string(i))
	}
	return args
}
