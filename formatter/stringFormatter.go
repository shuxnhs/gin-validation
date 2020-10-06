package formatter

import (
	"github.com/gin-gonic/gin"
)

type stringFormatter struct{}

func (s *stringFormatter) FormatterType() string {
	return ValidTypeString
}

func (s *stringFormatter) Parse(ctx *gin.Context, rule Rule) interface{} {
	key := rule["name"].(string)
	def, ok := rule["default"]
	str := ctx.Query(key)
	if len(str) == 0 {
		if !ok && rule["required"].(bool) {
			ctx.AbortWithStatusJSON(400, GetParamRequireError(key).Error())
		} else if ok {
			return def.(string)
		}
		return ""
	}
	return str
}
