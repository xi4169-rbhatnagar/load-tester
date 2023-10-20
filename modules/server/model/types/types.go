package types

import "context"

type HTTPRequestVerb string

const (
	POST   HTTPRequestVerb = "POST"
	GET                    = "GET"
	DELETE                 = "DELETE"
	PUT                    = "PUT"
)

type UserRequest struct {
	RequestURL        string          `json:"request_url"`
	RequestVerb       HTTPRequestVerb `json:"request_verb"`
	MaxConcurrentUser int             `json:"max_concurrent_users"`
}
type UserRequestStore interface {
	Store(context.Context, UserRequest) error
}
