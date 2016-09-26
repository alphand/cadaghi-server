package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alphand/skilltree-server/middleware"
	"github.com/icrowley/fake"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/urfave/negroni"
)

type mockObj struct {
	Name  string
	Email string
}

type mockDS struct{}

func (m *mockDS) GetByID(id string) (interface{}, error) {
	return &mockObj{
		Name:  fake.FirstName(),
		Email: fake.EmailAddress(),
	}, nil
}

func (m *mockDS) Create(o interface{}) error {
	return nil
}

func (m *mockDS) GetAll(q interface{}) ([]interface{}, error) {
	return nil, nil
}

func TestDataStoreMiddleware(t *testing.T) {
	Convey("Given can create datastore middleware", t, func() {
		rr := httptest.NewRecorder()
		mds := &mockDS{}
		dsmw := middleware.NewDataStoreMW(mds)

		n := negroni.New()
		n.Use(dsmw)

		Convey("Datastore Middleware is created", func() {
			var name, email string
			req, _ := http.NewRequest("Get", "/", nil)

			n.UseHandler(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				ds := middleware.GetDSFromCtx(req.Context())

				val, _ := ds.GetByID("")
				name = val.(*mockObj).Name
				email = val.(*mockObj).Email

				rw.WriteHeader(http.StatusOK)
				rw.Write([]byte(name + " " + email))
			}))

			n.ServeHTTP(rr, req)

			So(rr.Body.String(), ShouldEqual, name+" "+email)

		})
	})
}
