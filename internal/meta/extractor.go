package meta

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/sergivb01/unfurl-worker/internal/utils"
)

var ErrorType = errors.New("should not be non-ptr or nil")

func ExtractInfoFromReader(reader io.ReadCloser) (*PageInfo, error) {
	defer utils.BenchmarkFunction(utils.StartBench(), "ExtractInfoFromReader(io.ReadCloser)")

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("error creating document from reader: %w", err)
	}

	info := &PageInfo{}
	if err := getPageData(doc, info); err != nil {
		return nil, fmt.Errorf("error getting info: %w", err)
	}

	return info, nil
}

func getPageData(doc *goquery.Document, data interface{}) error {
	var (
		rv reflect.Value
		ok bool
	)

	if rv, ok = data.(reflect.Value); !ok {
		rv = reflect.ValueOf(data)
	}

	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return ErrorType
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
			if err := getPageData(doc, fv); err != nil {
				return err
			}
		case reflect.Struct:
			if err := getPageData(doc, fv.Addr()); err != nil {
				return err
			}
		case reflect.Slice:
			if fv.IsNil() {
				fv.Set(reflect.MakeSlice(fv.Type(), 0, 0))
			}

			switch field.Type.Elem().Kind() {
			case reflect.Struct:
				last := reflect.New(field.Type.Elem())
				for {
					data := reflect.New(field.Type.Elem())
					if err := getPageData(doc, data.Interface()); err != nil {
						return err
					}

					// Ugly solution (I can't remove nodes. Why?)
					if !reflect.DeepEqual(last.Elem().Interface(), data.Elem().Interface()) {
						fv.Set(reflect.Append(fv, data.Elem()))
						last.Elem().Set(data.Elem())

					} else {
						break
					}
				}
			case reflect.Ptr:
				last := reflect.New(field.Type.Elem().Elem())
				for {
					data := reflect.New(field.Type.Elem().Elem())
					if err := getPageData(doc, data.Interface()); err != nil {
						return err
					}

					// Ugly solution (I can't remove nodes. Why?)
					if !reflect.DeepEqual(last.Elem().Interface(), data.Elem().Interface()) {
						fv.Set(reflect.Append(fv, data))
						last.Elem().Set(data.Elem())

					} else {
						break
					}
				}
			default:
				if tag, ok := field.Tag.Lookup("meta"); ok {

					for _, t := range strings.Split(tag, ",") {
						var contents []reflect.Value

						processMeta := func(idx int, sel *goquery.Selection) {
							if c, existed := sel.Attr("content"); existed {
								if field.Type.Elem().Kind() == reflect.String {
									contents = append(contents, reflect.ValueOf(c))
								} else {
									i, e := strconv.Atoi(c)

									if e == nil {
										contents = append(contents, reflect.ValueOf(i))
									}
								}

								fv.Set(reflect.Append(fv, contents...))
							}
						}

						doc.Find("meta[property=\"" + t + "\"]").Each(processMeta)
						doc.Find("meta[name=\"" + t + "\"]").Each(processMeta)

						fv = reflect.Append(fv, contents...)
					}
				}
			}
		default:
			if tag, ok := field.Tag.Lookup("meta"); ok {
				var (
					content string
					existed bool
					sel     *goquery.Selection
				)

				for _, t := range strings.Split(tag, ",") {
					if sel = doc.Find("meta[property=\"" + t + "\"]").First(); sel.Size() > 0 {
						content, existed = sel.Attr("content")
					}

					if !existed {
						if sel = doc.Find("meta[name=\"" + t + "\"]").First(); sel.Size() > 0 {
							content, existed = sel.Attr("content")
						}
					}

					if existed {
						if fv.Type().Kind() == reflect.String {
							fv.Set(reflect.ValueOf(content))
						} else if fv.Type().Kind() == reflect.Int {
							if i, e := strconv.Atoi(content); e == nil {
								fv.Set(reflect.ValueOf(i))
							}
						}
						break
					}
				}
			}
		}
	}
	return nil
}
