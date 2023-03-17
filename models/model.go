package models

import (
	"database/sql"
	"fiber/pkg/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                error
		dbName, user, password, host, port string
	)

	dbSec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'database': %v", err)
	}

	dbName = dbSec.Key("NAME").String()
	user = dbSec.Key("USER").String()
	password = dbSec.Key("PASSWORD").String()
	host = dbSec.Key("HOST").String()
	port = dbSec.Key("PORT").String()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	mysqlDB, err := db.DB()
	mysqlDB.SetMaxOpenConns(10)
	mysqlDB.SetMaxIdleConns(10)

}

func CloseDB(mysqlDB *sql.DB) {
	defer mysqlDB.Close()
}
