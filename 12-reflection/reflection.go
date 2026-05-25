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

	if val.Kind() == reflect.Slice {
		for i := range val.Len() {
			walk(val.Index(i).Interface(), fn)
		}
		return
	}

	for _, field := range val.Fields() {
		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}

func getValue(i any) reflect.Value {
	val := reflect.ValueOf(i)

	if val.Kind() == reflect.Pointer {
		return val.Elem()
	}

	return val
}
