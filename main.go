package struct2flag

import (
	"flag"

	"github.com/hymkor/struct2flag/core"
)

func Bind(fs core.FlagSet, cfg interface{}) {
	core.BindTag("flag", fs, cfg)
}

func BindDefault(cfg interface{}) {
	Bind(flag.CommandLine, cfg)
}
