package esimplejson

import (
	"reflect"
)

func TypeValue(t interface{}) (reflect.Type, reflect.Value) {
	tt := reflect.TypeOf(t)
	tv := reflect.ValueOf(t)
	// 指针类型转换成非指针类型
	if tt.Kind() == reflect.Ptr {
		tv = tv.Elem()
		tt = tv.Type()
	}
	return tt, tv
}
