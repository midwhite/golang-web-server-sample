package serializers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func TranslateErrorMessage(err validator.FieldError, modelName string) string {
	switch err.Tag() {
	case "required":
		fieldName := fieldTranslations[modelName+"."+err.Field()]
		return fieldName + "は必須です。"
	default:
		message := fmt.Sprintf("unexpected error type: %+v", err.Tag())
		panic(message)
	}
}

var fieldTranslations = map[string]string{
	"Todo.Title": "件名",
}
