package model

type HTTPRequestVerb string

const (
	POST   HTTPRequestVerb = "POST"
	GET                    = "GET"
	DELETE                 = "DELETE"
	PUT                    = "PUT"
)

type UserRequest struct {
	RequestURL  string          `json:"request_url"`
	RequestVerb HTTPRequestVerb `json:"request_verb"`
}
