package api

import (
	"database_project/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"time"
)

type Teacher struct {
	Teacher_id  int       `json:"teacher_id"`
	First_name  string    `json:"first_name"`
	Last_name   string    `json:"last_name"`
	Login       string    `json:"login"`
	Password    string    `json:"password"`
	Registredon time.Time `json:"registredon"`
	Lastlogin   time.Time `json:"lastlogin"`
}

func CreateTeacher(context *gin.Context) {
	if auth.IsAdmin() {

		first := context.PostForm("FirstName")
		last := context.PostForm("LastName")
		login := context.PostForm("Login")
		password := context.PostForm("Password")
		reg := time.Now()

		var teacher = Teacher{
			Teacher_id:  -1,
			First_name:  first,
			Last_name:   last,
			Login:       login,
			Password:    password,
			Registredon: reg,
		}

		auth.DB.Create(&teacher)

		location := url.URL{Path: "/teacher"}
		context.Redirect(http.StatusSeeOther, location.RequestURI())
	} else {
		context.Redirect(http.StatusSeeOther, "")
	}
}

func GetTeachers(context *gin.Context) {
	if auth.IsAdmin() { // if isAdmin opens all teachers
		teachers := []Teacher{}
		auth.DB.Order("teacher_id asc").Find(&teachers)
		context.HTML(http.StatusOK, "Teacher.gohtml", gin.H{
			"teachers": teachers,
		})
	} else { // else you'll redirect to url "admin/"
		context.Redirect(http.StatusSeeOther, "")
	}
}
func GetTeacherByID(context *gin.Context) {
	//if auth.IsAdmin() {
	teacher := Teacher{}
	auth.DB.Where("Teacher_id", context.Param("teacher_id")).Find(&teacher)
	//context.HTML(http.StatusOK, "Teacher.gohtml", gin.H{
	//	"teachers": teacher,
	//})
	context.IndentedJSON(http.StatusOK, &teacher)
	//} else {
	//	context.Redirect(http.StatusSeeOther, "")
	//}
}

func DeleteTeacher(context *gin.Context) {
	if auth.IsAdmin() {

		var teacher Teacher
		auth.DB.Where("teacher_id = ?", context.Param("id")).Delete(&teacher)
		location := url.URL{Path: "/teacher"}
		context.Redirect(http.StatusTemporaryRedirect, location.RequestURI())
	} else {
		context.Redirect(http.StatusSeeOther, "")
	}
}
