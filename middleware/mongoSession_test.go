package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/alphand/skilltree-server/middleware"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/urfave/negroni"
)

func TestMongoSessionMiddleware(t *testing.T) {
	Convey("MongoSession middleware can be created", t, func() {
		var insess *mgo.Session

		connStr := "192.168.18.129"
		dbName := "testMWDB"

		rr := httptest.NewRecorder()

		mgosessMW := middleware.NewMongoSessionMW(connStr)

		n := negroni.New()
		n.Use(mgosessMW)

		Convey("Can get mongo session in request", func() {
			var upid string
			req, _ := http.NewRequest("Get", "/", nil)

			n.UseHandler(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				sess := middleware.GetMongoSessionFromCtx(req.Context())
				insess = sess.Copy()

				info, _ := insess.DB(dbName).C("mw").UpsertId(bson.NewObjectId().Hex(), struct{ Name string }{Name: "Niko"})
				upid = info.UpsertedId.(string)

				rw.WriteHeader(http.StatusOK)
				rw.Write([]byte(upid))
			}))

			n.ServeHTTP(rr, req)

			So(rr.Body.String(), ShouldEqual, upid)
		})

		Reset(func() {
			insess.DB(dbName).DropDatabase()
			insess.Close()
		})

	})

}
