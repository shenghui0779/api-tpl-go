package client

import (
	"net/http"
	"net/url"
	"strings"
)

// Action is the interface that handle wechat api
type Action interface {
	// URL returns request url
	URL(host string) string

	// Method returns action method
	Method() string

	// Body returns body for post request
	Body() ([]byte, error)

	// Decode decodes response
	Decode(b []byte) error
}

type action struct {
	method string
	path   string
	query  url.Values
	body   func() ([]byte, error)
	decode func(resp []byte) error
}

func (a *action) Method() string {
	return a.method
}

func (a *action) URL(host string) string {
	var builder strings.Builder

	builder.WriteString(host)

	if len(a.path) > 0 && a.path[:1] != "/" {
		builder.WriteString("/")
	}

	builder.WriteString(a.path)

	if len(a.query) != 0 {
		builder.WriteString("?")
		builder.WriteString(a.query.Encode())
	}

	return builder.String()
}

func (a *action) Body() ([]byte, error) {
	if a.body != nil {
		return a.body()
	}

	return nil, nil
}

func (a *action) Decode(b []byte) error {
	if a.decode != nil {
		return a.decode(b)
	}

	return nil
}

// NewAction returns a new action
func NewAction(method string, path string, options ...ActionOption) Action {
	a := &action{
		method: method,
		path:   path,
		query:  url.Values{},
	}

	for _, f := range options {
		f(a)
	}

	return a
}

// NewGetAction returns a new action with GET method
func NewGetAction(path string, options ...ActionOption) Action {
	return NewAction(http.MethodGet, path, options...)
}

// NewPostAction returns a new action with POST method
func NewPostAction(path string, options ...ActionOption) Action {
	return NewAction(http.MethodPost, path, options...)
}
