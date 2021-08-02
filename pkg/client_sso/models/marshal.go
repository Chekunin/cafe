package models

import (
	"encoding"
	"encoding/json"
	"reflect"
	"strings"
)

//Marshal API function
func Marshal(data interface{}, tagName string) ([]byte, error) {
	return json.Marshal(Convert(data, tagName))
}

//Convert API function
func Convert(value interface{}, tagName string) interface{} {
	return convert(reflect.ValueOf(value), tagName)
}

func convert(value reflect.Value, tagName string) interface{} {
	if _, ok := value.Interface().(json.Marshaler); ok {
		return value.Interface()
	}
	if _, ok := value.Interface().(encoding.TextMarshaler); ok {
		return value.Interface()
	}

	switch value.Type().Kind() {
	case reflect.Ptr:
		if !value.IsNil() {
			return convert(value.Elem(), tagName)
		}
		return nil
	case reflect.Array, reflect.Slice:
		ret := []interface{}{}
		for i := 0; i < value.Len(); i++ {
			ret = append(ret, convert(value.Index(i), tagName))
		}
		return ret
	case reflect.Map:
		ret := map[interface{}]interface{}{}
		for _, key := range value.MapKeys() {
			ret[key.Interface()] = convert(value.MapIndex(key), tagName)
		}
		return ret
	case reflect.Struct:
		ret := map[string]interface{}{}
		for i := 0; i < value.NumField(); i++ {
			tagValue := value.Type().Field(i).Tag.Get(tagName)
			if tagValue == "" {
				continue
			}
			tagFields := strings.Split(tagValue, ",")
			retVal := convert(value.Field(i), tagName)
			omitempty := false
			for i := 1; i < len(tagFields); i++ {
				if tagFields[i] == "omitempty" {
					omitempty = true
					break
				}
			}
			if omitempty {
				if retVal == nil || isEmptyValue(reflect.ValueOf(retVal)) {
					continue
				}
			}
			ret[tagFields[0]] = retVal
		}
		return ret
	default:
		return value.Interface()
	}
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}
