package gobutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeepCopy(t *testing.T) {
	foo := &Foo{"Jason", 100}
	bar, err := DeepCopy[Bar](foo)
	if !assert.Nil(t, err) {
		return
	}
	assert.Equal(t, bar.Name, foo.Name)
	assert.Equal(t, bar.Value, foo.Value)
}

func TestDeepCopyAny(t *testing.T) {
	foo := &Foo{"Jason", 100}
	var bar Bar
	err := DeepCopyAny(&bar, foo)
	if !assert.Nil(t, err) {
		return
	}
	assert.Equal(t, bar.Name, foo.Name)
	assert.Equal(t, bar.Value, foo.Value)
}
