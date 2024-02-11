package main

import "github.com/gin-gonic/gin"

func initRoutes(router *gin.Engine) {
	router.Use(setUserStatus())

	router.GET("/", showIndexPage)

	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/login", ensureNotLoggedIn(), showLoginPage)

		userRoutes.GET("/register", ensureNotLoggedIn(), showRegisterPage)

		userRoutes.GET("/message", ensureLoggedIn(), showMessagePage)

		userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)

		userRoutes.POST("/logout", ensureLoggedIn(), performLogout)
	}
}
