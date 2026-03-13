package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/Buddhadebgarai/student-api/internal/types"
	"github.com/Buddhadebgarai/student-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJsonResponse(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		slog.Info("Creating a student")

		//request validation using go-playground/validator
		validate := validator.New()
		err = validate.Struct(student)

		response.WriteJsonResponse(w, http.StatusCreated, map[string]string{"message": "Student created successfully", "student": student.Name})
	}
}
