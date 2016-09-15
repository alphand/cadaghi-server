package middleware

import (
	"net/http"

	mgo "gopkg.in/mgo.v2"

	db "github.com/alphand/skilltree-server/database"
	uuid "github.com/satori/go.uuid"

	"golang.org/x/net/context"
)

type key int

const reqIDKey key = 0
const mongoConnKey key = 1

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
	idgen        IDGenerator
	mongoConnStr string
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

func (c *Context) setupMongoConn(ctx context.Context, connStr string) context.Context {
	sess, err := db.NewSession(connStr)
	if err != nil {
		panic(err)
	}
	return context.WithValue(ctx, mongoConnKey, sess)
}

func (c *Context) ServeHTTP(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	reqID := c.getRequestID(req)

	if reqID == "" {
		reqID = c.IDGenerator()
		rw.Header().Add("X-Request-ID", reqID)
	}

	ctx := req.Context()

	ctx = c.setupContextReqID(ctx, reqID)
	ctx = c.setupMongoConn(ctx, c.mongoConnStr)
	next(rw, req.WithContext(ctx))

	defer GetMongoConn(ctx).Close()
}

// NewContext - create new middleware
func NewContext(idGen IDGenerator, mongoConnStr string) *Context {
	generator = idGen
	return &Context{
		idgen:        generator,
		mongoConnStr: mongoConnStr,
	}
}

//GetMongoConn - Get mongo session from context
func GetMongoConn(ctx context.Context) *mgo.Session {
	return ctx.Value(mongoConnKey).(*mgo.Session)
}
