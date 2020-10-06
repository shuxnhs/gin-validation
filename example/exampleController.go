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
