//go:build run

package main

import (
	"flag"
	"fmt"

	"github.com/hymkor/struct2flag"
)

type SubConfig struct {
	Debug bool `flag:"debug,Enable debug mode"`
}

type Config struct {
	Name string    `flag:"name,Set your name"`
	Sub  SubConfig // struct
}

func (c *Config) Run() {
	fmt.Printf("Name=%#v\n", c.Name)
	fmt.Printf("Sub.Debug=%#v\n", c.Sub.Debug)
}

func main() {
	var cfg Config
	struct2flag.BindDefault(&cfg)
	flag.Parse()

	cfg.Run()
}
