package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/xi4169-rbhatnagar/load-tester/modules/server/model"
	"github.com/xi4169-rbhatnagar/load-tester/modules/server/model/types"
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
	var request types.UserRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		reportError(w, fmt.Sprintf("error decoding request in UserRequest: %v", err), http.StatusBadRequest)
		return
	}

	// Validate the user request params
	if err = validateUserRequest(request); err != nil {
		reportError(w, fmt.Sprintf("error validating request: %v", err), http.StatusBadRequest)
		return
	}

	// Serve the request
	if err = model.SubmitRequest(request, nil, nil); err != nil {
		reportError(w, fmt.Sprintf("error serving user request: %v", err), http.StatusInternalServerError)
		return
	}
}

func reportError(w http.ResponseWriter, errorMessage string, statusCode int) {
	slog.Error(errorMessage)
	w.WriteHeader(statusCode)
	w.Write([]byte(errorMessage))
}

func validateUserRequest(req types.UserRequest) error {
	if req.MaxConcurrentUser <= 0 {
		return fmt.Errorf("there must be atleast one concurrent users in the system to be load-testing")
	}

	if _, err := url.ParseQuery(req.RequestURL); err != nil {
		return fmt.Errorf("invalid url: %v", err)
	}

	switch req.RequestVerb {
	case types.POST, types.GET, types.DELETE, types.PUT:
		break
	default:
		return fmt.Errorf("invalid request-verb '%s'", req.RequestVerb)
	}

	return nil
}
