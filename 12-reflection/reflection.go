package reflection

import "reflect"

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func walk(x any, fn func(string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.Struct:
		for _, field := range val.Fields() {
			walk(field.Interface(), fn)
		}
	case reflect.Slice:
		for i := range val.Len() {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
	}
}

func getValue(i any) reflect.Value {
	val := reflect.ValueOf(i)

	if val.Kind() == reflect.Pointer {
		return val.Elem()
	}

	return val
}
