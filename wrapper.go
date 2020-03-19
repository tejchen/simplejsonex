package simplejson

import (
	"encoding/json"
	"github.com/tejchen/go-simplejson-enhancer/go-simplejson"
	"io"
	"reflect"
)

type EJson struct {
	*sourcesimplejson.Json
}

func ENew() *EJson {
	sourceJson := sourcesimplejson.New()
	resultJson := &EJson{
		Json: sourceJson,
	}
	return resultJson
}

func ENewJson(body []byte) (*EJson, error) {
	sourceJson, err := sourcesimplejson.NewJson(body)
	if err != nil {
		return &EJson{}, err
	}
	resultJson := &EJson{
		Json: sourceJson,
	}
	return resultJson, nil
}

func ENewJsonRequired(body []byte) *EJson {
	resultJson, err := ENewJson(body)
	if err != nil {
		panic(err)
	}
	return resultJson
}

func EMustNewJson(body []byte) *EJson {
	resultJson, err := ENewJson(body)
	if err != nil {
		return ENew()
	}
	return resultJson
}

func ENewFromReader(r io.Reader) (*EJson, error) {
	sourceJson, err := sourcesimplejson.NewFromReader(r)
	if err != nil {
		return &EJson{}, err
	}
	resultJson := &EJson{
		Json: sourceJson,
	}
	return resultJson, nil
}

func ENewFromReaderRequired(r io.Reader) *EJson {
	resultJson, err := ENewFromReader(r)
	if err != nil {
		panic(err)
	}
	return resultJson
}

func EMustNewFromReader(r io.Reader) *EJson {
	resultJson, err := ENewFromReader(r)
	if err != nil {
		return ENew()
	}
	return resultJson
}

func EFrom(data interface{}) (*EJson, error) {
	sourceJson, e := json.Marshal(data)
	if e != nil {
		return nil, e
	}
	return ENewJson(sourceJson)
}

func EFromRequired(data interface{}) *EJson {
	sourceJson, e := json.Marshal(data)
	if e != nil {
		panic(e)
	}
	return EMustNewJson(sourceJson)
}

func EMustFrom(data interface{}) *EJson {
	sourceJson, e := json.Marshal(data)
	if e != nil {
		return nil
	}
	return EMustNewJson(sourceJson)
}

func (esj *EJson) ESet(key string, val interface{}) {
	// 自动拆包
	if data, ok := val.(sourcesimplejson.Json); ok {
		esj.Set(key, data.Interface())
	}
	if data, ok := val.(*sourcesimplejson.Json); ok {
		esj.Set(key, data.Interface())
	}
	if data, ok := val.(EJson); ok {
		esj.Set(key, data.Interface())
	}
	if data, ok := val.(*EJson); ok {
		esj.Set(key, data.Interface())
	}
	// 默认装填
	esj.Set(key, val)
}

func (esj *EJson) EGet(key string) *EJson {
	sourceJson := esj.Get(key)
	return &EJson{
		Json: sourceJson,
	}
}

func (esj *EJson) EGetIndex(idx int) *EJson {
	sourceJson := esj.GetIndex(idx)
	return &EJson{
		Json: sourceJson,
	}
}

func (esj *EJson) EMustArray(args ...[]interface{}) []interface{} {
	if data, err := esj.Array(); err == nil {
		return data
	}

	def := make([]interface{}, 0)
	if len(def) == 1 {
		def = args[0]
	}

	if esj.Interface() == nil {
		return def
	}

	tt, tv := TypeValue(esj.Interface())
	if tt.Kind() != reflect.Array && tt.Kind() != reflect.Slice {
		return def
	}

	data := make([]interface{}, 0)
	for i := 0; i < tv.Len(); i++ {
		data = append(data, tv.Index(i).Interface())
	}
	return data
}

func (esj *EJson) EMustMap(args ...map[string]interface{}) map[string]interface{} {
	if data, err := esj.Map(); err == nil {
		return data
	}

	def := make(map[string]interface{})
	if len(def) == 1 {
		def = args[0]
	}

	if esj.Interface() == nil {
		return def
	}

	tt, tv := TypeValue(esj.Interface())
	if tt.Kind() != reflect.Map {
		return def
	}

	data := make(map[string]interface{})
	for _, key := range tv.MapKeys() {
		data[key.String()] = tv.MapIndex(key).Interface()
	}
	return data
}

func (esj *EJson) EMustString(args ...string) string {
	if data, err := esj.String(); err == nil {
		return data
	}

	def := ""
	if len(args) == 1 {
		def = args[0]
	}

	if esj.Interface() == nil {
		return def
	}

	return InterfaceToString(esj.Interface())
}

func (esj *EJson) EMustInt64(args ...int64) int64 {
	if data, err := esj.Int64(); err == nil {
		return data
	}

	def := int64(0)
	if len(args) == 1 {
		def = args[0]
	}

	if esj.Interface() == nil {
		return def
	}

	return InterfaceToInt64(esj.Interface())
}

func (esj *EJson) EMustInt(args ...int) int {
	if data, err := esj.Int(); err == nil {
		return data
	}

	def := 0
	if len(args) == 1 {
		def = args[0]
	}

	if esj.Interface() == nil {
		return def
	}

	data64 := InterfaceToInt64(args)
	data := int(data64)
	return data
}

func (esj *EJson) EMustInterface(args ...interface{}) interface{} {
	if esj.Interface() != nil {
		return esj.Interface()
	}

	var def interface{} = nil
	if len(args) == 1 {
		def = args[0]
	}
	return def
}

func (esj *EJson) EToJson() (string, error) {
	bs, err := esj.MarshalJSON()
	return string(bs), err
}

func (esj *EJson) EToJsonRequired() string {
	bs, err := esj.MarshalJSON()
	if err != nil {
		panic(err)
	}
	return string(bs)
}
