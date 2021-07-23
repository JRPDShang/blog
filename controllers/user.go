package controllers

import (
	"demo1/models"
	beego "github.com/beego/beego/v2/server/web"
	"path"
	"strconv"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	u, _ := c.Input()
	isExit := u.Get("IsExit")
	if isExit == "true" {
		c.Ctx.SetCookie("nickname", "", -1, "/")
		c.Ctx.SetCookie("pwd", "", -1, "/")
		c.Redirect("/", 302)
		return
	}
	user, is := models.CheckUser(c.Ctx)
	if checkAccount(c.Ctx) {
		c.TplName = "user.html"
		users, _ := models.GetAllUsers()
		c.Data["Users"] = users
	} else if is {
		c.Redirect("/user/info/"+strconv.FormatInt(user.Id, 10), 302)
	} else {
		c.Redirect("/user/login", 302)
		return
	}
}
func (c *UserController) Add() {
	if c.Ctx.Input.Param("0") == "false" {
		c.Data["FalseCond"] = true
	}
	c.TplName = "user_add.html"
}
func (c *UserController) Login() {
	if c.Ctx.Input.Param("0") == "false" {
		c.Data["FalseCond"] = true
	}
	c.TplName = "user_login.html"
}
func (c *UserController) Post() {
	u, _ := c.Input()
	op := u.Get("op")
	switch op {
	case "add":
		err := models.AddUser(u.Get("nickname"), u.Get("password"))
		if err != nil {
			c.Redirect("/user/add/false", 302)
			return
		}
		c.Redirect("/user/login", 302)
	case "check":
		user, is := models.UserLogin(u.Get("nickname"), u.Get("password"))
		if is {
			c.Ctx.SetCookie("nickname", user.NickName, 1<<31-1, "/")
			c.Ctx.SetCookie("pwd", user.Password, 1<<31-1, "/")
			c.Redirect("/user/info/"+strconv.FormatInt(user.Id, 10), 302)
		} else {
			c.Redirect("/user/login/false", 302)
		}
	case "head":
		user,is:=models.CheckUser(c.Ctx)
		if is {
			_=c.SaveToFile("head",path.Join("static/head",user.NickName+".jpg"))
		}
		c.Redirect("/user/info/"+strconv.FormatInt(user.Id, 10),302)
	}
}
func (c *UserController) Info() {
	c.TplName = "user_info.html"
	user, _ := models.GetUser(c.Ctx.Input.Param("0"))
	c.Data["User"] = user
}
