package controllers

import (
	"demo1/models"
	beego "github.com/beego/beego/v2/server/web"
	"path"
)

type TopicController struct {
	beego.Controller
}
func (c *TopicController) Get() {
	c.Data["IsTopic"]=true
	c.Data["IsLogin"]=checkAccount(c.Ctx)
	var err error
	c.Data["Topics"],err=models.GetAllTopics("","",false)
	if err != nil {
		panic(err)
	}
	c.TplName = "topic.html"
}
func (c *TopicController)Add()  {
	c.TplName="topic_add.html"
	c.Data["IsLogin"]=checkAccount(c.Ctx)
	var err error
	c.Data["Categories"],err=models.GetAllCategories()
	if err != nil {
		panic(err)
	}
}
func (c *TopicController)Post()  {
	if !checkAccount(c.Ctx){
		c.Redirect("/login",302)
		return
	}
	u,err:=c.Input()
	if err!=nil{
		panic(err)
	}
	tid:=u.Get("id")
	title:=u.Get("title")
	if len(tid)==0{
		err=models.AddTopic(title,u.Get("content"),u.Get("category"),u.Get("label"))

		_=c.SaveToFile("background",path.Join("static/background",title+".jpg"))

	}else {
		err=models.ModifyTopic(tid,u.Get("title"),u.Get("content"),u.Get("category"),u.Get("label"))
	}

	if err != nil {
		panic(err)
	}
	c.Redirect("/topic",301)
}
func (c *TopicController) View()  {
	c.TplName="topic_view.html"
	topic,err:=models.GetTopic(c.Ctx.Input.Param("0"))
	if err!=nil {
		c.Redirect("/topic",302)
	}
	c.Data["Topic"]=topic
	c.Data["IsLogin"]=checkAccount(c.Ctx)
	replies,err:= models.GetAllReplies(c.Ctx.Input.Param("0"))
	if err != nil {
		panic(err)
	}
	c.Data["Replies"]=replies
}
func (c *TopicController)Modify()  {
	c.TplName="topic_modify.html"
	tid:=c.Ctx.Input.Param("0")
	topic,err:=models.GetTopic(tid)
	if err != nil {
		c.Redirect("/",302)
		return
	}
	c.Data["Categories"],err=models.GetAllCategories()
	if err != nil {
		panic(err)
	}
	c.Data["Topic"]=topic
}
func (c *TopicController)Delete(){
	if !checkAccount(c.Ctx){
		c.Redirect("/login",302)
		return
	}
	err:=models.DeleteTopic(c.Ctx.Input.Param("0"))
	if err != nil {
		panic(err)
	}
	c.Redirect("/topic",302)
}