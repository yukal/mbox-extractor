package main

import (
	"os"
	"testing"
)

// https://blog.logrocket.com/benchmarking-golang-improve-function-performance/
// https://www.linkedin.com/learning/go-performance-tuning-and-benchmarking/using-benchstat

// Benchmarking options:
//
// go test -bench=. -benchmem
// go test -bench=. -benchtime=500000x -benchmem
// go test -bench=. -benchtime=5s -benchmem
// go test -bench=. -count 5 -benchmem
//
// go test -bench BenchmarkExtract -count 10 |& tee /tmp/benchmark1
// go test -benchmem -bench BenchmarkExtract -count 10 |& tee /tmp/benchmark1

func BenchmarkExtract(b *testing.B) {
	const mboxFile = "letters.mbox"

	var (
		file *os.File
		err  error
	)

	if file, err = os.Open(mboxFile); err != nil {
		b.Fatalf("error opening .mbox file: %v", err)
	}
	defer file.Close()

	// for i := 0; i < b.N; i++ {
	// 	if err := ExtractWithoutCursor(file); err != nil {
	// 		b.Error(err)
	// 	}
	// }

	for i := 0; i < b.N; i++ {
		if err := ExtractWithCursor(file); err != nil {
			b.Fatal(err)
		}
	}
}
