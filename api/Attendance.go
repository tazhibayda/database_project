package api

import (
	"database_project/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Attendance struct {
	Course_id  int       `json:"course_id"`
	Student_id int       `json:"student_id"`
	AttDate    time.Time `json:"attDate"`
	Stamp      string    `json:"stamp"`
}

func CreateAttendance(context *gin.Context) {
	if auth.IsAdmin() {

		course_id, _ := strconv.Atoi(context.PostForm("course"))

		student_id, _ := strconv.Atoi(context.PostForm("student"))
		date, _ := time.Parse("2006-01-02", context.PostForm("date"))
		stamp := context.PostForm("stamp")

		if stamp == "on" {
			stamp = "Y"
		} else {
			stamp = "N"
		}
		var attendance = Attendance{Course_id: course_id,
			Student_id: student_id,
			AttDate:    date,
			Stamp:      stamp}

		auth.DB.Create(&attendance)

		location := url.URL{Path: "/admin/attendance"}
		context.Redirect(http.StatusSeeOther, location.RequestURI())
	} else {
		context.Redirect(http.StatusSeeOther, "")
	}
}
func GetAttendanceByCourseID(context *gin.Context) {
	//if auth.IsAdmin() {
	attendances := []Attendance{}
	auth.DB.Where("Course_id", context.Param("course_id")).Find(&attendances)
	//context.HTML(http.StatusOK, "Teacher.gohtml", gin.H{
	//	"teachers": teacher,
	//})
	context.IndentedJSON(http.StatusOK, &attendances)
	//} else {
	//	context.Redirect(http.StatusSeeOther, "")
	//}
}

func GetAttendanceByStudentID(context *gin.Context) {

	attendances := []Attendance{}
	auth.DB.Distinct().Joins("join coursestudent cs on attendances.course_id = cs.course_id and attendances.student_id = cs.student_id").
		Where("cs.student_id = ? and cs.course_id = ?", context.Param("student_id"), context.Param("course_id")).
		Find(&attendances)
	context.IndentedJSON(http.StatusOK, &attendances)

}

func SetAttendance(context *gin.Context) {
	students := []Student{}
	auth.DB.Model(&Student{}).
		Select("students.student_id,students.first_name,students.last_name,students.login,students.password, students.birth_date,students.registredon,students.lastlogin").
		Joins("left join coursestudent on coursestudent.student_id = students.student_id "+
			"join courses c on c.course_id = coursestudent.course_id").
		Where("c.course_id = ?", context.Param("course_id")).
		Scan(&students)

	for _, student := range students {
		id := student.Student_id
		date, _ := time.Parse("2006-01-02", context.PostForm("date"))
		course := context.Param("course_id")
		s := fmt.Sprintf("%d", id)
		stamp := context.Param(s)
		if stamp == "on" {
			stamp = "Y"
		} else {
			stamp = "N"
		}
		auth.DB.Exec("Insert into attendances(course_id, student_id, stamp, att_date) values(?,?,?,?)", course, id, stamp, date)
	}

}

func GetAttendance(context *gin.Context) {
	if auth.IsAdmin() {

		attendances := []Attendance{}
		auth.DB.Find(&attendances)
		context.HTML(http.StatusOK, "Attendance.gohtml", gin.H{
			"attendance": attendances,
		})
	} else {
		context.Redirect(http.StatusSeeOther, "")
	}
}
