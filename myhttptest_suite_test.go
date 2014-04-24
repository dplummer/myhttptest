package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMyhttptest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Myhttptest Suite")
}

var _ = Describe("http", func() {
	It("can route foo", func() {
		recorder := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/foo", nil)

		RunHTTP(recorder, req)

		Expect(recorder.Code).To(Equal(http.StatusOK))
	})

	It("does not route baz", func() {
		recorder := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/baz", nil)

		RunHTTP(recorder, req)

		Expect(recorder.Code).To(Equal(http.StatusNotFound))
	})

	It("requires auth for secret", func() {
		recorder := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/secret", nil)

		RunHTTP(recorder, req)

		Expect(recorder.Code).To(Equal(http.StatusUnauthorized))
	})

	It("allows a password to access the secret", func() {
		recorder := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/secret", nil)
		req.SetBasicAuth("valid", "password")

		RunHTTP(recorder, req)

		Expect(recorder.Code).To(Equal(http.StatusOK))
	})
})
