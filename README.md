struct2flag
===========

`struct2flag` automatically registers struct fields as flags for your Go command-line tools.

`example.go`

```example.go
package main

import (
    "flag"
    "fmt"

    "github.com/hymkor/struct2flag"
)

type Env struct {
    B bool   `flag:"b,This is a boolean flag"`
    N int    `flag:"n,This is an integer flag"`
    S string `flag:"s,this is a string flag"`
}

func (e Env) Run() {
    fmt.Printf("B=%#v\n", e.B)
    fmt.Printf("N=%#v\n", e.N)
    fmt.Printf("S=%#v\n", e.S)
}

func main() {
    var env Env
    struct2flag.BindDefault(&env)
    flag.Parse()
    env.Run()
}
```

`go run example.go -h`

```go run example.go -h |
Usage of R:\a4\a44a86327d99b87e3f152a25b4797c07342ef4e9a8505c7770013709f4b45dc6-d\example.exe:
  -b	This is a boolean flag
  -n int
    	This is an integer flag
  -s string
    	this is a string flag
```

`go run example.go -b -n 1 -s foo`

```go run example.go -b -n 1 -s foo |
B=true
N=1
S="foo"
```
