package middleware

import (
	"net/http"

	"golang.org/x/net/context"

	db "github.com/alphand/skilltree-server/database"
	mgo "gopkg.in/mgo.v2"
)

//MgoSessionMW - mongo session declaration
type MgoSessionMW struct {
	session *mgo.Session
	connStr string
}

func (m *MgoSessionMW) setMongoCtx(ctx context.Context, sess *mgo.Session) context.Context {
	return context.WithValue(ctx, GetMgoSessKey(), sess)
}

func (m *MgoSessionMW) ServeHTTP(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	ctx := req.Context()
	mgoSess := db.InitMongoSession(m.connStr)

	ctx = m.setMongoCtx(ctx, mgoSess)
	next(rw, req.WithContext(ctx))
}

//CloseSession - close mongo session
func (m *MgoSessionMW) CloseSession() {
	m.session.Close()
}

//GetSession - get mongo session
func (m *MgoSessionMW) GetSession() *mgo.Session {
	return m.session
}

//NewMongoSessionMW - Create new Mongo session middleware
func NewMongoSessionMW(connStr string) *MgoSessionMW {
	return &MgoSessionMW{
		connStr: connStr,
	}
}

//GetMongoSessionFromCtx - Retrieve original mongo session from context
func GetMongoSessionFromCtx(ctx context.Context) *mgo.Session {
	return ctx.Value(GetMgoSessKey()).(*mgo.Session)
}
