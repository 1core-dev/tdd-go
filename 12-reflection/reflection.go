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
	var valuesCount int
	var getField func(int) reflect.Value

	val := getValue(x)

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		valuesCount = val.NumField()
		getField = val.Field
	case reflect.Slice, reflect.Array:
		valuesCount = val.Len()
		getField = val.Index
	}

	for i := range valuesCount {
		walk(getField(i).Interface(), fn)
	}
}

func getValue(i any) reflect.Value {
	val := reflect.ValueOf(i)

	if val.Kind() == reflect.Pointer {
		return val.Elem()
	}

	return val
}
