package main

import (
	"log"
)

func main() {
	destinationDir := "eml"
	mboxFile := "letters.mbox"

	if err := ExtractTo(destinationDir, mboxFile); err != nil {
		log.Fatal(err)
	}
}
