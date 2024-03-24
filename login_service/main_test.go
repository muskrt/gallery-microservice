package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginCheckSuccess(t *testing.T) {
	router := newServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/check/admin/toor", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

}
func TestLoginCheckError(t *testing.T) {
	router := newServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/check/mustafa/toor", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)

}
