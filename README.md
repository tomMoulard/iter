# Iter

A go package that provides a ways to [iterate](https://go.dev/wiki/RangefuncExperiment) over elements.

## Installation

```bash
go get github.com/tommoulard/iter
```

## Usage

```go
package main

import (
    "fmt"

    "github.com/tommoulard/iter"
)


func main() {
    for a, b := range iter.Zip([]int{1, 2, 3}, []int{4, 5, 6}) {
        fmt.Println(a, b)
    }

    // Output:
    // 1 4
    // 2 5
    // 3 6
}
```

See the [GoDoc](https://pkg.go.dev/github.com/tommoulard/iter) for more information.
