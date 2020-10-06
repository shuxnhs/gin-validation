[Chinese](README_CN.md)
# gin-validation

This is a BaseController and Validation for [Gin](https://github.com/gin-gonic/gin) framework.

It uses [govalidator](https://github.com/asaskevich/govalidator) to support the request param validaton. It also provides baseController that other handler struct can combine this BaseController.


## Usage

Download and install using [go module](https://blog.golang.org/using-go-modules):

```sh
export GO111MODULE=on
go get -u github.com/shuxnhs/gin-validation
```

Import it in your code:

```go
import (
    validation "github.com/shuxnhs/gin-validation"
)
```


## Example

Please see [the example Controller](example/exampleController.go) and you can use `Declarative parameter validation` to validator the request data.

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

## explain