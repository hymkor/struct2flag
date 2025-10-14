//go:build run

package main

import (
	"flag"
	"fmt"

	"github.com/hymkor/struct2flag"
)

type SubConfig1 struct {
	Debug bool `flag:"d,Enable debug mode"`
}

type SubConfig2 struct {
	Verbose bool `flag:"v,Enable verbose mode"`
}

type Config struct {
	Name string `flag:"name,Set your name"`
	Sub1 SubConfig1
	Sub2 SubConfig2 `flag:""`
}

func (c *Config) Run() {
	fmt.Printf("Name=%#v\n", c.Name)
	fmt.Printf("Sub1.Debug=%#v\n", c.Sub1.Debug)
	fmt.Printf("Sub2.Verbose=%#v\n", c.Sub2.Verbose)
}

func main() {
	var cfg Config
	struct2flag.BindDefault(&cfg)
	flag.Parse()

	cfg.Run()
}
