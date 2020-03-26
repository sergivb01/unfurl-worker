package meta

import (
	"reflect"
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
			break
		case reflect.Struct:
			registerCache(fv.Addr(), field.Name+"-")
		case reflect.String:
			tag, ok := field.Tag.Lookup("meta")
			if !ok || tag == "-" {
				continue
			}

			for _, t := range strings.Split(tag, ",") {
				internalReflectionCache[t] = baseName + field.Name
				// fmt.Printf("added %s from %q\n", t, baseName+field.Name)
			}
		}
	}
}

func (m *PageInfo) updateField(name, value string) {
	args, ok := internalReflectionCache[name]
	if !ok {
		return
	}
	updateField(reflect.ValueOf(m).Elem(), strings.Split(args, "-"), value)
}

func updateField(v reflect.Value, args []string, value string) {
	if len(args) > 1 {
		updateField(v.FieldByName(args[0]), args[1:], value)
		return
	}
	v.FieldByName(args[0]).SetString(value)
}
