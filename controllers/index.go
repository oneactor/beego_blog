package controllers

import (
	_ "bingo_blog/bootstrap"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/context"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	this.TplName = "index/index.html"
}
