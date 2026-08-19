package core

import (
	"reflect"
	"strings"
)

type FlagSet interface {
	BoolVar(p *bool, name string, value bool, usage string)
	IntVar(p *int, name string, value int, usage string)
	UintVar(p *uint, name string, value uint, usage string)
	StringVar(p *string, name string, value string, usage string)
	Float64Var(p *float64, name string, value float64, usage string)
}

func BindTag(tag string, fs FlagSet, cfg interface{}) {
	v := reflect.ValueOf(cfg).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		field := t.Field(i)

		if !field.IsExported() {
			continue
		}
		desc, ok := field.Tag.Lookup(tag)
		if !ok {
			continue
		}
		switch f.Kind() {
		case reflect.Struct:
			BindTag(tag, fs, f.Addr().Interface())
			continue
		case reflect.Pointer:
			if f.Type().Elem().Kind() == reflect.Struct && !f.IsNil() {
				BindTag(tag, fs, f.Interface())
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
		case reflect.Float64:
			fs.Float64Var(f.Addr().Interface().(*float64), name, f.Float(), usage)
		}
	}
}
