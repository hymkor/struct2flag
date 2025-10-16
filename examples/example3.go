//go:build run

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
