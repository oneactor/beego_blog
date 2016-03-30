package main

import (
	"bingo_blog/bootstrap"
	_ "bingo_blog/routers"
	"github.com/astaxie/beego"
)

var (
	bt *bootstrap.Bootstrap
)

func main() {
	bt = bootstrap.NewBootstrap(bootstrap.InitOptions{})
	bt.InitEnv()
	beego.Run()
}
