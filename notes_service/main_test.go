package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginGetRoute(t *testing.T) {
	router := newServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/login", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

}
func TestLoginPostRoute(t *testing.T) {
	router := newServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login/mustafa/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)

}