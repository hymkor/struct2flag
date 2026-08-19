struct2flag
===========

<!-- badges.cmd |-->
[![Go Test](https://github.com/hymkor/struct2flag/actions/workflows/go.yml/badge.svg)](https://github.com/hymkor/struct2flag/actions/workflows/go.yml)
[![License](https://img.shields.io/badge/License-MIT-red)](https://github.com/hymkor/struct2flag/blob/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/hymkor/struct2flag.svg)](https://pkg.go.dev/github.com/hymkor/struct2flag)
<!-- -->

`struct2flag` automatically registers struct fields as flags for your Go command-line programs.

Minimal example
---------------

`example.go`

```examples/example.go
package main

import (
    "flag"
    "fmt"
    "time"

    "github.com/hymkor/struct2flag"
)

type Env struct {
    B bool          `flag:"b,This is a boolean flag"`
    N int           `flag:"n,This is an integer flag"`
    S string        `flag:"s,this is a string flag"`
    F float64       `flag:"f,this is a float flag"`
    D time.Duration `flag:"d,this is a time.duration flag"`
}

func (e Env) Run() {
    fmt.Printf("B=%#v\n", e.B)
    fmt.Printf("N=%#v\n", e.N)
    fmt.Printf("S=%#v\n", e.S)
    fmt.Printf("F=%#v\n", e.F)
    fmt.Printf("D=%v\n", e.D)
}

func main() {
    var env Env
    struct2flag.BindDefault(&env)
    flag.Parse()
    env.Run()
}
```

`go run examples/example.go -h`

```go run examples/example.go -h |
Usage of R:\go-build3962870177\b001\exe\example.exe:
  -b	This is a boolean flag
  -d duration
    	this is a time.duration flag
  -f float
    	this is a float flag
  -n int
    	This is an integer flag
  -s string
    	this is a string flag
```

`go run examples/example.go -b -n 1 -s foo -f 2.0 -d 1m30s`

```go run examples/example.go -b -n 1 -s foo -f 2.0 -d 1m30s |
B=true
N=1
S="foo"
F=2
D=1m30s
```

Reading default values from JSON and overriding them with command-line flags
----------------------------------------------------------------------------

`examples/example3.go`

```examples/example3.go
package main

import (
    "flag"
    "fmt"
    "os"

    "encoding/json"

    "github.com/hymkor/struct2flag"
)

type Env struct {
    B bool   `flag:"b,This is a boolean flag"  json:"b"`
    N int    `flag:"n,This is an integer flag" json:"n"`
    S string `flag:"s,this is a string flag"   json:"s"`
}

func (e Env) Run() {
    fmt.Printf("B=%#v\n", e.B)
    fmt.Printf("N=%#v\n", e.N)
    fmt.Printf("S=%#v\n", e.S)
}

func main() {
    var env Env

    if data, err := os.ReadFile("example3.json"); err == nil {
        err = json.Unmarshal(data, &env)
    }
    struct2flag.BindDefault(&env)
    flag.Parse()
    env.Run()
}
```

`examples/example3.json`

```examples/example3.json
{
    "b": true,
    "n": 100,
    "s": "hogehoge"
}
```

`go run -C examples example3.go -s changed`

```go run -C examples example3.go -s changed |
B=true
N=100
S="changed"
```
