package util

import (
	"fmt"
	"reflect"
	"strconv"
)

func ToMap(in interface{}) (map[string]any, error) {
	out := make(map[string]any)

	v := reflect.ValueOf(in)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct { // Non-structural return error
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i).Name
		val := v.Field(i)

		out[fi] = val

	}

	fmt.Println(out)

	return out, nil
}

func IntToStr(integer int) string {
	return strconv.Itoa(integer)
}

func StrToInt(str string) int {
	integer, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return integer
}

func StructToSliceOfFieldNames(a interface{}) (fieldNameSlice []string) {
	t := reflect.TypeOf(a)

	fieldNameSlice = make([]string, t.NumField())

	for i := range fieldNameSlice {
		fieldNameSlice[i] = t.Field(i).Name
	}

	return fieldNameSlice
}

// func GenerateTableDataFromType(row types.TableRow, t types.Table) (TableData types.TableData) {

// 	TableData = types.TableData{
// 		Fields:    StructToSliceOfFieldNames(row),
// 		Data:      t.CreateMapFromType(row),
// 		TableName: t.GetTableName(),
// 	}

// 	return TableData

// }
