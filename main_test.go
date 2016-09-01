package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWebserver(t *testing.T) {
	Convey("Given the webserver is started", t, func() {
		Convey("Webserver responding to HELLO", func() {
			req, _ := http.NewRequest("GET", "/hello", nil)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(helloFunc)

			handler.ServeHTTP(rr, req)

			log.Println("body", rr.Body.String())

			So(rr.Code, ShouldEqual, http.StatusOK)
			So(rr.Body.String(), ShouldContainSubstring, "Hello")

		})
	})
}
