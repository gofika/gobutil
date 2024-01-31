package gobutil

import (
	"bytes"
)

// DeepCopy clone object. can clone without same type
//
// Example:
//
//	type Foo struct {
//	    Name string
//	    Value int
//	}
//
//	type Bar struct {
//	    Name string
//	    Value int
//	}
//
//	foo := &Foo { "Jason", 100}
//	bar, err := DeepCopy[Bar](foo)
//	fmt.Printf("%+v\n", bar)
func DeepCopy[T any](src any) (T, error) {
	var v T
	err := DeepCopyAny(&v, src)
	return v, err
}

// DeepCopyAny clone object to dst. can clone without same type
func DeepCopyAny(dst, src any) error {
	var buf bytes.Buffer
	if err := WriteWriter(&buf, src); err != nil {
		return err
	}
	return ReadReaderAny(&buf, dst)
}
