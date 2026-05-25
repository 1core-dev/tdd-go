package reflection

import "reflect"

func walk(x any, fn func(string)) {
	val := reflect.ValueOf(x)

	for _, field := range val.Fields() {
		fn(field.String())
	}
}
