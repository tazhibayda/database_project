package api

import (
	"database_project/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strconv"
)

type Course struct {
	Course_id    int    `json:"course_id"`
	Classroom_id int    `json:"classroom_id"`
	Course_Code  string `json:"coursecode"`
	Teacher_id   int    `json:"teacher_id"`
}

func CreateCourse(context *gin.Context) {
	if auth.IsAdmin() {

		classroom, _ := strconv.Atoi(context.PostForm("classroom"))
		code := context.PostForm("code")
		var teacher, _ = strconv.Atoi(context.PostForm("teacher"))

		var course = Course{
			Course_id:    -1,
			Classroom_id: classroom,
			Course_Code:  code,
			Teacher_id:   teacher,
		}
		auth.DB.Create(&course)
		location := url.URL{Path: "admin/course"}
		context.Redirect(http.StatusSeeOther, location.RequestURI())
	} else {
		context.Redirect(http.StatusSeeOther, "")
	}
}

func GetCourseByID(context *gin.Context) {
	//if auth.IsAdmin() {
	course := Course{}
	auth.DB.Where("Course_id", context.Param("course_id")).Find(&course)
	//context.HTML(http.StatusOK, "Teacher.gohtml", gin.H{
	//	"teachers": teacher,
	//})
	context.IndentedJSON(http.StatusOK, &course)
	//} else {
	//	context.Redirect(http.StatusSeeOther, "")
	//}
}

func GetCoursesByTeacherID(context *gin.Context) {

	courses := []Course{}
	auth.DB.Where("Teacher_id", context.Param("teacher_id")).Find(&courses)
	context.IndentedJSON(http.StatusOK, &courses)

}

func GetCoursesByStudentID(context *gin.Context) {

	courses := []Course{}
	auth.DB.Joins("courses join coursestudent cs on cs.course_id = courses.course_id").Where("cs.student_id = ?", context.Param("student_id")).Find(&courses)
	context.IndentedJSON(http.StatusOK, &courses)

}

func ShowAllCourses(context *gin.Context) {
	courses := []Course{}
	auth.DB.Order("course_id asc").Find(&courses)
	context.HTML(http.StatusOK, "AddCourse.gohtml", gin.H{
		"courses": courses,
	})
}

func ShowCourseStudent(context *gin.Context) {
	courses := []Course{}
	auth.DB.Order("course_id asc").Find(&courses)
	context.HTML(http.StatusOK, "AddCourse.gohtml", gin.H{
		"courses": courses,
	})
}

func GetCourse(context *gin.Context) {
	if auth.IsAdmin() {
		courses := []Course{}
		auth.DB.Order("course_id asc").Find(&courses)
		context.HTML(http.StatusOK, "Course.gohtml", gin.H{
			"courses": courses,
		})
	} else {
		context.Redirect(http.StatusSeeOther, "")
	}
}

func DeleteCourse(context *gin.Context) {
	if auth.IsAdmin() {
		var course Course
		auth.DB.Where("course_id = ?", context.Param("id")).Delete(&course)
		location := url.URL{Path: "/course"}
		context.Redirect(http.StatusTemporaryRedirect, location.RequestURI())
	} else {
		context.Redirect(http.StatusSeeOther, "")
	}
}
