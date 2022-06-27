# doener

[![Go Report Card](https://goreportcard.com/badge/github.com/j0hax/doener)](https://goreportcard.com/report/github.com/j0hax/doener)

Have you ever wanted to represent Germany's favorite street food as a data structure? Here's your chance.

## Usage

```go
package main

import (
    "fmt"
    "github.com/j0hax/doener"
)

func main() {
    doener := doener.Random()
    fmt.Println(doener)
}
```

```console
Lamm Döner mit scharf, Tsatsiki und Schafskäse, Gurken, Zwiebeln
```

## Documentation

Please see the official documentation on [go.dev](https://pkg.go.dev/github.com/j0hax/doener).

## Purpose

While it's easy to *just write* Go, I wanted to take the opportunity to also learn how to properly document and maintain a Go package
