package controllers

import (
	"demo1/models"
	beego "github.com/beego/beego/v2/server/web"
)

type ReplyController struct {
	beego.Controller
}

func (c *ReplyController)Add()  {
	u,err:=c.Input()
	if err != nil {
		panic(err)
	}
	tid:=u.Get("tid")
	user,is:=models.CheckUser(c.Ctx)
	var nickname string
	if checkAccount(c.Ctx){
		nickname="JRPDShang"
	}else if is{
		nickname=user.NickName
	}else {
		c.Redirect("/user",302)
		return
	}
	err=models.AddReply(tid,nickname,u.Get("content"))
	if err != nil {
		panic(err)
	}
	c.Redirect("/topic/view/"+tid,302)
}
func (c *ReplyController)Delete()  {
	if !checkAccount(c.Ctx){
		return
	}
	tid:=c.Ctx.Input.Param("0")
	models.DeleteReply(c.Ctx.Input.Param("1"))
	c.Redirect("/topic/view/"+tid,302)
}