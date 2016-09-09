package accounts_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"

	acc "github.com/alphand/skilltree-server/controllers/accounts"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	ctx context.Context
)

type mockTransport struct {
	Response string
}

func newMockTransport(jsonResp string) http.RoundTripper {
	return &mockTransport{
		Response: jsonResp,
	}
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
			req, _ := http.NewRequest("POST", "/accounts/github", reader)

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

	})
}
