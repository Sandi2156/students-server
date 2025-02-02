package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	request "github.com/sandipan/students-api/internal/types"
	"github.com/sandipan/students-api/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Creating a student.")
		// decode request body
		var studentPayload request.Student
		err := json.NewDecoder(r.Body).Decode(&studentPayload)

		// if EOF error send bad request
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GetError(errors.New("Payload should contain data"), "empty_body"))
			return
		}
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GetError(err, "invalid_request"))
			return
		}

		// validation
		if err := validator.New().Struct(studentPayload); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.GetValidationError(validationErrors))
			return
		}

		// send status ok
		response.WriteJson(w, http.StatusCreated, map[string] string {"status": "OK"})
	}
}