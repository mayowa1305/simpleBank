package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/mayowa1305/simpleBank/util"
)

var validCurrency validator.Func = func(fieldlevel validator.FieldLevel) bool {
	if currency, ok := fieldlevel.Field().Interface().(string); ok {
		//check if currency is surported
		return util.IsSupportedCurrency(currency)
	}
	return false
}
