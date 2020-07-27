package validator

import (
	"context"
	"errors"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/howood/kangaroochat/library/utils"
)

// Validator struct
type Validator struct {
	validate *val.Validate
	trans    ut.Translator
	ctx      context.Context
}

// NewValidator creates a new Validator
func NewValidator(ctx context.Context) *Validator {
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator(utils.GetOsEnv("VALIDATE_LANG", "en"))
	I := &Validator{
		validate: val.New(),
		trans:    trans,
		ctx:      ctx,
	}
	en_translations.RegisterDefaultTranslations(I.validate, I.trans)
	return I

}

// Validate process to validate
func (v *Validator) Validate(structData interface{}) error {
	err := v.validate.Struct(structData)
	if err != nil {
		errmsg := []string{}
		errs := err.(val.ValidationErrors)
		for _, e := range errs {
			errmsg = append(errmsg, e.Translate(v.trans))
		}
		return errors.New(strings.Join(errmsg, " / "))
	}
	return nil
}
