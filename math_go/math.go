package math_go

import (
	"fmt"
	"github.com/RandomEstimate/util_go/common"
	"reflect"
	"strconv"
)

type Mode = int

const (
	MaxMode Mode = iota
	MinMode
)

// Max
func MaxOrMin(slice interface{}, m Mode) interface{} {
	var Type = []reflect.Kind{
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr,
		reflect.Float32,
		reflect.Float64,
	}

	if slice := reflect.ValueOf(slice); slice.Kind() == reflect.Slice {
		if slice.Len() == 0 {
			return nil
		}

		if max := reflect.ValueOf(slice.Index(0).Interface()); common.Index(max.Kind(), Type) == -1 {
			return nil
		}
		slice2 := make([]float64, 0, slice.Len())
		for i := 0; i < slice.Len(); i++ {
			slice2 = append(slice2, common.ParamSlice(strconv.ParseFloat(fmt.Sprint(slice.Index(i).Interface()), 64))[0].(float64))
		}

		switch m {
		case MaxMode:
			return (interface{})(max(slice2))
		case MinMode:
			return (interface{})(min(slice2))
		}

	}
	return nil
}

// max
func max(slice []float64) float64 {
	m := slice[0]
	for _, v := range slice {
		if m < v {
			m = v
		}
	}
	return m
}

// min
func min(slice []float64) float64 {
	m := slice[0]
	for _, v := range slice {
		if m > v {
			m = v
		}
	}
	return m
}
