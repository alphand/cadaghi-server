package main

import (
	"net/http"
	"net/http/httptest"
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

	})
}
