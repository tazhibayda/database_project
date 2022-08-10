package api

import (
	"database_project/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strconv"
)

type Classroom struct {
	Classroom_id int `json:"classroom_id"`
	Capacity     int `json:"capacity"`
}

func CreateClass(context *gin.Context) {
	if auth.IsAdmin() {

		capasity, _ := strconv.Atoi(context.PostForm("capacity"))

		var classroom = Classroom{
			Classroom_id: -1,
			Capacity:     capasity,
		}

		auth.DB.Create(&classroom)

		location := url.URL{Path: "/classroom"}
		context.Redirect(http.StatusSeeOther, location.RequestURI())
	} else {
		context.Redirect(http.StatusSeeOther, "")
	}
}

func GetClassByID(context *gin.Context) {
	//if auth.IsAdmin() {
	classroom := Classroom{}
	auth.DB.Where("Classroom_id", context.Param("classroom_id")).Find(&classroom)
	//context.HTML(http.StatusOK, "Teacher.gohtml", gin.H{
	//	"teachers": teacher,
	//})
	context.IndentedJSON(http.StatusOK, &classroom)
	//} else {
	//	context.Redirect(http.StatusSeeOther, "")
	//}
}

func GetClass(context *gin.Context) {
	if auth.IsAdmin() {

		classrooms := []Classroom{}
		auth.DB.Order("classroom_id asc").Find(&classrooms)
		context.HTML(http.StatusOK, "Classroom.gohtml", gin.H{
			"classrooms": classrooms,
		})
	} else {
		context.Redirect(http.StatusSeeOther, "")
	}
}

func DeleteClass(context *gin.Context) {
	if auth.IsAdmin() {

		var class Classroom
		auth.DB.Where("classroom_id = ?", context.Param("id")).Delete(&class)
		location := url.URL{Path: "/classroom"}
		context.Redirect(http.StatusTemporaryRedirect, location.RequestURI())
	} else {
		context.Redirect(http.StatusSeeOther, "")
	}
}
