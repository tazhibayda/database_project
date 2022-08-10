package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	Login    = "admin"
	Password = "admin"
)

type Session struct {
	Id       int
	Username string
	Password string
	Type     string
}

var user Session

func Enter(context *gin.Context) {

	if context.Request.Method == "POST" {
		validate(context)
		IsAdmin()
	} else {
		context.HTML(http.StatusOK, "login.gohtml", gin.H{})
	}

}

func validate(context *gin.Context) {

	login := context.PostForm("login")
	password := context.PostForm("password")

	if Login == login && Password == password {
		logining()
		context.HTML(http.StatusOK, "Admin.gohtml", gin.H{
			"username": GetUser(),
		})
	}
	if Login == login && Password != password {
		context.HTML(http.StatusOK, "login.gohtml", gin.H{
			"message": "Invalid password",
		})
	} else {
		context.HTML(http.StatusOK, "login.gohtml", gin.H{
			"message": "Invalid credentials",
		})
	}
}

func logining() {

	DB.Exec("Insert into sessions(id, username,password, type) values(?,?,?,?)", 100, Login, Password, "admin", time.Now()).First(&user)

}

func Outing(context *gin.Context) {
	DB.Exec("delete from sessions where id = ?", 100)
	user = Session{
		Id:       0,
		Username: "",
		Password: "",
		Type:     "",
	}
	IsAdmin()
	context.Redirect(http.StatusTemporaryRedirect, "/admin")
}

func IsAdmin() bool {
	if user.Type == "admin" {
		return true
	}
	return false
}

func GetUser() string {

	DB.First(&user)
	return user.Username

}

func Main(context *gin.Context) {

	context.HTML(http.StatusOK, "Main.gohtml", gin.H{})

}
