package bootstrap

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

var (
	Cache           cache.Cache
	log             *logs.BeeLogger
	blogCacheExpire = time.Duration(3600)
)

type InitOptions struct {
}
type Bootstrap struct {
	options InitOptions
}

func NewBootstrap(options InitOptions) *Bootstrap {
	bt := new(Bootstrap)
	bt.options = options
	return bt
}

//init app env
func (bt *Bootstrap) InitEnv() {

	beego.BConfig.AppName = "blog"            //设置app的name
	beego.BConfig.Listen.HTTPPort = 8088      //设置监听的端口
	beego.BConfig.WebConfig.AutoRender = true //是否开启自动加载模板
	beego.BConfig.RouterCaseSensitive = false //路由是否区分大小写
	beego.BConfig.ServerName = "nginx1.8"     //作假
	beego.BConfig.RecoverPanic = true         //异常恢复
	beego.BConfig.WebConfig.StaticExtensionsToGzip = []string{".css", ".js"}
	//initlilazation setup session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "redis" //用redis存储session
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379,100"
	beego.BConfig.WebConfig.Session.SessionName = "bingo_blog"
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = int64(3600)
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 3600

	//globalSessions, _ := session.NewManager("redis", `{"cookieName":"gosessionid","gclifetime":3600,"ProviderConfig":"127.0.0.1:6379,100"}`)
	//go globalSessions.GC()

	// setup log init  beego 支持 console、file、conn、smtp
	//Using redis storage log
	os.Mkdir("./log", os.ModePerm)
	if err := beego.BeeLogger.SetLogger(`file`, `{"filename":"./log/log.log"}`); err != nil {
		beego.Error(`Initialization failed log settings:%s`, err)
	}
	//beego.BeeLogger.Async()

	// log = logs.NewLogger(10000)
	// if err := log.SetLogger(`file`, `{"filename":"./log/log.log"}`); err != nil {
	// 	log.Error(`Initialization failed log settings:%s`, err)
	// }

	//###################################################
	//#####        初始化数据库                         ###
	//###################################################
	DB_USER := beego.AppConfig.String("mysql::DB_USER")
	DB_PASS := beego.AppConfig.String("mysql::DB_PASS")
	DB_PORT := beego.AppConfig.String("mysql::DB_PORT")
	DB_HOST := beego.AppConfig.String("mysql::DB_HOST")
	DB_NAME := beego.AppConfig.String("mysql::DB_NAME")

	//beego.AppConfig.Int64("mysql::DB_MaxOpenConns")
	DB_MaxOpenConns, err := beego.AppConfig.Int("mysql::DB_MaxOpenConns") //最大空闲连接数
	if err != nil {
		beego.Error(`get DB_MaxOpenConns err : %s`, err)
		DB_MaxOpenConns = int(200) //获取失败自动设置为默认值
		beego.Error(`setup DB_MaxOpenConns default %d`, 200)
	}

	DB_MaxIdleConns, err := beego.AppConfig.Int("mysql::DB_MaxIdleConns") //最大连接数
	if err != nil {
		beego.Error(`get DB_MaxIdleConns err : %s`, err)
		DB_MaxOpenConns = int(30) //获取失败自动设置为默认值
		beego.Error(`setup DB_MaxIdleConns default %d`, 30)
	}

	DB_Link := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME) + "&loc=Asia%2FChongqing"
	if err := orm.RegisterDataBase("default", "mysql", DB_Link, DB_MaxIdleConns, DB_MaxOpenConns); err != nil {
		beego.Error(`Database connection failed:%s`, err)
	}

	//初始化blog缓存机制  //cache 封装好的方法在cachefunc.go
	if Cache, err = cache.NewCache("redis", `{"key":"gobin_blog","conn":"127.0.0.1:6379","dbNum":"1","password":""}`); err != nil {
		beego.Error(`Failed to initialize cache:%s`, err)
	}
	//Check the configuration file for the last time.
	if err := Cache.StartAndGC(`{"key":"gobin_blog","conn":"127.0.0.1:6379","dbNum":"1","password":""}`); err != nil {
		beego.Error(`cache Configuration file has a problem, please check:%s`, err)
	}
}

//将log 开放接口出去
// func (bt *Bootstrap) Log() *logs.BeeLogger {
// 	return log
// }

//将Cache 开放接口
func (bt *Bootstrap) Cache() cache.Cache {
	return Cache
}
