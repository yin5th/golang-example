package models

import (
	"dothis.top/example/gin_blog/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedAt  int `json:"created_at"`
	ModifiedAt int `json:"modified_at"`
}

func init() {
	var (
		err                                                     error
		dbType, dbName, user, password, host, port, tablePrefix string
	)

	section, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'database': %v", err)
	}

	dbType = section.Key("TYPE").String()
	dbName = section.Key("NAME").String()
	user = section.Key("USER").String()
	password = section.Key("PASSWORD").String()
	host = section.Key("HOST").String()
	port = section.Key("PORT").String()
	tablePrefix = section.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, port, dbName))
	if err != nil {
		log.Fatalln(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
