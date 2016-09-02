package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alphand/skilltree-server/middleware"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/urfave/negroni"
)

func TestContext(t *testing.T) {
	Convey("Context is setup", t, func() {
		rr := httptest.NewRecorder()
		ctx := middleware.NewContext(&middleware.FakeGen{Content: "abc123"})
		n := negroni.New()
		n.Use(ctx)

		Convey("WebRequest is contexted properly", func() {
			req, _ := http.NewRequest("GET", "/", nil)

			n.UseHandler(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.WriteHeader(http.StatusOK)
			}))

			n.ServeHTTP(rr, req)

			So(rr.Header().Get("X-Request-ID"), ShouldEqual, "abc123")
		})

	})
}
