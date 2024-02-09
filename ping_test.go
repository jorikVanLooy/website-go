package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"website/routers"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := routers.Routers("db")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
