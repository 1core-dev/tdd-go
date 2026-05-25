package reflection

import "reflect"

func walk(x any, fn func(string)) {
	val := reflect.ValueOf(x)

	for _, field := range val.Fields() {
		if field.Kind() == reflect.String {
			fn(field.String())
		}
	}
}
