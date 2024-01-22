package yandexlogistic

import (
	"encoding/json"
	"net/url"
	"reflect"
	"strconv"
)

func addValues(vals url.Values, i any) url.Values {
	if i == nil {
		return vals
	}
	if vals == nil {
		vals = make(url.Values)
	}
	return scan(i, vals)
}

func scan(i any, v url.Values) url.Values {
	e := reflect.ValueOf(i)

	if e.Kind() == reflect.Pointer {
		e = e.Elem()
	}

	if e.Kind() == reflect.Invalid {
		return v
	}

	for i := 0; i < e.NumField(); i++ {
		fTag := e.Type().Field(i).Tag

		if name := fTag.Get("query"); name != "" && !e.Field(i).IsZero() {
			v.Set(name, toString(e.Field(i)))
		}
	}

	return v
}

func toString(v reflect.Value) string {
	switch v.Kind() {
	case reflect.String:
		return v.String()

	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)

	case reflect.Int, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)

	case reflect.Bool:
		return strconv.FormatBool(v.Bool())

	case reflect.Struct, reflect.Interface, reflect.Slice, reflect.Array:
		b, _ := json.Marshal(v.Interface())
		return string(b)

	default:
		return ""
	}
}
