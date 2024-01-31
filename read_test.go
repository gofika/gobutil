package gobutil

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadBytes(t *testing.T) {
	foo := &Foo{"Jason", 100}
	b, err := WriteBytes(foo)
	if !assert.Nil(t, err) {
		return
	}

	bar, err := ReadBytes[Foo](b)
	if !assert.Nil(t, err) {
		return
	}
	assert.Equal(t, bar.Name, "Jason")
}

func TestReadReader(t *testing.T) {
	foo := &Foo{"Jason", 100}
	b, err := WriteBytes(foo)
	if !assert.Nil(t, err) {
		return
	}

	bar, err := ReadReader[Foo](bytes.NewReader(b))
	if !assert.Nil(t, err) {
		return
	}
	assert.Equal(t, bar.Name, "Jason")
}
