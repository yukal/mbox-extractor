package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("letters.mbox")

	if err != nil {
		log.Fatalf("error opening .mbox file: %v", err)
	}

	defer file.Close()

	if err := ExtractTo("eml", file); err != nil {
		log.Fatal(err)
	}
}
