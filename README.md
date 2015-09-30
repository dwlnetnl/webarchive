# webarchive ![GoDoc](https://godoc.org/github.com/dwlnetnl/webarchive?status.svg)
Package webarchive reads Safari .webarchive files.

## Example
```go
package main

import "github.com/dwlnetnl/webarchive"

func main() {
  r, err := os.Open("/path/to/saved.webarchive")
  if err != nil {
    log.Fatal(err)
    return
  }

  if a, err := New(r); err == nil {
    fmt.Println(a.Content.URL) // URL of web page.

    // Copy resource to stdout.
    io.Copy(os.Stdout, a.Content.Reader())
  }
}

```
