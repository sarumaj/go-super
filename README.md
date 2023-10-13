[![Go Report Card](https://goreportcard.com/badge/github.com/sarumaj/go-super)](https://goreportcard.com/report/github.com/sarumaj/go-super)
[![Maintainability](https://img.shields.io/codeclimate/maintainability-percentage/sarumaj/go-super.svg)](https://codeclimate.com/github/sarumaj/go-super/maintainability)

---

# go-super
Making Go more convenient

## github.com/sarumaj/go-super/errors

```Go

package main

import (
  "fmt"
  "os"

  supererrors "github.com/sarumaj/go-super/errors"
)

func main() {
  // Register callback for error
  supererrors.RegisterCallback(func(err error) {
    _, _ = fmt.Fprintln(os.Stderr, err)
  })

  // returns *os.File directly and calls callback only if error occures
  file := supererrors.ExceptFn(supererrors.W(os.Create("file.txt")))

  // try to close file and call callback on error which is not os.ErrClosed
  defer supererrors.Except(file.Close(), os.ErrClosed)
}  

```
