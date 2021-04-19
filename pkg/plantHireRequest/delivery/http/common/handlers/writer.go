package handlers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"gitlab.com/alexgtn/buildit/pkg/plantHireRequest/delivery/http/common/dto"
	"net/http"
)

// WriteErrorMessage wraps message in standard API response
func WriteErrorMessage(w http.ResponseWriter, statusCode int, message string) {
	WriteErrors(w, statusCode, []*dto.Error{dto.NewError(message)})
}

// WriteErrors wraps errors in standard API response
func WriteErrors(w http.ResponseWriter, statusCode int, errors []*dto.Error) {
	res := dto.NewErrorResponse(errors)
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}
