package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func showIndexPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.tmpl.html", gin.H{
		"title": "Welcome to Chatty",
	})
}

func showLoginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.tmpl.html", gin.H{
		"title": "Login",
	})
}

func showRegisterPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register.tmpl.html", gin.H{
		"title": "register",
	})
}

func showMessagePage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "message.tmpl.html", gin.H{
		"title": "Send your message",
	})
}

func performLogin(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	if username == "test" && password == "test" {
		ctx.SetCookie("token", "1234", 3600, "", "", false, true)
		ctx.Set("user", "username")
		ctx.Set("is_logged_in", "true")
	}
	ctx.Redirect(http.StatusMovedPermanently, "/")
}

func performLogout(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "", "", false, true)
	ctx.Redirect(http.StatusTemporaryRedirect, "/")
}
