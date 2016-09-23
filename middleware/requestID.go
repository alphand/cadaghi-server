package middleware

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/context"
)

// IReqIDGenerator - interface to specify generator type
type IReqIDGenerator interface {
	GenerateID() string
}

//UUIDGen - UUID Generator for context request id
type UUIDGen struct{}

// GenerateID - uuid generator string
func (u *UUIDGen) GenerateID() string {
	return uuid.NewV4().String()
}

// RequestIDMW - RequestID middleware who will handle request id generation
type RequestIDMW struct {
	generator IReqIDGenerator
}

func (u *RequestIDMW) getRequestID(req *http.Request) string {
	return req.Header.Get("X-Request-ID")
}

func (u *RequestIDMW) setRequestID(rw http.ResponseWriter, reqID string) {
	rw.Header().Add("X-Request-ID", reqID)
}

func (u *RequestIDMW) setReqIDContext(ctx context.Context, reqID string) context.Context {
	return context.WithValue(ctx, GetReqIDKey(), reqID)
}

func (u *RequestIDMW) ServeHTTP(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	reqID := u.getRequestID(req)

	if reqID == "" {
		reqID = u.generator.GenerateID()
	}

	rw.Header().Add("X-Request-ID", reqID)

	ctx := req.Context()
	ctx = u.setReqIDContext(ctx, reqID)

	next(rw, req.WithContext(ctx))
}

//NewRequestIDMW - create new RequestID middleware
func NewRequestIDMW(generator IReqIDGenerator) *RequestIDMW {
	return &RequestIDMW{
		generator: generator,
	}
}
