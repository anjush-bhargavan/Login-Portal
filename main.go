package main

import (
	"sample/controllers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	
	r.Static("/static","./static")
	r.LoadHTMLGlob("templates/*.html")

	r.Use(controllers.ClearCache())

	store := cookie.NewStore([]byte("10111"))
	r.Use(sessions.Sessions("login-session", store))
	
	
	r.GET("/login",controllers.LoginPage)
	r.POST("/login",controllers.LoginCheck)
	r.GET("/dashboard",controllers.DashboardPage)
	r.GET("/logout",controllers.Logout)

	r.Run("localhost:8080")
}



