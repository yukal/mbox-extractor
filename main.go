package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var lettersCount int
	var file *os.File
	var err error

	if file, err = os.Open("letters.mbox"); err != nil {
		log.Fatalf("error opening .mbox file: %v", err)
	}

	defer file.Close()

	if lettersCount, err = ExtractTo("eml", file); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("extracted %d letters\n", lettersCount)
}
