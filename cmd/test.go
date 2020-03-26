package main

import (
	"fmt"
	"reflect"
	"strings"
)

var cache map[string]string

type MetaTest struct {
	Title   string `name:"title"`
	Twitter struct {
		Description string `name:"og:description"`
	}
}

func test() {
	cache = make(map[string]string)
	registerCache(&MetaTest{}, "")
	for name, f := range cache {
		fmt.Printf("(%s) %s\n", name, f)
	}
	test := &MetaTest{}
	test.updateField("title", "titleee")
	fmt.Printf("1 -> %+v\n", test)

	test.updateField("og:description", "test123")
	fmt.Printf("2 -> %+v\n", test)
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
			registerCache(fv, "")
		case reflect.Struct:
			registerCache(fv.Addr(), field.Name+"-")
		case reflect.String:
			tag, ok := field.Tag.Lookup("name")
			if !ok || tag == "-" {
				continue
			}

			for _, t := range strings.Split(tag, ",") {
				cache[t] = baseName + field.Name
				fmt.Printf("added %s from %q\n", t, baseName+field.Name)
			}
		}
	}
}

func (m *MetaTest) updateField(name, value string) {
	args, ok := cache[name]
	if !ok {
		fmt.Printf("%q does not exist in cache\n", name)
		return
	}

	updateField(reflect.ValueOf(m).Elem(), strings.Split(args, "-"), value)
}

func updateField(v reflect.Value, args []string, value string) {
	// fmt.Printf("called updateField with %T type and %s args\n", v, args)
	if len(args) > 1 {
		// fmt.Printf("    args > 1 for %s (%s) in type %T\n", args, value, v)
		updateField(v.FieldByName(args[0]), args[1:], value)
		return
	}
	// fmt.Println("should now be updating", args, value)

	v.FieldByName(args[0]).SetString(value)
}
