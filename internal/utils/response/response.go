package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GetError(err error, code string) ErrorResponse {
	return ErrorResponse{
		Code: code,
		Message: err.Error(),
	}
}

func GetValidationError(errs validator.ValidationErrors) ErrorResponse {
	var errMessages []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMessages = append(errMessages, fmt.Sprintf("'%s' is required field", err.Field()))
		default:
			errMessages = append(errMessages, fmt.Sprintf("'%s' is invalid", err.Field()))
		}
	}

	return ErrorResponse{
		Code: "invalid_request",
		Message: strings.Join(errMessages, ", "),
	}
}