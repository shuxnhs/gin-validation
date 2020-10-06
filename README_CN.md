[English](README.md)
# gin-validation

本项目为 [Gin](https://github.com/gin-gonic/gin) 框架的handler进一步封装，提供了参数验证，参数获取等功能.

使用了 [govalidator](https://github.com/asaskevich/govalidator) 去支持参数规则判断，实现一种较为优雅的参数校验获取判断方式`声明式参数校验`.


## 用法

通过 [go module](https://blog.golang.org/using-go-modules) 下载使用:

```sh
export GO111MODULE=on
go get -u github.com/shuxnhs/gin-validation
```

在你的项目中引入:

```go
import (
    validation "github.com/shuxnhs/gin-validation"
)
```


## 用例

你可以参考example示例 [the example Controller](example/exampleController.go) 体验下优雅的参数验证方式 `声明式参数验证` 去对你的请求参数验证，简化掉繁杂的参数校验.

[embedmd]:# (example/exampleController.go go)
```go
package main

import (
	"github.com/gin-gonic/gin"
	validation "github.com/shuxnhs/gin-validation"
	"github.com/shuxnhs/gin-validation/formatter"
)

type ExampleController struct {
	validation.BaseController
}

func (e *ExampleController) Ping(ctx *gin.Context) {
	ruleMap := map[string]formatter.Rule{
		"objectId": {"name": "object_id", "type": formatter.ValidTypeInt, "required": true, "rule": "", "default": 123},
		"objectName": {"name": "object_name", "type": formatter.ValidTypeString, "required": true,
			"rule": "length(1|10),in(string1|string2|...|stringN)"},
	}
	paramMap := e.Rules(ctx, ruleMap)
	if ctx.IsAborted() {
		return
	}
	ctx.JSON(200, paramMap["objectId"])
}
```


## 解析

1. 在我们每一个新的Controller/Handler中的struct都需要去组合validation.BaseController这个结构体，才能使用Rules方法去做参数判断和参数获取。

2. 在每一个api开始我们都需要去构建一个规则的map：map[string]formatter.Rule 存放我们请求参数的规则
    
    + ruleMap的key为每一个请求的参数/参数别名
    
    + ruleMap的value则对应每个参数的规则Rule：formatter.Rule
    
3. formatter.Rule为一个Map，主要有五个key，如：{"name": "object_id", "type": formatter.ValidTypeInt, "required": true, "rule": "", "default": 123}
    
    + name：请求参数名称
    
    + type：参数类型，主要有(int -> formatter.ValidTypeInt, bool -> formatter.ValidTypeBool, float64 -> formatter.ValidTypeFloat
                         string->formatter.ValidTypeString,  strings->formatter.ValidTypeStrings map->formatter.ValidTypeMap)
    
    + required: bool类型，true表示必传，false表示可传
    
    + default：参数默认值，类型必须跟type类型一致
    
    + rule：拓展的规则,下面详解
    
4. 定义好对应的参数规则后，调用 e.Rules(ctx, ruleMap)则会对参数进行验证，如果不满足我们设定的规则则会抛出参数错误的400错误，不会进行下面的业务流程。

5. 调用e.Rules会返回Map，为获取的参数的值，如上文例子中paramMap["objectId"]则为请求参数object_id的值。



> 拓展的rule，可以参考[govalidator](https://github.com/asaskevich/govalidator), 常用规则如下：
> + "range(min|max)":         大小范围,
> + "length(min|max)":        长度范围,
> + "stringlength(min|max)":  字节长度范围,
> + "matches(pattern)":       正则规则,
> + "in(string1|string2|...|stringN)": 范围规则,
> + "email":  email格式,
> + "url":    url格式,
> + "json"    json格式，
> + "ip":     ip格式,
> + "port":   IsPort,
> + "ipv4":   IsIPv4,
> + "ipv6":   IsIPv6,
> + "dns":    IsDNSName,
> + "host":   IsHost,
> + "mac":    IsMAC,
> + "latitude":   IsLatitude,
> + "longitude":  IsLongitude,
> + "IMEI":       IsIMEI,

## 示例截图

1.  object_name的规则required为true但未传参

![](http://cd7.yesapi.net/89E670FD80BA98E7F7D7E81688123F32_20201006215021_a61aab57717122b96bff2e54c342213f.png)

2. 不满足rule规则in(string1|string2|...|stringN)

![](http://cd7.yesapi.net/89E670FD80BA98E7F7D7E81688123F32_20201006215353_b8a822b5aa1ee7664256a024404946da.png)


