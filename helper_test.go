package esimplejson

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
