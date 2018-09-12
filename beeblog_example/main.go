package main

import (
	"dothis.top/example/beeblog_example/models"
	_ "dothis.top/example/beeblog_example/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbUser := beego.AppConfig.String("mysql_user")
	dbPwd := beego.AppConfig.String("mysql_pwd")
	dbHost := beego.AppConfig.String("mysql_host")
	dbPort := beego.AppConfig.String("mysql_port")
	dbName := beego.AppConfig.String("mysql_db_name")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPwd, dbHost, dbPort, dbName)
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库
	orm.RegisterDataBase("default", "mysql", dataSource)
}

func main() {
	//开启ORM调试模式
	orm.Debug = true
	//自动建表
	orm.RunSyncdb("default", false, true)

	db := orm.NewOrm()

	category := new(models.Category)
	category.Title = "golang"

	fmt.Println(db.Insert(category))
	beego.Run()
}
