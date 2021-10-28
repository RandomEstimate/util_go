package common

import (
	"reflect"
)

// ParamSlice
func ParamSlice(param ...interface{}) []interface{} {
	return param
}

// index
func Index(v interface{}, slice interface{}) int {
	if slice := reflect.ValueOf(slice); slice.Kind() == reflect.Slice {
		for i := 0; i < slice.Len(); i++ {
			if reflect.DeepEqual(v, slice.Index(i).Interface()) {
				return i
			}
		}
	}
	return -1
}
