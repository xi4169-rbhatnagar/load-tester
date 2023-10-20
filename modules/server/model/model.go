package model

import (
	"context"

	"github.com/xi4169-rbhatnagar/load-tester/modules/server/model/types"
)

// SubmitRequest adds an entry to the task-queue for the incoming request
// Params:
//
//	req: UserRequest to be worked upon
//	s: Store in which persisting our data
func SubmitRequest(ctx context.Context, req types.UserRequest, s types.UserRequestStore) error {
	return s.Store(ctx, req)
}
