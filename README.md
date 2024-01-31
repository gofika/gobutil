[![codecov](https://codecov.io/gh/gofika/gobutil/branch/main/graph/badge.svg)](https://codecov.io/gh/gofika/gobutil)
[![Build Status](https://github.com/gofika/gobutil/workflows/build/badge.svg)](https://github.com/gofika/gobutil)
[![go.dev](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/gofika/gobutil)
[![Go Report Card](https://goreportcard.com/badge/github.com/gofika/gobutil)](https://goreportcard.com/report/github.com/gofika/gobutil)
[![Licenses](https://img.shields.io/github/license/gofika/gobutil)](LICENSE)

# gobutil

golang gob utils for common use


## Basic Usage

### Installation

To get the package, execute:

```bash
go get github.com/gofika/gobutil
```

### Example

```go
package main

import (
	"fmt"

	"github.com/gofika/gobutil"
)

func main() {
	type Foo struct {
		Name  string
		Value int
	}
	type Bar struct {
		Name  string
		Value int
	}
	foo := &Foo{"Jason", 100}
	// deep copy for different struct
	bar, err := gobutil.DeepCopy[Bar](foo)
	if err != nil {
		fmt.Printf("DeepCopy failed. err: %s\n", err.Error())
		return
	}
	fmt.Printf("bar.Name: %s\n", bar.Name)
	fmt.Printf("bar.Value: %d\n", bar.Value)
}
```