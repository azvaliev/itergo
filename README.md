# itergo

[![Go Reference](https://pkg.go.dev/badge/github.com/azvaliev/itergo.svg)](https://pkg.go.dev/github.com/azvaliev/itergo)
[![Go Report Card](https://goreportcard.com/badge/github.com/azvaliev/itergo)](https://goreportcard.com/report/github.com/azvaliev/itergo)
![License](https://img.shields.io/github/license/azvaliev/itergo)

Collection of functional-style array iteration functions for Go.
Should be familiar to users of languages like JavaScript, Python, Ruby, C#.

## Installation

```bash
go get github.com/azvaliev/itergo
```

## Usage

See [GoDoc](https://pkg.go.dev/github.com/azvaliev/itergo) for detailed documentation on each function.

```go
package main

import "github.com/azvaliev/itergo"

func main() {
    arr := []int{1, 2, 3, 4, 5}
    itergo.FindLast(arr, func(x int, idx int) bool {
        return x % 2 == 0
    }) // 4
}
```