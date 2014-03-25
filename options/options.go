package options

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"
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

func (o *Options) Read(r io.Reader) error {
	scanner := bufio.NewScanner(r)

	lineNum := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		splitLine := bytes.Split(line, []byte{'='})

		if len(splitLine) != 2 {
			return errors.New("Line "+strconv.Itoa(lineNum)+" has invalid syntax:"+string(line))
		}

		name := string(bytes.Trim(splitLine[0], " \n"))
		value := string(bytes.Trim(splitLine[1], " \n"))

		err := o.SetOption(name, value)
		if err != nil {
			return err
		}

		lineNum++
	}

	return nil
}

func (o *Options) Write(w io.Writer) {
	for name, value := range o.fields {
		w.Write([]byte(name+"="+fmt.Sprint(value.Interface())+"\n"))
	}
}

func (o *Options) SetOption(name, value string) error {
	kind := o.fields[name].Kind()

	switch kind {
	case reflect.Bool:
		if val, err := strconv.ParseBool(value); err == nil {
			o.fields[name].SetBool(val)
		} else {
			return err
		}
		break
	case reflect.String:
		o.fields[name].SetString(value)
		break
	case reflect.Int:
		if val, err := strconv.ParseInt(value, 10, 32); err == nil {
			o.fields[name].SetInt(val)
		} else {
			return err
		}
		break
	case reflect.Int8:
		if val, err := strconv.ParseInt(value, 10, 8); err == nil {
			o.fields[name].SetInt(val)
		} else {
			return err
		}
		break
	case reflect.Int16:
		if val, err := strconv.ParseInt(value, 10, 16); err == nil {
			o.fields[name].SetInt(val)
		} else {
			return err
		}
		break
	case reflect.Int32:
		if val, err := strconv.ParseInt(value, 10, 32); err == nil {
			o.fields[name].SetInt(val)
		} else {
			return err
		}
		break
	case reflect.Int64:
		if val, err := strconv.ParseInt(value, 10, 64); err == nil {
			o.fields[name].SetInt(val)
		} else {
			return err
		}
		break
	case reflect.Uint:
		if val, err := strconv.ParseUint(value, 10, 32); err == nil {
			o.fields[name].SetUint(val)
		} else {
			return err
		}
		break
	case reflect.Uint8:
		if val, err := strconv.ParseUint(value, 10, 8); err == nil {
			o.fields[name].SetUint(val)
		} else {
			return err
		}
		break
	case reflect.Uint16:
		if val, err := strconv.ParseUint(value, 10, 16); err == nil {
			o.fields[name].SetUint(val)
		} else {
			return err
		}
		break
	case reflect.Uint32:
		if val, err := strconv.ParseUint(value, 10, 32); err == nil {
			o.fields[name].SetUint(val)
		} else {
			return err
		}
		break
	case reflect.Uint64:
		if val, err := strconv.ParseUint(value, 10, 64); err == nil {
			o.fields[name].SetUint(val)
		} else {
			return err
		}
		break
	}

	return nil
}

// Helper function to recursively go through a reflected value and map names to field values
func mapFieldRecursive(fieldMap *map[string]reflect.Value, name string, field reflect.Value) {
	if field.Kind() == reflect.Struct {
		for i := 0; i < field.NumField(); i++ {
			mapFieldRecursive(fieldMap, field.Type().Field(i).Name, field.Field(i))
		}
	} else if (field.CanSet()) { // It's a settable field
		(*fieldMap)[name] = field
	}
}
