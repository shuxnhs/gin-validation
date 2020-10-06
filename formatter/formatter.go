package formatter

import (
	"errors"
	"reflect"

	"github.com/gin-gonic/gin"
)

type Formatter interface {
	FormatterType() string
	Parse(ctx *gin.Context, rule Rule) interface{}
}

func CheckRule(rule Rule) error {
	def, ok := rule["default"]
	if ok && reflect.TypeOf(def).String() != rule["type"].(string) {
		return errors.New("Rule Is Illegal ")
	}
	return nil
}

func GetParamRequireError(key string) error {
	return errors.New("Param " + key + " Is Required ")
}
