package main

import (
	"encoding/xml"
	"os"

	"github.com/raitucarp/gown"
)

func readLexicalResource(filePath string) (resource gown.LexicalResource, err error) {
	s, err := os.ReadFile(filePath)
	if err != nil {
		return
	}

	err = xml.Unmarshal(s, &resource)
	if err != nil {
		return
	}

	return
}
