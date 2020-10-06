package formatter

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type boolFormatter struct{}

func (b *boolFormatter) FormatterType() string {
	return ValidTypeBool
}

func (b *boolFormatter) Parse(ctx *gin.Context, rule Rule) interface{} {
	key := rule["name"].(string)
	def, ok := rule["default"]
	str := ctx.Query(key)
	if len(str) == 0 {
		if !ok && rule["required"].(bool) {
			ctx.AbortWithStatusJSON(400, GetParamRequireError(key).Error())
		}
		return def.(bool)
	}
	value, err := strconv.ParseBool(str)
	if err != nil {
		ctx.AbortWithStatusJSON(400, err.Error())
	}
	return value
}
