package simplejson

import (
	"fmt"
	"testing"
)

func TestEJson_EMustArray(t *testing.T) {
	eJson := New()
	val := make([]int, 0)
	val = append(val, 1)
	val = append(val, 2)
	val = append(val, 3)
	eJson.Set("key", val)
	// 原生方式 取不到这份数据
	if len(eJson.Get("key").MustArray()) > 0 {
		t.FailNow()
	}
	fmt.Println(eJson.Get("key").MustArray())
	// 新方式 可以取到数据
	if len(eJson.EGet("key").EMustArray()) != 3 {
		t.FailNow()
	}
	fmt.Println(eJson.EGet("key").EMustArray())
}

func TestEJson_EMustMap(t *testing.T) {
	eJson := New()
	val := make(map[string]string, 0)
	val["1"] = "1"
	val["2"] = "2"
	eJson.Set("key", val)
	// 原生方式 取不到这份数据
	if len(eJson.Get("key").MustMap()) > 0 {
		t.FailNow()
	}
	fmt.Println(eJson.Get("key").MustMap())
	// 新方式 可以取到数据
	if len(eJson.EGet("key").EMustMap()) != 2 {
		t.FailNow()
	}
	fmt.Println(eJson.EGet("key").EMustMap())
}
