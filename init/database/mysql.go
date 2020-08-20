package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
	"learn-go/conf"
)

var Eloquent *gorm.DB

type Mysql struct {
}

func (e *Mysql) Setup() {
	var err error
	var db Database

	db = new(Mysql)

	Eloquent, err := db.Open(db.GetDriver(), db.GetConnect())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("connect success")
	}

	if Eloquent.Error != nil {
		fmt.Println(Eloquent.Error)
	}

	Eloquent.LogMode(true)
}

// 打开数据库连接
func (e *Mysql) Open(dbType string, conn string) (db *gorm.DB, err error) {
	return gorm.Open(dbType, conn)
}

// 获取数据库连接
func (e *Mysql) GetConnect() string {
	return conf.DatabaseConfig.Source
}

func (e *Mysql) GetDriver() string {
	return conf.DatabaseConfig.Driver
}
