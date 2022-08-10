package oracle

import (
	"github.com/cengsin/oracle"
	"gorm.io/gorm"
)

const (
	Username string = "admin"
	Password string = "admin"
	Host     string = "localhost"
	Database string = "orcl"
)

func NewConnection() (*gorm.DB, *gorm.DB) {
	conn, err := gorm.Open(oracle.Open("c##admin/secret@localhost:1521/orcl"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return conn, nil
}
