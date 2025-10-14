package struct2flag

import (
	"flag"
	"reflect"
	"strings"
)

func Bind(fs *flag.FlagSet, cfg interface{}) {
	v := reflect.ValueOf(cfg).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		field := t.Field(i)

		if !field.IsExported() {
			continue
		}
		desc, ok := field.Tag.Lookup("flag")
		if !ok {
			continue
		}
		switch f.Kind() {
		case reflect.Struct:
			Bind(fs, f.Addr().Interface())
			continue
		case reflect.Pointer:
			if f.Type().Elem().Kind() == reflect.Struct && !f.IsNil() {
				Bind(fs, f.Interface())
				continue
			}
		}
		name, usage, ok := strings.Cut(desc, ",")
		if !ok {
			name = strings.ToLower(field.Name)
			usage = desc
		}
		switch f.Kind() {
		case reflect.Bool:
			fs.BoolVar(f.Addr().Interface().(*bool), name, f.Bool(), usage)
		case reflect.Int:
			fs.IntVar(f.Addr().Interface().(*int), name, int(f.Int()), usage)
		case reflect.Uint:
			fs.UintVar(f.Addr().Interface().(*uint), name, uint(f.Uint()), usage)
		case reflect.String:
			fs.StringVar(f.Addr().Interface().(*string), name, f.String(), usage)
		}
	}
}

func BindDefault(cfg interface{}) {
	Bind(flag.CommandLine, cfg)
}
