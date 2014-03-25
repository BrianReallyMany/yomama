package options

import (
	"reflect"
)

type Options struct {
	fields map[string]reflect.Value
}

func NewOptions(optionsStruct interface{}) *Options {
	options := &Options{}

	// Build the map
	options.fields = make(map[string]reflect.Value)
	mapFieldRecursive(&options.fields, "options", reflect.ValueOf(optionsStruct).Elem())

	return options
}

func mapFieldRecursive(fieldMap *map[string]reflect.Value, name string, field reflect.Value) {
	if field.Kind() == reflect.Struct {
		for i := 0; i < field.NumField(); i++ {
			mapFieldRecursive(fieldMap, field.Type().Field(i).Name, field.Field(i))
		}
	} else if (field.CanSet()) { // It's a settable field
		(*fieldMap)[name] = field
	}
}
