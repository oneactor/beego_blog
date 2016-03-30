package controllers

import (
	_ "bingo_blog/bootstrap"
	//"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) GetBlogInfo() {
	c.Ctx.WriteString("hello world!")
}
func (c *MainController) GetCategoryInfo() {
	c.Ctx.WriteString(`{
    "labels": [
        "GO",
        "PHP",
        "Python",
        "Linux",
        "Elasticsearch"
    ],
    "datasets": [
        {
            "data": [170,50,
               100,
                60,
                90
            ],
            "backgroundColor": [
                "#F7464A",
                "#00CC66",
                "#FDB45C",
                "#66CCFF",
                "#000000",
                "#46BFBD"
            ],
            "hoverBackgroundColor": [
                "#FF5A5E",
                "#00CC66",
                "#FFC870",
                "#66CCFF",
                "#000000",
                "#46BFBD"
            ]
        }
    ]
}`)
}

func (c *MainController) Get() {
	c.Ctx.WriteString(`{
    "result": 1, 
    "count": 7, 
    "data": [
        {
            "title": "马来西亚", 
            "content": "最近老在搞和115网盘有关的东西，博客流量很大一部分都是和115有关的关键词带来的。但是我的标签小程序被封杀之后，我就懒得研究了。反正我又不用115网盘，爱封不封，对我没有丝毫影响。sajsdjalkjdaljdlajfljfljsfljslfjlsjflsdj", 
            "user": "毕滨滨", 
            "timestamp": "2016-12-01 12:23:34", 
            "category": "golang", 
            "from": "content", 
            "image": "/static/img/logo.jpg"
        },
        {
            "title": "马来西亚", 
            "content": "Donec id elit non mi porta gravida at eget metus. Fusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh, ut fermentum massa justo sit amet risus. Etiam porta sem malesuada magna mollis euismod. Donec sed odio dui...", 
            "user": "毕滨滨", 
            "timestamp": "2016-12-01", 
            "category": "golang", 
            "from": "content", 
            "image": "/static/img/logo.jpg"
        },
        {
            "title": "马来西亚", 
            "content": "Donec id elit non mi porta gravida at eget metus. Fusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh, ut fermentum massa justo sit amet risus. Etiam porta sem malesuada magna mollis euismod. Donec sed odio dui...", 
            "user": "毕滨滨", 
            "timestamp": "2016-12-01", 
            "category": "golang", 
            "from": "content", 
            "image": "/static/img/logo.jpg"
        },
        {
            "title": "马来西亚", 
            "content": "Donec id elit non mi porta gravida at eget metus. Fusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh, ut fermentum massa justo sit amet risus. Etiam porta sem malesuada magna mollis euismod. Donec sed odio dui...", 
            "user": "毕滨滨", 
            "timestamp": "2016-12-01", 
            "category": "golang", 
            "from": "content", 
            "image": "/static/img/logo.jpg"
        },
        {
            "title": "马来西亚", 
            "content": "Donec id elit non mi porta gravida at eget metus. Fusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh, ut fermentum massa justo sit amet risus. Etiam porta sem malesuada magna mollis euismod. Donec sed odio dui...", 
            "user": "毕滨滨", 
            "timestamp": "2016-12-01", 
            "category": "golang", 
            "from": "content", 
            "image": "/static/img/logo.jpg"
        },
        {
            "title": "马来西亚", 
            "content": "Donec id elit non mi porta gravida at eget metus. Fusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh, ut fermentum massa justo sit amet risus. Etiam porta sem malesuada magna mollis euismod. Donec sed odio dui...", 
            "user": "毕滨滨", 
            "timestamp": "2016-12-01", 
            "category": "golang", 
            "from": "content", 
            "image": "/static/img/logo.jpg"
        },
        {
            "title": "马来西亚", 
            "content": "Donec id elit non mi porta gravida at eget metus. Fusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh, ut fermentum massa justo sit amet risus. Etiam porta sem malesuada magna mollis euismod. Donec sed odio dui...", 
            "user": "毕滨滨", 
            "timestamp": "2016-12-01", 
            "category": "golang", 
            "from": "content", 
            "image": "/static/img/logo.jpg"
        }
    ]
}
`)
	c.GetString("offset")
	c.GetString("limit")
}
