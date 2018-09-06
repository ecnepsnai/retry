# retry
A go package to retry tasks that might eventually succeed

# Installation

```
go get github.com/ecnepsnai/retry
```

# Usage

```golang
package main

import "github.com/ecnepsnai/retry"

func main() {
    // Wait until it succeeds or fails
    retry.Try(func() error {
        return SomethingThatMightFail()
    }, 5)

    // Try in the background
    retry.TryAsync(func() error {
        return SomethingThatMightFail()
    }, 5, func(err error) {
        if err != nil {
            panic(err.Error())
        }
    })
}
```