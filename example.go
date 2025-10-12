//go:build run

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
