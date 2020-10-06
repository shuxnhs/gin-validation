package formatter

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type floatFormatter struct{}

func (f *floatFormatter) FormatterType() string {
	return ValidTypeFloat
}

func (f *floatFormatter) Parse(ctx *gin.Context, rule Rule) interface{} {
	key := rule["name"].(string)
	def, ok := rule["default"]
	str := ctx.Query(key)
	if len(str) == 0 {
		if !ok && rule["required"].(bool) {
			ctx.AbortWithStatusJSON(400, GetParamRequireError(key).Error())
		}
		return def.(float64)
	}
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(400, err.Error())
	}
	return value
}
