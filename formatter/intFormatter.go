package formatter

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type intFormatter struct{}

func (i *intFormatter) FormatterType() string {
	return ValidTypeInt
}

func (i *intFormatter) Parse(ctx *gin.Context, rule Rule) interface{} {
	key := rule["name"].(string)
	def, ok := rule["default"]
	str := ctx.Query(key)
	if len(str) == 0 {
		if !ok && rule["required"].(bool) {
			ctx.AbortWithStatusJSON(400, GetParamRequireError(key).Error())
		}
		return def.(int)
	}
	value, err := strconv.Atoi(str)
	if err != nil {
		ctx.AbortWithStatusJSON(400, err.Error())
	}
	return value
}
