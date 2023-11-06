package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func TranslateError(err error) []string {
	validate := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(validate, trans)

	if err == nil {
		return []string{}
	}

	validatorErrs := err.(validator.ValidationErrors)
	var errs []string

	for _, e := range validatorErrs {
		translatedErr := e.Translate(trans)
		// Hapus "Key: 'FieldName' " dari pesan kesalahan
		translatedErr = strings.Replace(translatedErr, fmt.Sprintf("Key: '%s' ", e.Namespace()), "", -1)
		errs = append(errs, translatedErr)
	}

	return errs
}
