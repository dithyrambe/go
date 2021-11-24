package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"lbc/model"

	"lbc/db"
	"lbc/db/sqlite"
)

// SQL is the connexion with the postgres.
type SQL = sqlite.SQlite

func New(dsn string) db.Store {
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	conn.AutoMigrate(&model.User{})

	return &SQL{
		Conn: conn,
	}
}
