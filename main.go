package main

import (
	"fmt"
	"net/http"
	"github.com/go-martini/martini"
	"strings"
	"encoding/base64"
)

func requireAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", `Basic realm="So secret"`)
	w.WriteHeader(401)
	w.Write([]byte("401 Unauthorized\n"))
}

func extractBasicAuth(r *http.Request) (username string, password string) {
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 || s[0] != "Basic" {
		return
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return
	}
	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return
	}
	username = pair[0]
	password = pair[1]
	return
}

func checkAuth(r *http.Request) string {
	username, password := extractBasicAuth(r)

	if username == "valid" && password == "password" {
		return username
	} else {
		return ""
	}
}

func wrapAuth(wrapped http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := checkAuth(r)
		if username != "" {
			wrapped(w, r)
		} else {
			requireAuth(w, r)
		}
	}
}

func secretHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is so secret")
}

func RunHTTP(res http.ResponseWriter, req *http.Request) {
	m := martini.New()
	router := martini.NewRouter()

	router.Get("/foo", func() string {
		return "bar"
	})

	router.Get("/secret", wrapAuth(secretHandler))

	m.Action(router.Handle)
	m.ServeHTTP(res, req)
}

func main() {
	http.HandleFunc("/", RunHTTP)
	http.ListenAndServe(":8080", nil)
}

