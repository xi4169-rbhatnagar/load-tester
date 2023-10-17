package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/model"

	"golang.org/x/exp/slog"
)

func HandleSubmit(w http.ResponseWriter, r *http.Request) {
	// Read the Request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		reportError(w, fmt.Sprintf("error reading request body: %v", err), http.StatusBadRequest)
		return
	}

	// Decode request body in a UserRequest
	var request model.UserRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		reportError(w, fmt.Sprintf("error decoding request in UserRequest: %v", err), http.StatusBadRequest)
		return
	}

	err = model.SubmitRequest(request)
	if err != nil {
		reportError(w, fmt.Sprintf("error serving user request: %v", err), http.StatusInternalServerError)
		return
	}
}

func reportError(w http.ResponseWriter, errorMessage string, statusCode int) {
	slog.Error(errorMessage)
	w.WriteHeader(statusCode)
	w.Write([]byte(errorMessage))
}
