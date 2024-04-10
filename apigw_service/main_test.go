package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginCheckSuccess(t *testing.T) {
	router := newServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/check/test_user/toor", nil)
	fmt.Println(req)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

}
func TestLoginCheckError(t *testing.T) {
	router := newServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/check/test_user/root", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)

}
