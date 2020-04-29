package simplejson

import (
	"encoding/json"
	"reflect"
)

func ENewJsonRequired(body []byte) *Json {
	resultJson, err := NewJson(body)
	if err != nil {
		panic(err)
	}
	return resultJson
}

func EMustNewJson(body []byte) *Json {
	resultJson, err := NewJson(body)
	if err != nil {
		return New()
	}
	return resultJson
}

func EFrom(data interface{}) (*Json, error) {
	sourceJson, e := json.Marshal(data)
	if e != nil {
		return nil, e
	}
	return NewJson(sourceJson)
}

func EFromRequired(data interface{}) *Json {
	sourceJson, e := json.Marshal(data)
	if e != nil {
		panic(e)
	}
	return EMustNewJson(sourceJson)
}

func EMustFrom(data interface{}) *Json {
	sourceJson, e := json.Marshal(data)
	if e != nil {
		return nil
	}
	return EMustNewJson(sourceJson)
}

func (j *Json) ESet(key string, val interface{}) {
	// 自动拆包
	if data, ok := val.(Json); ok {
		j.Set(key, data.Interface())
	}
	if data, ok := val.(*Json); ok {
		j.Set(key, data.Interface())
	}
	// 默认装填
	j.Set(key, val)
}

func (j *Json) EGet(key string) *Json {
	return j.Get(key)
}

func (j *Json) EGetIndex(idx int) *Json {
	sourceJson := j.GetIndex(idx)
	return sourceJson
}

func (j *Json) EMustArray(args ...[]interface{}) []interface{} {
	if data, err := j.Array(); err == nil {
		return data
	}

	def := make([]interface{}, 0)
	if len(def) == 1 {
		def = args[0]
	}

	if j.Interface() == nil {
		return def
	}

	tt, tv := TypeValue(j.Interface())
	if tt.Kind() != reflect.Array && tt.Kind() != reflect.Slice {
		return def
	}

	data := make([]interface{}, 0)
	for i := 0; i < tv.Len(); i++ {
		data = append(data, tv.Index(i).Interface())
	}
	return data
}

func (j *Json) EMustMap(args ...map[string]interface{}) map[string]interface{} {
	if data, err := j.Map(); err == nil {
		return data
	}

	def := make(map[string]interface{})
	if len(def) == 1 {
		def = args[0]
	}

	if j.Interface() == nil {
		return def
	}

	tt, tv := TypeValue(j.Interface())
	if tt.Kind() != reflect.Map {
		return def
	}

	data := make(map[string]interface{})
	for _, key := range tv.MapKeys() {
		data[key.String()] = tv.MapIndex(key).Interface()
	}
	return data
}

func (j *Json) EMustString(args ...string) string {
	if data, err := j.String(); err == nil {
		return data
	}

	def := ""
	if len(args) == 1 {
		def = args[0]
	}

	if j.Interface() == nil {
		return def
	}

	return InterfaceToString(j.Interface())
}

func (j *Json) EMustInt64(args ...int64) int64 {
	if data, err := j.Int64(); err == nil {
		return data
	}

	def := int64(0)
	if len(args) == 1 {
		def = args[0]
	}

	if j.Interface() == nil {
		return def
	}

	return InterfaceToInt64(j.Interface())
}

func (j *Json) EMustInt(args ...int) int {
	if data, err := j.Int(); err == nil {
		return data
	}

	def := 0
	if len(args) == 1 {
		def = args[0]
	}

	if j.Interface() == nil {
		return def
	}

	data64 := InterfaceToInt64(args)
	data := int(data64)
	return data
}

func (j *Json) EMustInterface(args ...interface{}) interface{} {
	if j.Interface() != nil {
		return j.Interface()
	}

	var def interface{} = nil
	if len(args) == 1 {
		def = args[0]
	}
	return def
}

func (j *Json) EToJson() (string, error) {
	bs, err := j.MarshalJSON()
	return string(bs), err
}

func (j *Json) EToJsonRequired() string {
	bs, err := j.MarshalJSON()
	if err != nil {
		panic(err)
	}
	return string(bs)
}
