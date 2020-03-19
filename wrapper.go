package esimplejson

import (
	"github.com/tejchen/go-simplejson-enhancer/go-simplejson"
	"io"
	"log"
)

type EJson simplejson.Json

func NewJson(body []byte) (*simplejson.Json, error) {
	return simplejson.NewJson(body)
}

func New() *simplejson.Json {
	return simplejson.New()
}

func NewFromReader(r io.Reader) (*simplejson.Json, error) {
	return simplejson.NewFromReader(r)
}

func (esj *EJson) EMustArray(args ...[]interface{}) []interface{} {
	def := make([]interface{}, 0)

	switch len(args) {
	case 0:
	case 1:
		def = args[0]
	default:
		log.Panicf("MustArray() received too many arguments %d", len(args))
	}

	tt, tv := TypeValue(esj)

	return def
}
