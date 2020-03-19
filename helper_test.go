package simplejson

import (
	"reflect"
	"testing"
)

func TestTypeValue(t *testing.T) {
	// 指针类型测试
	sample := "110"
	ptrSample := &sample
	tt, tv := TypeValue(sample)
	if tt.Kind() != reflect.String || tv.Kind() != reflect.String {
		t.FailNow()
	}
	tt, tv = TypeValue(ptrSample)
	if tt.Kind() != reflect.String || tv.Kind() != reflect.String {
		t.FailNow()
	}
}

func TestInterfaceToString(t *testing.T) {
	// int8
	i8 := int8(8)
	i8s := InterfaceToString(i8)
	if i8s != "8" {
		t.FailNow()
	}
	// int64
	i64 := int64(64)
	i64s := InterfaceToString(i64)
	if i64s != "64" {
		t.FailNow()
	}
	// *int8
	i8p := &i8
	i8ps := InterfaceToString(i8p)
	if i8ps != "8" {
		t.FailNow()
	}
	// *int64
	i64p := &i64
	i64ps := InterfaceToString(i64p)
	if i64ps != "64" {
		t.FailNow()
	}
}

func TestInterfaceInt64(t *testing.T) {
	// string
	s := "1"
	i64 := InterfaceToInt64(s)
	if i64 != 1 {
		t.FailNow()
	}
	// *string
	sp := &s
	i64 = InterfaceToInt64(sp)
	if i64 != 1 {
		t.FailNow()
	}
	// int32
	i32 := int32(32)
	i64 = InterfaceToInt64(i32)
	if i64 != 32 {
		t.FailNow()
	}
	// *int32
	i32p := &i32
	i64 = InterfaceToInt64(i32p)
	if i64 != 32 {
		t.FailNow()
	}
}
