package math_go

import (
	"fmt"
	"github.com/RandomEstimate/util_go/common"
	"math"
	"reflect"
	"strconv"
)

type Series []interface{}

func NewSeries(s interface{}) Series {
	if s := reflect.ValueOf(s); s.Kind() == reflect.Slice {
		series := make(Series, 0, s.Len())
		for i := 0; i < s.Len(); i++ {
			series = append(series, s.Index(i).Interface())
		}
		return series
	}
	return nil
}

func (s Series) Mean() (float64, error) {

	sum := 0.
	for i := 0; i < s.Len(); i++ {
		switch reflect.ValueOf(s.Index(i)).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint,
			reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64:
			sum += common.ParamSlice(strconv.ParseFloat(fmt.Sprint(s.Index(i)), 64))[0].(float64)
		default:
			return 0, fmt.Errorf("type error %v ", reflect.ValueOf(s.Index(i)).Kind())
		}
	}
	return sum / float64(s.Len()), nil

}

func (s Series) Sum() (float64, error) {

	sum := 0.
	for i := 0; i < s.Len(); i++ {
		switch reflect.ValueOf(s.Index(i)).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint,
			reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64:
			sum += common.ParamSlice(strconv.ParseFloat(fmt.Sprint(s.Index(i)), 64))[0].(float64)
		default:
			return 0, fmt.Errorf("type error %v ", reflect.ValueOf(s.Index(i)).Kind())
		}
	}
	return sum, nil

}

func (s Series) Var() (float64, error) {
	sum2 := 0.
	sum := 0.
	for i := 0; i < s.Len(); i++ {
		switch reflect.ValueOf(s.Index(i)).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint,
			reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64:
			x := common.ParamSlice(strconv.ParseFloat(fmt.Sprint(s.Index(i)), 64))[0].(float64)
			sum2 += math.Pow(x, 2)
			sum += x
		default:
			return 0, fmt.Errorf("type error ")
		}
	}
	n := float64(s.Len())
	mean := sum / n

	return (sum2 - n*math.Pow(mean, 2)) / n, nil
}

func (s Series) Std() (float64, error) {
	var_, err := s.Var()
	if err != nil {
		return 0, err
	}
	return math.Sqrt(var_), nil
}

func (s Series) Max() (float64, error) {
	max := 0.
	for i := 0; i < s.Len(); i++ {
		switch reflect.ValueOf(s.Index(i)).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint,
			reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64:
			j := common.ParamSlice(strconv.ParseFloat(fmt.Sprint(s.Index(i)), 64))[0].(float64)
			if max < j {
				max = j
			}
		default:
			return 0, fmt.Errorf("type error ")
		}
	}
	return max, nil
}

func (s Series) Min() (float64, error) {
	min := 0.
	for i := 0; i < s.Len(); i++ {
		switch reflect.ValueOf(s.Index(i)).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint,
			reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64:
			j := common.ParamSlice(strconv.ParseFloat(fmt.Sprint(s.Index(i)), 64))[0].(float64)
			if min > j {
				min = j
			}
		default:
			return 0, fmt.Errorf("type error ")
		}
	}
	return min, nil
}

func (s Series) Agg(f func(x interface{}) interface{}) Series {
	sCopy := make(Series, 0, s.Len())
	for i := 0; i < s.Len(); i++ {
		sCopy = append(sCopy, f(s.Index(i)))
	}
	return (Series)(sCopy)
}

type Rolling []Series

func (s Series) Rolling(period int) Rolling {

	v := reflect.ValueOf(s)
	if v.Len() < period {
		return Rolling{}
	}

	r := make(Rolling, 0)
	for i := 0; i < v.Len(); i++ {
		slice := make([]interface{}, 0)
		if period-(i+1) > 0 {
			slice = Full(period-(i+1), nil)
		}
		for j := int(math.Max(0, float64(i-period+1))); j <= i; j++ {
			slice = append(slice, v.Index(j).Interface())
		}

		r = append(r, slice)
	}

	return (Rolling)(r)

}

func (r Rolling) Mean() Series {
	return r.Agg(func(x interface{}) interface{} {
		x_ := x.(Series)
		x1 := ([]interface{})(x_)
		if common.Index(nil, x1) == -1 {
			return common.ParamSlice(x_.Mean())[0]
		} else {
			return math.NaN()
		}

	})
}

func (r Rolling) Var() Series {
	return r.Agg(func(x interface{}) interface{} {
		x_ := x.(Series)
		x1 := ([]interface{})(x_)
		if common.Index(nil, x1) == -1 {
			return common.ParamSlice(x_.Var())[0]
		} else {
			return math.NaN()
		}

	})
}

func (r Rolling) Std() Series {
	return r.Agg(func(x interface{}) interface{} {
		x_ := x.(Series)
		x1 := ([]interface{})(x_)
		if common.Index(nil, x1) == -1 {
			return common.ParamSlice(x_.Std())[0]
		} else {
			return math.NaN()
		}

	})
}

func (r Rolling) Max() Series {
	return r.Agg(func(x interface{}) interface{} {
		x_ := x.(Series)
		x1 := ([]interface{})(x_)
		if common.Index(nil, x1) == -1 {
			return common.ParamSlice(x_.Max())[0]
		} else {
			return math.NaN()
		}

	})
}

func (r Rolling) Min() Series {
	return r.Agg(func(x interface{}) interface{} {
		x_ := x.(Series)
		x1 := ([]interface{})(x_)
		if common.Index(nil, x1) == -1 {
			return common.ParamSlice(x_.Min())[0]
		} else {
			return math.NaN()
		}

	})
}

func (r Rolling) Agg(f func(x interface{}) interface{}) Series {
	rCopy := make(Series, 0, r.Len())
	for i := 0; i < r.Len(); i++ {
		rCopy = append(rCopy, f(r.Index(i)))
	}
	return (Series)(rCopy)
}

type DataFrame []Series

func NewDataFrame(x ...Series) DataFrame {
	d := make(DataFrame, 0)
	l := -1
	for i := 0; i < len(x); i++ {
		if x[i].Len() != l && l != -1 {
			return nil
		}
		d = append(d, x[i])
		l = x[i].Len()
	}
	return d
}

func (d *DataFrame) Agg(f func(x Series) interface{}) Series {
	series := make(Series, 0, d.Len())
	for i := 0; i < d.Len(); i++ {
		series = append(series, f(d.Row(i)))
	}
	return series
}
