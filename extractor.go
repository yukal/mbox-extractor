package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
)

const Kb = 1024

type SequenceMap map[string]int

func (s SequenceMap) ToUnique(token, postfix string) string {
	if n, exist := s[token]; !exist {

		s[token] = 1

	} else {

		s[token] = n + 1
		token += "_" + strconv.Itoa(s[token])

	}

	return token + postfix
}

func ExtractTo(destinationDir string, file io.ReadCloser) (int, error) {
	searchingPhrase := []byte("\r\n\r\nFrom ")
	phraseLen := len(searchingPhrase)
	sequence := make(SequenceMap)

	cursor := 0
	lettersCount := 0

	letters := make([]byte, 0)
	buf := make([]byte, 4*Kb)

	for {
		n, err := file.Read(buf)

		if err != nil {
			if err == io.EOF {
				break
			}

			return lettersCount, fmt.Errorf("error reading file: %v", err)
		}

		letters = append(letters, buf[:n]...)

		for {
			if posEnding := bytes.Index(letters[cursor:], searchingPhrase); posEnding > -1 {
				posEnding += cursor + 4 // last [\r\n]+

				filename := sequence.ToUnique(getLetterId(letters[:posEnding]), ".eml")
				filepath := path.Join(destinationDir, filename)

				if err := os.WriteFile(filepath, letters[:posEnding], 0644); err != nil {
					// log.Printf("error saving file: %v", err)
					return lettersCount, fmt.Errorf("error saving file: %v", err)
				}

				lettersCount++
				letters = letters[posEnding:]

				cursor = len(letters)

				if startFrom := cursor - phraseLen + 1; startFrom > 0 {
					cursor = startFrom
				}

			} else {

				cursor = len(letters)

				if startFrom := cursor - phraseLen + 1; startFrom > 0 {
					cursor = startFrom
				}

				break

			}
		}
	}

	if len(letters) > 0 {
		filename := sequence.ToUnique(getLetterId(letters), ".eml")
		filepath := path.Join(destinationDir, filename)

		if err := os.WriteFile(filepath, letters, 0644); err != nil {
			// log.Printf("error saving file: %v", err)
			return lettersCount, fmt.Errorf("error saving file: %v", err)
		}

		lettersCount++
	}

	return lettersCount, nil
}

func getLetterId(message []byte) string {
	pos := bytes.IndexByte(message[5:], '@')

	// From 1826299905545011489@xxx Tue Mar 11 12:15:13 +0000 2025
	return string(message[5:pos])
}
