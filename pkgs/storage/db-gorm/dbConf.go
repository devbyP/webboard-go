package models

import (
	"database/sql"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	SqlDB  *sql.DB
	GormDB *gorm.DB
}

var db Database = Database{}

func ConnectDB() *gorm.DB {
	dns := "host=localhost user=untitled_user password=1234 dbname=post port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	gorm, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot connect to database")
	}
	return gorm
}

func SetGorm(d *gorm.DB) {
	if db.GormDB != nil {
		CloseGorm()
	}
	db.GormDB = d
}

func GetGorm() *gorm.DB {
	return db.GormDB
}

func SetSql(d *sql.DB) {
	if db.SqlDB != nil {
		CloseSqlDB()
	}
	db.SqlDB = d
}

func GetSql() *sql.DB {
	return db.SqlDB
}

func CloseGorm() {
	sql, err := db.GormDB.DB()
	if err != nil {
		log.Fatal("fail to get db from gorm")
	}
	sql.Close()
	db.GormDB = nil
}

func CloseSqlDB() {
	db.SqlDB.Close()
	db.SqlDB = nil
}
