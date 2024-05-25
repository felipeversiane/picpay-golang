package validation

import (
	"encoding/json"
	"errors"
	"github.com/felipeversiane/picpay-golang.git/config/http_error"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateError(
	validation_err error,
) *http_error.HttpError {

	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return http_error.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []http_error.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := http_error.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return http_error.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return http_error.NewBadRequestError("Error trying to convert fields")
	}
}
