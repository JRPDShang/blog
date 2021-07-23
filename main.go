package main

import (
	"demo1/models"
	_ "demo1/routers"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)
func init(){
	models.RegisterDb()
}
func main() {
	err:=orm.RunSyncdb("default",false,true)
	orm.Debug=true
	if err!=nil{
		fmt.Println(err)
	}
	beego.Run()
}

