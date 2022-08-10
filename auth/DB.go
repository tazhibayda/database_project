package auth

import (
	"database_project/pkg/oracle"
)

var DB, err = oracle.NewConnection()

func Init() {
	//DB.AutoMigrate(&api.Course{})
	//DB.AutoMigrate(&api.Teacher{})
	//DB.AutoMigrate(&api.Student{})
	//DB.AutoMigrate(&api.Attendance{})
	//DB.AutoMigrate(&api.Classroom{})
}
