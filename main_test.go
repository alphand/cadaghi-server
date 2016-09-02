package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWebserver(t *testing.T) {
	Convey("Given the webserver is started", t, func() {
		n := SetupNegroni()
		rr := httptest.NewRecorder()

		Convey("Webserver responding to HELLO", func() {
			req, _ := http.NewRequest("GET", "/", nil)

			n.ServeHTTP(rr, req)

			So(rr.Code, ShouldEqual, http.StatusOK)
			So(rr.Body.String(), ShouldContainSubstring, "Hello")
		})

		Convey("Webserver Get GHToken", func() {
			codejson := `{"code":"1234"}`
			reader := strings.NewReader(codejson)

			req, _ := http.NewRequest("POST", "/accounts/github", reader)
			n.ServeHTTP(rr, req)

			log.Println("body", rr.Body.String())

			So(rr.Code, ShouldEqual, http.StatusOK)
		})

	})
}
