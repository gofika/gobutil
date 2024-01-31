package gobutil

import (
	"bytes"
	"encoding/gob"
	"io"
)

// WriteWriter write struct to gob stream writer
func WriteWriter(w io.Writer, v any) error {
	enc := gob.NewEncoder(w)
	return enc.Encode(v)
}

// WriteBytes write struct to gob stream bytes
func WriteBytes(v any) ([]byte, error) {
	var buf bytes.Buffer
	err := WriteWriter(&buf, v)
	return buf.Bytes(), err
}
