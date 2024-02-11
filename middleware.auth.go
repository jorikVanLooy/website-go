package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ensureNotLoggedIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//check if the user is logged in
		loggedInInterface, _ := ctx.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if loggedIn {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func ensureLoggedIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		loggedInInterface, _ := ctx.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if !loggedIn {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

// This middleware sets whether the user is logged in or not
func setUserStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if token, err := ctx.Cookie("token"); err == nil || token != "" {
			ctx.Set("is_logged_in", true)
		} else {
			ctx.Set("is_logged_in", false)
		}
	}
}
