package simplejson

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

func TypeValue(t interface{}) (reflect.Type, reflect.Value) {
	tt := reflect.TypeOf(t)
	tv := reflect.ValueOf(t)
	// 指针类型转换
	if tt.Kind() == reflect.Ptr {
		tv = tv.Elem()
		tt = tv.Type()
	}
	return tt, tv
}

//支持原类型为 string/int/float/number 的转换
func InterfaceToString(inter interface{}) string {
	if inter == nil {
		return ""
	}
	switch inter.(type) {
	case string, *string:
		return reflect.Indirect(reflect.ValueOf(inter)).String()
	case int, int8, int16, int32, int64, *int, *int8, *int16, *int32, *int64:
		return strconv.FormatInt(reflect.Indirect(reflect.ValueOf(inter)).Int(), 10)
	case float32, float64, *float32, *float64:
		return strconv.FormatFloat(reflect.Indirect(reflect.ValueOf(inter)).Float(), 'E', -1, 32)
	case json.Number:
		return inter.(json.Number).String()
	case *json.Number:
		return inter.(*json.Number).String()
	}
	panic(fmt.Sprintf("no support type:%v", reflect.TypeOf(inter)))
}

//支持原类型为 string/int/float/number 的转换
func InterfaceToInt64(inter interface{}) int64 {
	if inter == nil {
		return int64(0)
	}
	switch inter.(type) {
	case string, *string:
		return stringToInt64(reflect.Indirect(reflect.ValueOf(inter)).String())
	case int, int8, int16, int32, int64, *int, *int8, *int16, *int32, *int64:
		return reflect.Indirect(reflect.ValueOf(inter)).Int()
	case json.Number:
		re, err := inter.(json.Number).Int64()
		if err != nil {
			panic("interface no an Int64")
		}
		return re
	case *json.Number:
		re, err := inter.(*json.Number).Int64()
		if err != nil {
			panic("interface no an Int64")
		}
		return re
	}
	panic(fmt.Sprintf("no support type:%v", reflect.TypeOf(inter)))
}

func stringToInt64(s string) int64 {
	v, e := strconv.ParseInt(s, 10, 64)
	if e != nil {
		return 0
	}
	return v
}
