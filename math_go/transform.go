package math_go

import (
	"reflect"
)

// []struct 转 dataFrame
func Struct2DataFrame(structureList Series) DataFrame {
	// 检查是否字段相同
	if len(structureList) == 0 {
		panic("len cannot equal to 0 .")
	}
	k := reflect.TypeOf(structureList[0])
	colList := make([]string, 0)
	for i := 0; i < k.NumField(); i++ {
		colList = append(colList, k.Field(i).Name)
	}

	d := DataFrame{
		Data:   make([]Series, len(colList), len(colList)),
		Column: colList,
	}

	for i := 0; i < len(structureList); i++ {
		kn := reflect.TypeOf(structureList[i])
		if kn != k {
			panic("type is not consistence.")
		}
		m := make(map[string]interface{}, len(colList))

		v := reflect.ValueOf(structureList[i])

		for j := 0; j < kn.NumField(); j++ {
			m[kn.Field(j).Name] = v.Field(j).Interface()
		}

		tmp := make(Series, 0, len(colList))
		for _, colName := range colList {
			tmp = append(tmp, m[colName])
		}

		d.Concat1(tmp)

		k = kn
	}

	return d

}
