package validators

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate   *validator.Validate
	translator ut.Translator
)

func init() {
	validate = validator.New()

	en := en.New()
	uni := ut.New(en, en)
	translator, _ = uni.GetTranslator("en")

	enTranslations.RegisterDefaultTranslations(validate, translator)

	registerCustomMessages()
}

func registerCustomMessages() {
	validate.RegisterTranslation("required", translator, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is required!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	validate.RegisterTranslation("min", translator, func(ut ut.Translator) error {
		return ut.Add("min", "{0} must be at least {1} characters long", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("min", fe.Field(), fe.Param())
		return t
	})
}

func ValidateStruct(s interface{}) map[string]string {
	err := validate.Struct(s)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		errorMessages := make(map[string]string)
		for _, e := range errs {
			errorMessages[e.Field()] = e.Translate(translator)
		}
		return errorMessages
	}
	return nil
}
