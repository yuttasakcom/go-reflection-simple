package report

import (
	"fmt"
	"reflect"
	"strings"
)

func Text(v interface{}) string {
	rv := reflect.ValueOf(v)
	var sb strings.Builder
	text(&sb, &rv)
	return sb.String()
}

func text(sb *strings.Builder, rv *reflect.Value) {
	switch rv.Kind() {
	case reflect.Struct:
		structName := rv.Type().Name()
		if len(structName) == 0 {
			structName = "Anonymous"
		}
		sb.WriteString(structName + " {")
		for i := 0; i < rv.NumField(); i++ {
			val := rv.Field(i)
			if i > 0 {
				sb.WriteString(", ")
			}
			fieldName := rv.Type().Field(i).Name
			tagName, ok := rv.Type().Field(i).Tag.Lookup("report")
			if ok {
				tags := strings.Split(tagName, ",")
				if len(tags) == 2 {
					fieldName = tags[0]
					if tags[1] == "uppercase" {
						val = reflect.ValueOf(strings.ToUpper(val.String()))
					}
				}
			}
			sb.WriteString(fieldName + ": ")
			text(sb, &val)
		}
		sb.WriteString("}")
	case reflect.String:
		sb.WriteString(fmt.Sprintf("%q", rv.String()))
	case reflect.Int:
		sb.WriteString(fmt.Sprintf("%d", rv.Int()))
	}
}
