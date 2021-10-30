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
	return reflect.ValueOf(r).Index(i).Interface()
}

func (d DataFrame) Len() int {
	if d.ColLen() == 0 {
		return 0
	}
	return reflect.ValueOf(d).Index(0).Interface().(Series).Len()
}

func (d DataFrame) ColLen() int {
	return reflect.ValueOf(d).Len()
}

func (d DataFrame) Col(c int) Series {
	return reflect.ValueOf(d).Index(c).Interface().(Series)
}

func (d DataFrame) Row(r int) Series {
	s := make(Series, 0)
	for i := 0; i < d.ColLen(); i++ {
		s = append(s, d.Col(i).Index(r))
	}
	return s
}
