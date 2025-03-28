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

func ExtractTo(destinationDir string, file io.ReadCloser) error {
	searchingPhrase := []byte("\r\n\r\nFrom ")
	sequence := make(SequenceMap)

	cursor := 0
	letters := make([]byte, 0)
	buf := make([]byte, 4*Kb)

	for {
		n, err := file.Read(buf)

		if err != nil {
			if err == io.EOF {
				break
			}

			return fmt.Errorf("error reading file: %v", err)
		}

		buf = buf[:n]
		letters = append(letters, buf...)

		for {
			if posEnding := bytes.Index(letters[cursor:], searchingPhrase); posEnding > -1 {
				posEnding += cursor + 4 // last [\r\n]+

				filename := sequence.ToUnique(getLetterId(letters[:posEnding]), ".eml")
				filepath := path.Join(destinationDir, filename)

				if err := os.WriteFile(filepath, letters[:posEnding], 0644); err != nil {
					// log.Printf("error saving file: %v", err)
					return fmt.Errorf("error saving file: %v", err)
				}

				letters = letters[posEnding:]
				cursor = len(letters)

			} else {

				cursor = len(letters)
				break

			}
		}
	}

	if len(letters) > 0 {
		filename := sequence.ToUnique(getLetterId(letters), ".eml")
		filepath := path.Join(destinationDir, filename)

		if err := os.WriteFile(filepath, letters, 0644); err != nil {
			// log.Printf("error saving file: %v", err)
			return fmt.Errorf("error saving file: %v", err)
		}
	}

	return nil
}

func getLetterId(message []byte) string {
	pos := bytes.IndexByte(message[5:], '@')

	// From 1826299905545011489@xxx Tue Mar 11 12:15:13 +0000 2025
	return string(message[5:pos])
}
