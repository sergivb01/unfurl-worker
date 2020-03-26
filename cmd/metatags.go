package main

import (
	"reflect"
	"strings"
)

var tagCache map[string]int

/**
 * TODO:
 * Add https://ogp.me/
 * Add https://developer.twitter.com/en/docs/tweets/optimize-with-cards/overview/markup
 */

type Metatags struct {
	Title       string `name:"title, twitter:title"`
	Description string `name:"description, og:description, twitter:description"`
	Author      string `name:"author"`
	Keywords    string `name:"keywords"`
	ThemeColor  string `name:"theme-color, msapplication-TileColor"`

	OGAudio  string `name:"og:audio"`
	OGLocale string `name:"og:locale"`
	OGVideo  string `name:"og:video"`
	OGType   string `name:"og:type"`

	TwitterCard     string `name:"twitter:card"`
	TwitterImage    string `name:"twitter:image"`
	TwitterImageAlt string `name:"twitter:image:alt"`
	TwitterPlayer   string `name:"twitter:player"`
}

func init() {
	v := reflect.ValueOf(&Metatags{}).Elem()
	tagCache = make(map[string]int, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		typeField := v.Type().Field(i)
		tag, ok := typeField.Tag.Lookup("name")
		if !ok || tag == "-" {
			continue
		}

		for _, str := range strings.Split(tag, ",") {
			tagCache[strings.TrimSpace(str)] = i
		}
	}
}

func (m *Metatags) updateField(name, value string) {
	v := reflect.ValueOf(m).Elem()

	tagIdx, ok := tagCache[name]
	if !ok || v.Field(tagIdx).String() != "" {
		// fmt.Println("ignoring " + name)
		return
	}
	v.Field(tagIdx).SetString(value)
}
