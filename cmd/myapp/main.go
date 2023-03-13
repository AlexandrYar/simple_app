package main

import (
	"github.com/AlexandrYar/simple_app/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true
	router.LoadHTMLGlob("tmp/html/*.html")
	router.GET("/register", handler.Register)
	router.POST("/register", handler.Register)
	router.GET("/login", handler.Login)
	router.POST("/login", handler.Login)
	router.GET("/userpage/:login", handler.UserPage)
	router.POST("/userpage/:login", handler.UserPage)
	router.GET("/main_page", handler.MainPage)
	router.POST("/userpage/:login/addItem", handler.AddItem)
	router.GET("/userpage/:login/addItem", handler.AddItem)

	router.Run("localhost:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
