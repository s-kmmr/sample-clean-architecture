package client

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlHandler struct {
	conn *gorm.DB
}

func (s *SqlHandler) Conn() *gorm.DB {
	return s.conn
}

func NewSqlHandler() *SqlHandler {
	dsn := "docker:docker@tcp(127.0.0.1:3306)/prac?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db = db.Debug()

	return &SqlHandler{conn: db}
}
