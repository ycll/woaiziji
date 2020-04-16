package routers

import (
	"woaiziji/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HomeController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/exit", &controllers.ExitController{})
	beego.Router("/article/add", &controllers.AddArticleController{})
	beego.Router("/article/:id", &controllers.ShowArticleController{})
	beego.Router("/article/update", &controllers.UpdateArticleController{})
}
