package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alphand/skilltree-server/middleware"
	"github.com/urfave/negroni"

	. "github.com/smartystreets/goconvey/convey"
)

type mReqIDGen struct {
}

func (m *mReqIDGen) GenerateID() string {
	return "abc123"
}

func TestRequestIDMiddeware(t *testing.T) {
	Convey("Given Request Middleware can be created", t, func() {
		rr := httptest.NewRecorder()

		mockGen := &mReqIDGen{}
		ridmw := middleware.NewRequestIDMW(mockGen)
		n := negroni.New()
		n.Use(ridmw)

		Convey("WebRequest generating new Request ID", func() {
			req, _ := http.NewRequest("GET", "/", nil)

			n.UseHandler(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.WriteHeader(http.StatusOK)
			}))

			n.ServeHTTP(rr, req)

			So(rr.Header().Get("X-Request-ID"), ShouldEqual, mockGen.GenerateID())
		})

		Convey("WebRequest generating ReUse Request ID", func() {
			req, _ := http.NewRequest("GET", "/", nil)
			req.Header.Set("X-Request-ID", "Old123")

			n.UseHandler(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.WriteHeader(http.StatusOK)
			}))

			n.ServeHTTP(rr, req)

			So(rr.Header().Get("X-Request-ID"), ShouldEqual, "Old123")
		})

	})
}
