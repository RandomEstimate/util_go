package math_go

import (
	"reflect"
)

func (s Series) Len() int {
	return reflect.ValueOf(s).Len()
}

func (s Series) Index(i int) interface{} {
	return reflect.ValueOf(s).Index(i).Interface()
}

func (r Rolling) Len() int {
	return reflect.ValueOf(r).Len()
}

func (r Rolling) Index(i int) interface{} {
	return (interface{})(reflect.ValueOf(r).Index(i).Interface())
}
