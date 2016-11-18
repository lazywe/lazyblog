package main

import (
	"fmt"
	"lazyblog/functions"
	"lazyblog/models"
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
	debug, _ := beego.AppConfig.Bool("db.debug")
	orm.Debug = debug
	dbuser := beego.AppConfig.String("db.user")
	dbpass := beego.AppConfig.String("db.pass")
	dbprefix := beego.AppConfig.String("db.prefix")
	dbdatabase := beego.AppConfig.String("db.database")
	dbhost := beego.AppConfig.String("db.host")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dbuser+":"+dbpass+"@tcp("+dbhost+")/"+dbdatabase+"?charset=utf8")
	orm.RegisterModelWithPrefix(dbprefix, new(models.Admin), new(models.Node), new(models.Blog), new(models.Category), new(models.Option), new(models.Recommend), new(models.Menu))
	err := orm.RunSyncdb("default", false, false)
	if err != nil {
		fmt.Println(err)
	}
}
