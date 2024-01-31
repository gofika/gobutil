package gobutil

import (
	"os"

	"github.com/gofika/fileutil"
)

// ReadFileAny read struct from gob stream file
func ReadFileAny(filename string, v any) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return ReadReaderAny(f, v)
}

// ReadFile read struct from gob stream file
func ReadFile[T any](filename string) (T, error) {
	var v T
	err := ReadFileAny(filename, &v)
	return v, err
}

// WriteFile write struct to gob stream file
func WriteFile(filename string, v any) error {
	f, err := fileutil.OpenWrite(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return WriteWriter(f, v)
}
