package routers

import (
	"fmt"
	"net/http"
	usr "website/users"

	"github.com/gin-gonic/gin"
)

func Routers(db string) *gin.Engine {

	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")
	router.POST("/login", func(c *gin.Context) {
		user := c.PostForm("username")
		password := c.PostForm("password")
		result, id := usr.Login(user, password, db)
		if result {
			println(id)
			c.SetCookie("user", id, 2628288, "/", "localhost", false, true)
			c.Redirect(http.StatusMovedPermanently, "/")
		} else {
			c.Redirect(http.StatusMovedPermanently, "/forbidden")

		}
	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl.html", gin.H{
			"title": "login",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/", func(c *gin.Context) {
		userCookie, err := c.Cookie("user")
		if err != nil {
			fmt.Println(err)
		}
		if userCookie != "0" {
			c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
				"title": "Welcome to chatty",
				"user":  usr.GetUser(userCookie, db),
			})
		} else {
			c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
				"title": "Welcome to chatty",
				"user":  "anonymous",
			})
		}
	})

	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.tmpl.html", gin.H{
			"title": "register",
		})
	})

	router.POST("/register", func(c *gin.Context) {
		user := c.PostForm("username")
		password := c.PostForm("password")
		email := c.PostForm("email")
		phone := c.PostForm("phone")
		id := usr.Register(user, password, email, phone, db)
		fmt.Println(id)
	})

	router.GET("/forbidden", func(c *gin.Context) {
		c.HTML(http.StatusOK, "forbidden.tmpl.html", gin.H{})
	})

	router.GET("/messages", func(c *gin.Context) {
		if userCookie, err := c.Cookie("user"); err != nil {

			c.HTML(http.StatusOK, "login.tmpl.html", gin.H{
				"title": "login",
			})
		} else {

			c.HTML(http.StatusOK, "message.tmpl.html", gin.H{
				"title": "send your message",
				"user":  usr.GetUser(userCookie, db),
			})
		}
	})

	return router
}
