package gin_validation

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/shuxnhs/gin-validation/formatter"
)

type ValidatorParam interface {
	Rules(ctx *gin.Context, ruleMap map[string]formatter.Rule) map[string]interface{}
}

type BaseController struct{}

// 检验参数实现
func (bc *BaseController) Rules(ctx *gin.Context, ruleMap map[string]formatter.Rule) map[string]interface{} {
	inputMap := make(map[string]interface{})
	mapTemplate := make(map[string]interface{})
	for param, rule := range ruleMap {
		if err := formatter.CheckRule(rule); err != nil {
			ctx.AbortWithStatusJSON(503, err.Error())
		}
		queryName := rule["name"].(string)
		switch rule["type"] {
		case formatter.ValidTypeInt:
			inputMap[param] = formatter.IntFormatter.Parse(ctx, rule)
		case formatter.ValidTypeBool:
			inputMap[param] = formatter.BoolFormatter.Parse(ctx, rule)
		case formatter.ValidTypeFloat:
			inputMap[param] = formatter.FloatFormatter.Parse(ctx, rule)
		case formatter.ValidTypeMap:
			inputMap[param] = ctx.QueryMap(queryName)
		case formatter.ValidTypeString:
			inputMap[param] = formatter.StringFormatter.Parse(ctx, rule)
		case formatter.ValidTypeStrings:
			inputMap[param] = ctx.QueryArray(queryName)
		}
		mapTemplate[param] = rule["rule"]
	}
	if _, err := govalidator.ValidateMap(inputMap, mapTemplate); err != nil {
		ctx.AbortWithStatusJSON(400, err.Error())
	}
	return inputMap
}
