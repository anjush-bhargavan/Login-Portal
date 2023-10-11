package controllers

import (
	"net/http"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)



func LoginPage(c *gin.Context){
	session :=sessions.Default(c)
	email :=session.Get("email")
	if email == "sample@email.com"{
		c.HTML(http.StatusOK, "dashboard.html",nil)
		return
	}
	
	c.HTML(http.StatusOK,"login.html",nil)
}


func LoginCheck(c *gin.Context){
	session :=sessions.Default(c)

	email :=c.PostForm("email")
	password :=c.PostForm("password")

	session.Set("email",email)
	session.Set("password",password)
	session.Save()

	if Validate(email,password){
		c.Redirect(http.StatusSeeOther,"/dashboard")
	}else{
		c.HTML(http.StatusBadRequest,"login.html", gin.H{
			"error":"Invalid credentials",
		})
	}
}

func Validate(email, password string) bool{
	return email =="sample@email.com" && password == "sample@123"
}


func DashboardPage(c *gin.Context){
	session :=sessions.Default(c)
	email :=session.Get("email")
	if email != "sample@email.com"{
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}
	c.HTML(http.StatusOK,"dashboard.html",gin.H{
		"email":email,
	})
}

func Logout(c *gin.Context){
	session :=sessions.Default(c)
	session.Delete("email")
	session.Delete("password")
	session.Save()
	
	c.Redirect(http.StatusSeeOther,"/login")
}

func ClearCache() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")

		c.Next()
	}
}
