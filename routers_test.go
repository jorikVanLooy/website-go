package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestShowLoginPageAuthenticated(t *testing.T) {
	w := httptest.NewRecorder()

	r := getRouter(true)

	http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})

	r.GET("/u/login", ensureNotLoggedIn(), showLoginPage)

	req, _ := http.NewRequest("GET", "/u/login", nil)
	req.Header = http.Header{"Cookie": w.Result().Header["Set-Cookie"]}

	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusUnauthorized)
}

func TestShowLoginPageNotAuthenticated(t *testing.T) {
	w := httptest.NewRecorder()

	r := getRouter(true)

	r.GET("/u/login", ensureNotLoggedIn(), showLoginPage)

	req, _ := http.NewRequest("GET", "/u/login", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestShowRegisterPageAuthenticated(t *testing.T) {
	w := httptest.NewRecorder()

	r := getRouter(true)

	http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})

	r.GET("/u/register", ensureNotLoggedIn(), showRegisterPage)

	req, _ := http.NewRequest("GET", "/u/register", nil)
	req.Header = http.Header{"Cookie": w.Result().Header["Set-Cookie"]}

	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusUnauthorized)
}

func TestShowRegistPageNotAuthenticated(t *testing.T) {
	w := httptest.NewRecorder()

	r := getRouter(true)

	r.GET("/u/register", ensureNotLoggedIn(), showRegisterPage)

	req, _ := http.NewRequest("GET", "/u/register", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestShowMessagePageAuthenticated(t *testing.T) {
	w := httptest.NewRecorder()

	r := getRouter(true)

	http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})

	r.GET("/u/message", ensureLoggedIn(), showMessagePage)

	req, _ := http.NewRequest("GET", "/u/message", nil)
	req.Header = http.Header{"Cookie": w.Result().Header["Set-Cookie"]}

	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)

}

func TestShowMessagePageNotAuthenticated(t *testing.T) {
	w := httptest.NewRecorder()

	r := getRouter(true)

	r.GET("/u/message", ensureLoggedIn(), showMessagePage)
	req, _ := http.NewRequest("GET", "/u/message", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusUnauthorized)
}
