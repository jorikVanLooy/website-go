package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"website/routers"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := routers.Routers(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestLoginRoute(t *testing.T) {
	router := routers.Routers(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/login", nil)
	router.ServeHTTP(w, req)

	p, _ := io.ReadAll(w.Body)

	assert.Equal(t, 200, w.Code)
	assert.Greater(t, strings.Index(string(p), "<title>login</title>"), 0)
}

func TestRegisterRoute(t *testing.T) {
	router := routers.Routers(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/register", nil)
	router.ServeHTTP(w, req)

	p, _ := io.ReadAll(w.Body)

	assert.Equal(t, 200, w.Code)
	assert.Greater(t, strings.Index(string(p), "<title>register</title>"), 0)
}

func TestIndexRoute(t *testing.T) {
	router := routers.Routers(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	p, _ := io.ReadAll(w.Body)

	assert.Equal(t, 200, w.Code)
	assert.Greater(t, strings.Index(string(p), "<title>Welcome to chatty</title>"), 0)
}

func TestMessageNoAuthentication(t *testing.T) {
	router := routers.Routers(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/messages", nil)
	router.ServeHTTP(w, req)

	p, _ := io.ReadAll(w.Body)

	assert.Equal(t, 200, w.Code)
	assert.Greater(t, strings.Index(string(p), "<title>login</title>"), 0)
}

func TestMessageAuthentication(t *testing.T) {
	router := routers.Routers(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/messages", nil)
	router.ServeHTTP(w, req)

	p, _ := io.ReadAll(w.Body)

	assert.Equal(t, 200, w.Code)
	assert.Greater(t, strings.Index(string(p), "<title>send your message</title>"), 0)
}
