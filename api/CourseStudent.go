package api

import (
	"database_project/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CourseStudent struct {
	Course_id  int `json:"course_id"`
	Student_id int `json:"student_id"`
}

func AddCourseToStudent(context *gin.Context) {
	auth.DB.Exec("Insert into coursestudent(course_id, student_id) values(?,?)", context.Param("course_id"), context.Param("student_id"))
	context.Redirect(http.StatusSeeOther, "/student/"+context.Param("student_id")+"/all_courses")
}
