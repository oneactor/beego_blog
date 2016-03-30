package routers

import (
	"bingo_blog/controllers"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/test", &controllers.MainController{})
	beego.Router("/category", &controllers.MainController{}, "*:GetCategoryInfo")
	beego.Router("/content", &controllers.MainController{}, "*:GetBlogInfo")

}
