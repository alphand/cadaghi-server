package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alphand/skilltree-server/middleware"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/urfave/negroni"
)

type fakeGen struct {
}

func (f *fakeGen) New() interface {
	IDGenerator() string
} {
	return &concreteGen{}
}

type concreteGen struct{}

func (c *concreteGen) IDGenerator() string {
	return c.test()
}

func (c *concreteGen) test() string {
	return "abc123"
}

func TestContext(t *testing.T) {
	Convey("Context is setup", t, func() {
		rr := httptest.NewRecorder()

		// ctx := middleware.NewContext(&middleware.FakeGen{Content: "abc123"})

		fgen := &fakeGen{}
		congen := fgen.New()
		// var _ middleware.IDGenerator = (*concreteGen)(nil)

		// m, ok := interface{}(congen).(middleware.IDGenerator)

		// gen := &fakeGen{
		// 	Gen: congen,
		// }

		// log.Println("fakeme details: ", m, ok, gen.Gen.idGenerator())
		ctx := middleware.NewContext(congen)
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
