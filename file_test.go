package gobutil

import (
	"path"
	"testing"

	"github.com/gofika/fileutil"
	"github.com/stretchr/testify/assert"
)

func TestFile(t *testing.T) {
	tempName := path.Join(tempDir, "foo.gob")
	foo := &Foo{"Jason", 100}

	err := WriteFile(tempName, foo)
	if !assert.Nil(t, err) {
		return
	}
	defer fileutil.Delete(tempName)

	bar, err := ReadFile[Foo](tempName)
	if !assert.Nil(t, err) {
		return
	}
	assert.Equal(t, bar.Name, "Jason")
}
