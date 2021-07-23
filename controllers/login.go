package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	context2 "github.com/beego/beego/v2/server/web/context"
)
type LoginController struct {
	beego.Controller
}
func (c *LoginController)Get(){
	u,_:=c.Input()
	IsExit:=u.Get("exit")=="true"
	if IsExit{
		c.Ctx.SetCookie("username", "", -1, "/")
		c.Ctx.SetCookie("password","",-1,"/")
		c.Redirect("/",301)
		return
	}
	c.TplName="login.html"
}
func (c *LoginController)Post(){
	u,_:=c.Input()
	uname:=u.Get("username")
	pwd:=u.Get("password")
	auto:=u.Get("AutoLogin")=="on"
	RightName,_:=beego.AppConfig.String("username")
	RightPwd,_:=beego.AppConfig.String("password")
	if uname==RightName&&
		pwd==RightPwd{
		maxAge:=0
		if auto{
			maxAge=1<<31-1
		}
		c.Ctx.SetCookie("username",uname,maxAge,"/")
		c.Ctx.SetCookie("password",pwd,maxAge,"/")
	}
	c.Redirect("/",301)
	return
}
func checkAccount(ctx *context2.Context) bool {
	ck,err:=ctx.Request.Cookie("username")
	if err!=nil{
		return false
	}
	uname:=ck.Value
	ck,err=ctx.Request.Cookie("password")
	if err!=nil{
		return false
	}
	pwd:=ck.Value
	RightName,_:=beego.AppConfig.String("username")
	RightPwd,_:=beego.AppConfig.String("password")
	return uname == RightName && pwd == RightPwd
}