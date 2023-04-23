package reflectutil

import (
	"errors"
	"reflect"
)

// GetStructTagsValue 得到该结构体 tag 标签对应的值
func GetStructTagsValue(i interface{}, tag string) []string {
	var res []string
	rt := reflect.TypeOf(i)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	if rt.Kind() != reflect.Struct {
		panic("not a struct")
	}
	//通过 json 标签得到表格头
	for i := 0; i < rt.NumField(); i++ {
		res = append(res, rt.Field(i).Tag.Get(tag))
	}
	return res
}

// GetStructFieldName 得到该结构体的全部字段名
func GetStructFieldName(i interface{}) []string {
	var res []string
	rt := reflect.TypeOf(i)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	if rt.Kind() != reflect.Struct {
		panic("not a struct")
	}
	//通过 json 标签得到表格头
	for i := 0; i < rt.NumField(); i++ {
		res = append(res, rt.Field(i).Name)
	}
	return res
}

// GetStructFieldValue 得到该结构体的字段值
func GetStructFieldValue(i interface{}) []interface{} {
	rv := reflect.ValueOf(i)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		panic("not a struct")
	}
	var row []interface{}
	for j := 0; j < rv.NumField(); j++ {
		fieldKind := rv.Field(j).Kind()
		fieldValue := rv.Field(j)
		switch fieldKind {
		case reflect.String:
			row = append(row, fieldValue.String())
		case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
			row = append(row, fieldValue.Int())
		case reflect.Float64, reflect.Float32:
			row = append(row, fieldValue.Float())
		case reflect.Bool:
			row = append(row, fieldValue.Bool())
		}
	}
	return row
}

// GetStructSliceFieldValue 得到 slice 结构体每个字段的值
func GetStructSliceFieldValue(i interface{}) ([][]interface{}, error) {
	rt := reflect.TypeOf(i)
	rv := reflect.ValueOf(i)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
		rv = rv.Elem()
	}
	if rt.Kind() != reflect.Slice {
		panic("not a slice")
	}
	if rv.Len() == 0 {
		return nil, nil
	}
	//通过结构体数组得到表格内容
	var res [][]interface{}
	for i := 0; i < rv.Len(); i++ {
		rvt := rv.Index(i)
		if rv.Index(i).Kind() == reflect.Ptr {
			rvt = rv.Index(i).Elem()
		}
		var row []interface{}
		for j := 0; j < rvt.NumField(); j++ {
			fieldKind := rvt.Field(j).Kind()
			fieldValue := rvt.Field(j)
			switch fieldKind {
			case reflect.String:
				row = append(row, fieldValue.String())
			case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
				row = append(row, fieldValue.Int())
			case reflect.Float64, reflect.Float32:
				row = append(row, fieldValue.Float())
			case reflect.Bool:
				row = append(row, fieldValue.Bool())
			}
		}
		res = append(res, row)
	}
	return res, nil
}

//得到结构体的反射，从 slice 结构体或者结构体指针中 (slice必须 >0 且为结构体)
func getStructReflectInSliceOrPtr(i interface{}) (reflect.Type, reflect.Value, error) {
	rt := reflect.TypeOf(i)
	rv := reflect.ValueOf(i)
	if rt.Kind() != reflect.Slice {
		if rt.Kind() == reflect.Ptr {
			rt = rt.Elem()
		}
		if rt.Kind() != reflect.Struct {
			panic("not a struct")
		}
		return rt, rv, nil
	}
	if rv.Len() == 0 {
		return nil, reflect.Value{}, errors.New("slice is empty")
	}
	if rv.Index(0).Kind() == reflect.Ptr {
		rt = rv.Index(0).Elem().Type()
		rv = rv.Index(0).Elem()
	}
	if rt.Kind() != reflect.Struct {
		panic("not a struct")
	}
	return rt, rv, nil
}
