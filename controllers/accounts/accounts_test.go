package accounts_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"

	acc "github.com/alphand/skilltree-server/controllers/accounts"
	db "github.com/alphand/skilltree-server/database"
	"github.com/icrowley/fake"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	connStr = "192.168.18.129"
	dbName  = "testapp"
)

var (
	ctx context.Context
)

type mockTransport struct {
	Response string
}

func (t *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	resp := &http.Response{
		Header:     make(http.Header),
		Request:    req,
		StatusCode: http.StatusOK,
	}
	resp.Header.Set("Content-Type", "application/json")
	resp.Body = ioutil.NopCloser(strings.NewReader(t.Response))

	return resp, nil
}

func newMockTransport(jsonResp string) http.RoundTripper {
	return &mockTransport{
		Response: jsonResp,
	}
}

type mockDS struct {
}

func (m *mockDS) Coll() *mgo.Collection {
	return &mgo.Collection{}
}
func (m *mockDS) Create(id bson.ObjectId, v interface{}) (interface{}, error) {
	return v, nil
}

func TestAccountHandler(t *testing.T) {
	Convey("Given the Acc Handler is setup", t, func() {
		rr := httptest.NewRecorder()

		client := http.DefaultClient
		client.Transport = newMockTransport(`{"access_token":"popo"}`)
		context.WithValue(ctx, "HTTPClient", client)

		accHdl := &acc.Handler{
			Context:    ctx,
			OAuth2Conf: &oauth2.Config{},
		}

		Convey("Webserver responding to HELLO", func() {
			req, _ := http.NewRequest("GET", "/", nil)
			accHdl.Hello().ServeHTTP(rr, req)
			So(rr.Code, ShouldEqual, http.StatusOK)
			So(rr.Body.String(), ShouldContainSubstring, "Hello")
		})

		Convey("Webserver Get GHToken", func() {
			codejson := `{"code":"1234"}`

			reader := strings.NewReader(codejson)
			req, _ := http.NewRequest("POST", "/token/github", reader)

			accHdl := &acc.Handler{
				Context:    ctx,
				OAuth2Conf: &oauth2.Config{},
			}

			accHdl.HandleGithubExchange().ServeHTTP(rr, req)

			var token oauth2.Token
			json.NewDecoder(rr.Body).Decode(&token)

			So(rr.Code, ShouldEqual, http.StatusOK)
			So(token.AccessToken, ShouldEqual, "popo")
		})

		Convey("Create new Account based on GH", func() {
			email := fake.EmailAddress()
			regoJson := fmt.Sprintf(
				`{"email":"%s", "token":"%s", "firstName":"%s", "lastName":"%s"}`,
				email,
				fake.SimplePassword(),
				fake.FirstName(),
				fake.LastName(),
			)

			reader := strings.NewReader(regoJson)
			req, _ := http.NewRequest("POST", "/accounts/github", reader)

			mongsess := db.InitMongoSession(connStr)
			dbStore, err := db.NewMongoStore(mongsess, dbName, "Users")

			accHdl := &acc.Handler{
				Context:    ctx,
				OAuth2Conf: &oauth2.Config{},
				UserDS:     dbStore,
			}

			accHdl.RegisterUser().ServeHTTP(rr, req)

			So(err, ShouldBeNil)
			So(rr.Body.String(), ShouldContainSubstring, email)
		})

	})
}
