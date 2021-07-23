package routers

import (
	"demo1/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login",&controllers.LoginController{})
    beego.Router("/category",&controllers.CategoryController{})
    beego.Router("/topic",&controllers.TopicController{})
    beego.AutoRouter(&controllers.TopicController{})
    beego.ErrorController(&controllers.ErrorController{})
    beego.Router("/reply",&controllers.ReplyController{})
    beego.AutoRouter(&controllers.ReplyController{})
    beego.Router("/user",&controllers.UserController{})
    beego.AutoRouter(&controllers.UserController{})
}
