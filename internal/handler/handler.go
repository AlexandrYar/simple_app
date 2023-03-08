package handler

import (
	"net/http"

	"github.com/AlexandrYar/simple_app/internal/db"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	login := c.PostForm("login")
	password := c.PostForm("password")
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	email := c.PostForm("email")
	date_of_birth := c.PostForm("date_of_birth")
	var user db.User
	user.Register(db.ConnDb.Connection(db.NewDb), login, password, first_name, last_name, email, date_of_birth)
	c.HTML(200, "register.html", gin.H{
		"Login": login,
	})
}

func Login(c *gin.Context) {
	login := c.PostForm("login")
	password := c.PostForm("password")
	var user db.User
	message, isRedirect := user.LoginUser(db.ConnDb.Connection(db.NewDb), login, password)
	if isRedirect {
		c.Redirect(http.StatusTemporaryRedirect, "/userpage/"+login)
	}
	c.HTML(http.StatusTemporaryRedirect, "login.html", gin.H{
		"Message": message,
	})
}

func UserPage(c *gin.Context) {
	login_given := c.Params.ByName("login")
	var user db.User
	user.Find_info(db.ConnDb.Connection(db.NewDb), login_given)
	c.HTML(http.StatusOK, "userPage.html", gin.H{
		"Login":         user.Login,
		"First_name":    user.First_name,
		"Last_name":     user.Last_name,
		"Email":         user.Email,
		"Date_of_birth": user.Date_of_birth,
	})
}
