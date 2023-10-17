package http_model

import (
	"server/model/types"
)

// SubmitRequest adds an entry to the queue for the incoming request for the
// workers to work on it and also persists it to any persistence to be read
// later
//
// Params:
//
//	req: UserRequest to be worked upon
//	queueStore: Store from which the listeners will be
func SubmitRequest(req types.UserRequest, requestStore, queueStore types.Persistence) error {
	return nil
}
