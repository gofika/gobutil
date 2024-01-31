package gobutil

import (
	"os"
)

var (
	tempDir, _ = os.MkdirTemp("", "util")
)

type Foo struct {
	Name  string
	Value int
}

type Bar struct {
	Name  string
	Value int
}
