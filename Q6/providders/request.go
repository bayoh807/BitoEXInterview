package providders

import (
	"github.com/google/uuid"
	"net/url"
)

type Request struct {
	Query     map[string][]string
	RequestID string
}

func (p *provider) NewRequest(query url.Values) *Request {
	return &Request{
		Query:     query,
		RequestID: uuid.NewString(),
	}
}
