package middleware

// Key - key for context
type Key int

const reqIDKey Key = 0
const iDataStoreKey Key = 1

// GetReqIDKey - Retreive RequestID Context Key
func GetReqIDKey() Key {
	return reqIDKey
}

// GetIDataStoreKey - Retreive IDataStore Context Key
func GetIDataStoreKey() Key {
	return iDataStoreKey
}

// var generator IDGenerator

// // Context - is a middleware handle context creation to setup request ID, user id and DB
// type Context struct {
// 	idgen     IDGenerator
// 	datastore db.IDataStore
// }

// // IDGenerator - context id generator
// func (c *Context) IDGenerator() string {
// 	return c.idgen.IDGenerator()
// }

// func (c *Context) getRequestID(req *http.Request) string {
// 	return req.Header.Get("X-Request-ID")
// }

// func (c *Context) setupContextReqID(ctx context.Context, reqID string) context.Context {
// 	return context.WithValue(ctx, reqIDKey, reqID)
// }

// func (c *Context) ServeHTTP(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
// 	reqID := c.getRequestID(req)

// 	if reqID == "" {
// 		reqID = c.IDGenerator()
// 		rw.Header().Add("X-Request-ID", reqID)
// 	}

// 	ctx := req.Context()
// 	ctx = c.setupContextReqID(ctx, reqID)

// 	next(rw, req.WithContext(ctx))
// }

// // NewContextHandler - create new middleware
// func NewContextHandler(idGen IDGenerator, ds db.IDataStore) *Context {
// 	generator = idGen
// 	return &Context{
// 		idgen:     generator,
// 		datastore: ds,
// 	}
// }

// //GetDatastore - Get mongo session from context
// func GetDatastore(ctx context.Context) db.IDataStore {
// 	return ctx.Value(mongoConnKey).(db.IDataStore)
// }
