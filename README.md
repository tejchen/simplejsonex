## go-simplejson-enhancer
go-simplejson-enhancer 是一个Go库，用于直接操作JSON数据

功能点：

- 提供最佳性能
- 提供更丰富、更"舒服"的API
- 支持直接操作JSON，不需要预先定义结构
- 支持动态结构（尽最大可能的所调即所得）
- 支持自动拆包自动装包
- 支持使用原生API

## 快速开始

```go
package main

import (
	"fmt"
	"github.com/tejchen/go-simplejson-enhancer"
)

func main() {
	sourceJson := `{
		"array": [1, "2", 3],
		"map": [
			{
				"mapKey1": 1, 
				"mapKey2": 2
			}
		],
		"string": "stringContent",
		"numberString": "123",
		"number": 1024,
	}`
	esj := simplejson.ENewJsonRequired([]byte(sourceJson))
	esj.EGet("string").EMustString()
	esj.EGet("number").EMustString()  // 数据支持自动转换，所调即所得
	esj.EGet("number").EMustInt64()   // 数据支持自动转换，所调即所得
	for idx := range esj.EGet("array").EMustArray() {
		item := esj.EGet("array").EGetIndex(idx)
		fmt.Println(item.EMustInt64())
	}
}
```

## 基于 bitly/go-simplejson 进行增强
### 增强思路
#### 痛点（详细痛点实例可见下方）
- 设值不友好
    - 设值时无法自动拆包
    - 设值时必须是interface类型才可以被重新读取，无法自动支持。 ![详见](wrapper_test.go)
- 取值必须强指定类型
    - 对于动态类型，无法通过API读取
        - 如某个字段可能是Int也可能是String
        
PS：本人由于使用Go语言处理Python老系统的历史数据，深受其限制，当然这也是我封装该库的初衷

#### 解决思路
- 自动拆解包
- 提供无 ERR 回参的 API（EMust*、E*Required 系列的API）
- 所调即所得，提供自动转换类型的API（EMust* 系列的API）

### 方法约定
- E* 开头表示该方法是本库新增增强方法
    - EMust* 表示该方法会自动转换类型，
        - 失败时，返回默认值或空值
    - E*Required 表示该方法封装了 ERR 处理，出现异常会 Panic

### 原生 API
原生 API 见 [gopkgdoc](http://godoc.org/github.com/bitly/go-simplejson)

### 痛点场景

#### 设置不支持重新通过API获取的场景

```go
package main

import (
	"fmt"
	"github.com/tejchen/go-simplejson-enhancer"
)

func main() {
    data := make([]string, 0)
    data = append(data, "1")
    data = append(data, "2")
	esj := simplejson.ENew()
	esj.Set("data", data)
	for idx := range esj.Get("array").MustArray() {
		// 无法执行以下
		// 原因是 MustArray 底层经过 interface.([]interface) 处理，无法转换成功
        // Map 同理
        esj.GetIndex(idx)
	}
    
    // 改进方法
    // 使用增强方法
    for idx := range esj.EGet("array").EMustArray() {
        esj.GetIndex(idx)
    }
}
```

#### 设置不支持自动拆包
