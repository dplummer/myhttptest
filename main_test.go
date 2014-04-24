package main

import (
	"github.com/go-martini/martini"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

/* Test Helpers */
func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func Test_Http(t *testing.T) {
	m := martini.New()
	router := martini.NewRouter()
	recorder := httptest.NewRecorder()
	routeMatched := false

	router.Get("/foo", func() {
		routeMatched = true
	})

	req, _ := http.NewRequest("GET", "/foo", nil)
	m.ServeHTTP(recorder, req)
	expect(t, routeMatched, true)
}
