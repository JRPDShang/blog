package controllers

import (
	"demo1/models"
	beego "github.com/beego/beego/v2/server/web"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController)Get()  {
	u,err:=c.Input()
	if err!=nil{
		panic(err)
	}
	op:=u.Get("op")
	switch op {
	case "add":
		var name =u.Get("name")
		if len(name)==0{
			break
		}
		err:=models.AddCategory(name)
		if err != nil {
			panic(err)
		}
		c.Redirect("/category",301)
	case "del":
		var id =u.Get("id")
		if len(id) == 0 {
			break
		}
		err:=models.DelCategory(id)
		if err!=nil {
			panic(err)
		}
	}
	c.TplName="category.html"
	c.Data["IsCategory"]=true
	c.Data["IsLogin"]=checkAccount(c.Ctx)
	c.Data["Categories"],err=models.GetAllCategories()
	if err != nil {
		panic(err)
	}
}