package utils

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/kennygrant/sanitize"
	"strings"
)

func ConvertToValidationError(err error, key string) validation.Errors {
	return validation.Errors{key: err}
}

func SanitizeString(stringParam string) string {

	cleanedString := stringParam
	// trim spaces
	cleanedString = strings.TrimSpace(cleanedString)

	// strip html tags
	cleanedString = sanitize.HTML(cleanedString)

	// remove special chars (non ascii chars)
	cleanedString = sanitize.Accents(cleanedString)

	return cleanedString
}
