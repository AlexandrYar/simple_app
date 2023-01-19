package handler

import (
	"example/mymodule/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	login := c.PostForm("login")
	password := c.PostForm("password")
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	email := c.PostForm("email")
	date_of_birth := c.PostForm("date_of_birth")
	db.Register(db.Connection(), login, password, first_name, last_name, email, date_of_birth)
	c.HTML(200, "register.html", gin.H{
		"Login": login,
	})
}

func Login(c *gin.Context) {
	login := c.PostForm("login")
	password := c.PostForm("password")
	message, isRedirect := db.FindLogin(db.Connection(), login, password)
	if isRedirect {
		log.Print("AAAAAAAAAAAAAAAAAAAAAAAAAA")
		c.Redirect(http.StatusTemporaryRedirect, "/userpage")
	}
	c.HTML(http.StatusTemporaryRedirect, "login.html", gin.H{
		"Message": message,
	})
}
