package math_go

import (
	"github.com/RandomEstimate/util_go/common"
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
	return reflect.ValueOf(d.Data).Index(0).Interface().(Series).Len()
}

func (d DataFrame) ColLen() int {
	return reflect.ValueOf(d.Data).Len()
}

func (d DataFrame) Col(c int) Series {
	return reflect.ValueOf(d.Data).Index(c).Interface().(Series)
}

func (d DataFrame) Row(r int) Series {
	s := make(Series, 0)
	for i := 0; i < d.ColLen(); i++ {
		s = append(s, d.Col(i).Index(r))
	}
	return s
}

func (d *DataFrame) ColSet(col []string) {
	if len(col) != d.ColLen() {
		panic("length is not consistence.")
	}

	d.Column = col
}

func (d DataFrame) Loc(rowRange [2]int, colName []string) DataFrame {
	if rowRange[0] < 0 || rowRange[1] > d.Len() {
		panic("exceed row range")
	}

	for _, col := range colName {
		if common.Index(col, d.Column) == -1 {
			panic("don't exist the column name.")
		}
	}

	newD := DataFrame{}
	s := make([]Series, 0, len(colName))
	for i := 0; i < d.ColLen(); i++ {
		if common.Index(d.Column[i], colName) != -1 {
			s = append(s, d.Col(i)[rowRange[0]:rowRange[1]])
		}
	}
	newD.Data = s
	newD.ColSet(colName)

	return newD

}

func (d DataFrame) Concat1(s Series) {
	if s.Len() != d.ColLen() {
		panic("length is not consistence.")
	}

	for i := 0; i < d.ColLen(); i++ {
		d.Data[i] = append(d.Data[i], s.Index(i))
	}

}
