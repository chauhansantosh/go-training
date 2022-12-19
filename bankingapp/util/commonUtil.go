package util

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"reflect"
	"strings"
	"fmt"
)

func ValidateRequest(c *gin.Context, reqObj interface{}) (errorList []string, err error) {
	var errorlist []string
	v := validator.New()

	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(v, trans)

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	err = v.Struct(reqObj)

	if err != nil {
		validatorErrs := err.(validator.ValidationErrors)
		for _, e := range validatorErrs {
			translatedErr := fmt.Sprintf(e.Translate(trans))
			errorlist = append(errorlist, translatedErr)
		}
		return errorlist, err
	}
	return errorlist, nil
}
