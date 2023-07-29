package main

import (
	"fmt"
	"golang-gin-note/mdemo2"
	"golang-gin-note/routers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"

	"github.com/gin-gonic/gin"
)

func routersGroup(r *gin.Engine) {
	routers.DefaultRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.AdminRoutersInit(r)
	routers.UsersRouters(r)
	routers.Test(r)
	routers.NavsRouters(r)
	routers.ArticleRoutersInit(r)
	routers.StudentRoutersInit(r)
	routers.BankRoutersInit(r)
}

func initMiddleware(r *gin.Context) {
	fmt.Println("全局中间件")
}
func initMiddlewareTwo(r *gin.Context) {
	fmt.Println("全局中间件2")
}

func main() {
	// 创建一个默认的路由引擎
	e := gin.Default() // 默认中加载了Logger(), Recovery() 中间件
	e.Use(initMiddleware)
	// todo: 如果要返回html模版,必须设置模版目录 html模版放到templates下
	// 如果你在templates文件夹下有一个html 你可以设置/templates/*
	// 如果你在templates文件夹下有一个html,还有一个文件夹，文件夹下有html 你可以设置/templates/*/*
	// 如果你在templates文件夹下有一层文件夹,文件夹下有html 你可以设置/templates/**/*
	e.LoadHTMLGlob("templates/*/*")
	// mdemo1.TestRouteMain(e)
	mdemo2.TestPostRouteMain(e)
	// 监听并在 0.0.0.0:8000 上启动服务 不传默认8080

	// 配置Session中间件 创建基于cookie的存储引擎 secret 是加密的密钥
	// store := cookie.NewStore([]byte("secret")) // store是存储引擎 这里使用的是cookie
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret")) // 使用redis localhost:6379是本地的redis地址
	e.Use(sessions.Sessions("mysession", store))
	routersGroup(e)
	e.Run(":80")
}
