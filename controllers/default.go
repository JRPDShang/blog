package controllers

import (
	"demo1/models"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	models.UpdateCategory()
	models.UpdateReply()
	models.UpdateUser()
	c.TplName = "main.html"
	c.Data["IsHome"]=true
	c.Data["IsLogin"]=checkAccount(c.Ctx)
	u,err:=c.Input()
	if err!=nil {
		panic("err")
	}
	label:=u.Get("label")
	c.Data["Label"]=label
	c.Data["Topics"], _ =models.GetAllTopics(u.Get("cate"),label,true)
	categories,err:=models.GetAllCategories()
	if err!=nil{
		panic(err)
	}
	c.Data["Categories"]=categories
}


