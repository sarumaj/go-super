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
  file := supererrors.ExceptFn2(supererrors.W2(os.Create("file.txt)))

  // try to close file and call callback on error which is not os.ErrClosed
  defer supererrors.Except(file.Close(), os.ErrClosed)
}  

```
