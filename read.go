package gobutil

import (
	"bytes"
	"encoding/gob"
	"io"
)

// ReadReaderAny read struct from reader
func ReadReaderAny(r io.Reader, v any) error {
	dec := gob.NewDecoder(r)
	return dec.Decode(v)
}

// ReadReader read struct from reader
func ReadReader[T any](r io.Reader) (T, error) {
	var v T
	err := ReadReaderAny(r, &v)
	return v, err
}

// ReadBytesAny read struct from bytes
func ReadBytesAny(b []byte, v any) error {
	return ReadReaderAny(bytes.NewReader(b), v)
}

// ReadBytes read typed struct from bytes
func ReadBytes[T any](b []byte) (T, error) {
	return ReadReader[T](bytes.NewReader(b))
}
