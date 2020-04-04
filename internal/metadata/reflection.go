package metadata

import (
	"reflect"
	"strconv"
	"strings"
)

var internalReflectionCache = make(map[string]string)

func init() {
	registerCache(&PageInfo{}, "")
}

func registerCache(data interface{}, baseName string) {
	var (
		rv reflect.Value
		ok bool
	)

	if rv, ok = data.(reflect.Value); !ok {
		rv = reflect.ValueOf(data)
	}

	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		panic(ErrorType)
	}

	rt := rv.Type()
	for i := 0; i < rv.Elem().NumField(); i++ {
		fv := rv.Elem().Field(i)
		field := rt.Elem().Field(i)

		switch fv.Type().Kind() {
		case reflect.Ptr:
			if fv.IsNil() {
				fv.Set(reflect.New(fv.Type().Elem()))
			}
			registerCache(fv, field.Name+"-")
		case reflect.Slice:
			// TODO: implement slices for Images, Videos and Audios
		case reflect.Struct:
			registerCache(fv.Addr(), field.Name+"-")
		case reflect.String:
			tag, ok := field.Tag.Lookup("metadata")
			if !ok || tag == "-" {
				continue
			}

			for _, t := range strings.Split(tag, ",") {
				internalReflectionCache[t] = baseName + field.Name
			}
		}
	}
}

func (m *PageInfo) updateField(name, value string) {
	if args, ok := internalReflectionCache[name]; ok {
		updateField(reflect.ValueOf(m).Elem(), strings.Split(args, "-"), value)
	}
}

func updateField(v reflect.Value, args []string, value string) {
	if len(args) > 1 {
		updateField(v.FieldByName(args[0]), args[1:], value)
		return
	}

	field := v.FieldByName(args[0])
	switch field.Type().Kind() {
	case reflect.String:
		// don't update meta-tag if already found a valid tag
		if field.String() != "" {
			field.SetString(value)
		}
	case reflect.Int:
		if n, err := strconv.Atoi(value); err == nil {
			field.SetInt(int64(n))
		}
	default:
		// TODO: field not supported, could log?
	}
	// v.FieldByName(args[0]).SetString(value)
}
