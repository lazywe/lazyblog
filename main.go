package main

import (
	"fmt"
	"lazyblog/functions"
	"lazyblog/models/admin"
	_ "lazyblog/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	functions.Init() //注册函数
	beego.Run()
}

func init() {

	orm.Debug = true
	dbuser := beego.AppConfig.String("db.user")
	dbpass := beego.AppConfig.String("db.pass")
	dbprefix := beego.AppConfig.String("db.prefix")
	dbdatabase := beego.AppConfig.String("db.database")
	dbhost := beego.AppConfig.String("db.host")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dbuser+":"+dbpass+"@tcp("+dbhost+")/"+dbdatabase+"?charset=utf8")
	orm.RegisterModelWithPrefix(dbprefix, new(admin.Admin), new(admin.Node), new(admin.Blog), new(admin.Category), new(admin.Option))
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
	}
}
