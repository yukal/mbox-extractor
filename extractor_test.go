package main

import (
	"io"
	"io/fs"
	"testing"
	"yu/mboxextractor/data"
)

// Efficient File Reading in Go: Examples and Benchmark Comparisons
// https://medium.com/@smart_byte_labs/efficient-file-reading-in-go-examples-and-benchmark-comparisons-2335b097431a
// https://github.com/SmartByteLabs/filereadingexamples-go
// https://pkg.go.dev/embed

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

var (
	result int
	n      int
)

func BenchmarkExtract(b *testing.B) {
	var (
		fileName = "22M.mbox"
		file     fs.File
		err      error
	)

	if file, err = data.Files.Open(fileName); err != nil {
		b.Fatalf("error opening .mbox file: %v", err)
	}
	defer file.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if n, err = ExtractWithoutCursor(file); err != nil {
			b.Error(err)
		}

		// if n, err = ExtractWithCursor(file); err != nil {
		// 	b.Error(err)
		// }

		// Store the result to prevent the compiler optimizations
		result = n

		// Reset the file pointer
		if _, err := file.(io.Seeker).Seek(0, io.SeekStart); err != nil {
			b.Fatalf("Failed to reset %s file pointer: %v", fileName, err)
		}
	}
}
