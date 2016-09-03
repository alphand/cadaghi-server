package middleware

import (
	"net/http"

	uuid "github.com/satori/go.uuid"

	"golang.org/x/net/context"
)

type key int

const reqIDKey key = 0

var generator IDGenerator

// IDGenerator - interface to specify generator type
type IDGenerator interface {
	IDGenerator() string
}

//UUIDGen - UUID Generator for context request id
type UUIDGen struct{}

// IDGenerator - uuid generator string
func (u *UUIDGen) IDGenerator() string {
	return uuid.NewV4().String()
}

// Context - is a middleware handle context creation to setup request ID, user id and DB
type Context struct {
	idgen IDGenerator
}

// IDGenerator - context id generator
func (c *Context) IDGenerator() string {
	return c.idgen.IDGenerator()
}

func (c *Context) getRequestID(req *http.Request) string {
	return req.Header.Get("X-Request-ID")
}

func (c *Context) setupContextReqID(ctx context.Context, reqID string) context.Context {
	return context.WithValue(ctx, reqIDKey, reqID)
}

func (c *Context) ServeHTTP(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	reqID := c.getRequestID(req)

	if reqID == "" {
		reqID = c.IDGenerator()
		rw.Header().Add("X-Request-ID", reqID)
	}

	ctx := c.setupContextReqID(req.Context(), reqID)
	next(rw, req.WithContext(ctx))
}

// NewContext - create new middleware
func NewContext(idGen IDGenerator) *Context {
	generator = idGen
	return &Context{
		idgen: generator,
	}
}
