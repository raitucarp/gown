package main

import (
	"bufio"
	"compress/gzip"
	"encoding/gob"
	"io"
	"os"

	"github.com/raitucarp/gown"
)

func createOewnGobFile() (gobFile *os.File, err error) {
	gobFile, err = os.Create("./data/oewn.gob")
	if err != nil {
		return
	}
	return
}

func newGzipAndBufferedWriter(w io.Writer) (*gzip.Writer, *bufio.Writer, error) {
	gzipWriter, err := gzip.NewWriterLevel(w, gzip.BestCompression)
	if err != nil {
		return nil, nil, err
	}
	bufferedWriter := bufio.NewWriter(gzipWriter)
	return gzipWriter, bufferedWriter, nil
}

func encodeOewn(w io.Writer, resource gown.LexicalResource) (err error) {
	encoder := gob.NewEncoder(w)
	err = encoder.Encode(resource)
	if err != nil {
		return
	}
	return
}
