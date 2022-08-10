package api

import (
	"database_project/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"time"
)

type Student struct {
	//gorm.Model
	Student_id  int       `json:"student_id"`
	First_name  string    `json:"first_name"`
	Last_name   string    `json:"last_name"`
	Login       string    `json:"login"`
	Password    string    `json:"password"`
	Birth_date  time.Time `json:"birth_Date"`
	Registredon time.Time `json:"registredon"`
	Lastlogin   time.Time `json:"lastlogin"`
}

func CreateStudent(context *gin.Context) {
	if auth.IsAdmin() {
		first := context.PostForm("FirstName")
		last := context.PostForm("LastName")
		login := context.PostForm("Login")
		password := context.PostForm("Password")
		date, _ := time.Parse("2006-01-02", context.PostForm("Birth"))
		reg := time.Now()

		var student = Student{
			Student_id:  -1,
			First_name:  first,
			Last_name:   last,
			Login:       login,
			Password:    password,
			Birth_date:  date,
			Registredon: reg,
		}

		auth.DB.Create(&student)
		context.IndentedJSON(http.StatusOK, &student)
		//location := url.URL{Path: "/student"}
		//context.Redirect(http.StatusSeeOther, location.RequestURI())
	} else {
		context.Redirect(http.StatusSeeOther, "")
	}
}

func GetStudentByID(context *gin.Context) {
	//if auth.IsAdmin() {
	student := Student{}
	auth.DB.Where("student_id", context.Param("student_id")).Find(&student)
	//context.HTML(http.StatusOK, "Teacher.gohtml", gin.H{
	//	"teachers": teacher,
	//})
	context.HTML(http.StatusOK, "Personal.gohtml", gin.H{
		"student": student,
	})
	//context.IndentedJSON(http.StatusOK, &student)
	//} else {
	//	context.Redirect(http.StatusSeeOther, "")
	//}
}

func GetStudentsByCourseID(context *gin.Context) {

	students := []Student{}
	//auth.DB.Joins("CourseStudent").Joins("Courses").Find(&students)
	auth.DB.Model(&Student{}).
		Select("students.student_id,students.first_name,students.last_name,students.login,students.password, students.birth_date,students.registredon,students.lastlogin").
		Joins("left join coursestudent on coursestudent.student_id = students.student_id "+
			"join courses c on c.course_id = coursestudent.course_id").
		Where("c.course_id = ?", context.Param("course_id")).
		Scan(&students)
	context.HTML(http.StatusOK, "Course-Att.gohtml", gin.H{
		"students": students,
	})
	//context.IndentedJSON(http.StatusOK, &students)
}

func GetStudents(context *gin.Context) {
	if auth.IsAdmin() {
		students := []Student{}
		auth.DB.Order("student_id asc").Find(&students)
		context.HTML(http.StatusOK, "Student.gohtml", gin.H{
			"student": students,
		})
	} else {
		context.Redirect(http.StatusSeeOther, "")
	}
}

func DeleteStudent(context *gin.Context) {

	if auth.IsAdmin() {
		var student Student
		auth.DB.Where("student_id = ?", context.Param("id")).Delete(&student)
		location := url.URL{Path: "/student"}
		context.Redirect(http.StatusTemporaryRedirect, location.RequestURI())
	} else {
		context.Redirect(http.StatusSeeOther, "")
	}
}
