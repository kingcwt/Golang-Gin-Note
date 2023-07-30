# Golang-note


---
title: go-gin
date: 2023-07-06
categories:
 - 前端
tags:
 - golang
---

# go-gin

- 初始化

```go
// 1 新建main.go

go mod init go-gin //2 go-gin是当前文件夹名称
```

- 下载 gin

```go
go get -u github.com/gin-gonic/gin
```

- 启动服务

```go
package main

import "github.com/gin-gonic/gin"

func main() {
	// 创建一个默认的路由引擎
	e := gin.Default()
	e.GET("/", func(c *gin.Context) {
		c.String(200, "值:%v", "hello gin")
	})

	e.GET("/news", func(c *gin.Context) {
		c.String(200, "news")
	})

	// 监听并在 0.0.0.0:8000 上启动服务 不传默认8080
	e.Run(":8000")
}
```

- download - 热加载

```go
go get github.com/pilu/fresh
```

[fresh command - github.com/pilu/fresh - Go Packages](https://pkg.go.dev/github.com/pilu/fresh@v0.0.0-20190826141211-0fa698148017#section-readme)

[https://github.com/gravityblast/fresh](https://github.com/gravityblast/fresh)

关于热加载 会遇到找不到fresh命令的问题

- 1 执行go env 查看当前环境变量,找到GOPATH
- 2 在你项目下执行

```go
// 在你的项目下执行就可以了
export PATH="这里是你的GOPATH对应的路径/bin:$PATH" 

```

永久添加 可以在你的的~/.zshrc里面加入这个path

你可以执行echo $path来查看是否把这个GOPATH添加到环境变量中 

- go env

```go
GO111MODULE="on"
GOARCH="arm64"
GOBIN=""
GOCACHE="/Users/cuihongran/Library/Caches/go-build"
GOENV="/Users/cuihongran/Library/Application Support/go/env"
GOEXE=""
GOEXPERIMENT=""
GOFLAGS=""
GOHOSTARCH="arm64"
GOHOSTOS="darwin"
GOINSECURE=""
GOMODCACHE="/Users/cuihongran/go/pkg/mod"
GONOPROXY=""
GONOSUMDB=""
GOOS="darwin"
GOPATH="/Users/cuihongran/go"
GOPRIVATE=""
GOPROXY="https://proxy.golang.org,direct"
GOROOT="/usr/local/go"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/darwin_arm64"
GOVCS=""
GOVERSION="go1.20.5"
GCCGO="gccgo"
AR="ar"
CC="clang"
CXX="clang++"
CGO_ENABLED="1"
GOMOD="/dev/null"
GOWORK=""
CGO_CFLAGS="-O2 -g"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-O2 -g"
CGO_FFLAGS="-O2 -g"
CGO_LDFLAGS="-O2 -g"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -arch arm64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/qt/qfgjw21x1rl7zwyylqy453mc0000gn/T/go-build4047264937=/tmp/go-build -gno-record-gcc-switches -fno-common"
```

- 在你的项目里运行热加载

```go
fresh
```

### gin - request & template

- main

```go
package main

import (
	"go-gin/mdemo1"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	e := gin.Default()
	// todo: 如果要返回html模版,必须设置模版目录 html模版放到templates下
	// 如果你在templates文件夹下有一个html 你可以设置/templates/*
	// 如果你在templates文件夹下有一个html,还有一个文件夹，文件夹下有html 你可以设置/templates/*/*
	// 如果你在templates文件夹下有一层文件夹,文件夹下有html 你可以设置/templates/**/*
	e.LoadHTMLGlob("templates/*/*")
	mdemo1.TestRouteMain(e)
	// 监听并在 0.0.0.0:8000 上启动服务 不传默认8080
	e.Run(":8000")
}
```

- mdemo1

```go
package mdemo1

import "github.com/gin-gonic/gin"

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestRouteMain(e *gin.Engine) {
	e.GET("/", func(c *gin.Context) {
		c.String(200, "值:%v", "hello gin") //返回string
	})

	e.GET("/json", func(c *gin.Context) {
		data := &User{
			Name: "张三",
			Age:  18,
		}
		c.JSON(200, data)
	})
	e.GET("/json1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "张三",
		})
	})
	e.POST("/add", func(c *gin.Context) {
		c.String(200, "add")
	})
	e.PUT("/put", func(c *gin.Context) {
		c.String(200, "put-----1123")
	})
	e.DELETE("/delete", func(c *gin.Context) {
		c.String(200, "delete")
	})
	// todo: jsonp: /jsonp?callback=fn 会把数据放到fn里
	e.GET("/jsonp", func(c *gin.Context) {
		c.JSONP(200, gin.H{
			"msg":  "jsonp",
			"code": 200,
		})
	})
	// todo: xml
	e.GET("/xml", func(c *gin.Context) {
		c.XML(200, gin.H{
			"msg":  "xml",
			"code": 200,
		})
	})

	// todo: 返回html模版
	e.GET("/html1", func(ctx *gin.Context) {
		ctx.HTML(200, "tem1.html", gin.H{
			"msg":  "msg-html1",
			"name": "name-我是请求html1返回的数据name",
			"code": 200,
		})
	})
	// todo: 如果你在templates文件夹下有两个文件夹admin和default，在这俩文件下都有一个index1.html 这个时候你需要在在各自的index1.html中定义模版名称
	// 定义模版名称 admin下的index1.html中: {{ define "admin/index1.html"}} html {{end}}
	// {{ define "admin/index1.html"}} 定义在html最上面
	// {{end}} 定义在html最下面
	// default下的html同理
	// 在模版目录templates下新建文件夹 一定要对文件夹下的html指定模版名称
	e.GET("/html2", func(ctx *gin.Context) {
		ctx.HTML(200, "default/index1.html", gin.H{
			"msg":  "msg-index1",
			"name": "name-我是请求default/index1.html返回的数据name",
			"code": 200,
		})
	})

	e.GET("/admin/index1", func(ctx *gin.Context) {
		ctx.HTML(200, "admin/index1.html", gin.H{
			"msg":  "msg-index1",
			"name": "name-我是请求admin/index1.html返回的数据name",
			"var":  "这个是在html中用来赋值给变量的值",
			"code": 200,
		})
	})

}
```

- templates/admin/index1.html

```go
{{define "admin/index1.html"}}
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Document</title>
  </head>
  <body>
    <h1>这是admin模版</h1>
    <p>路径：我是templates/admin/index1.html</p>
    <!-- {{.name}} 可以取到返回的name的值 => .返回的字段名称-->
    <div>{{.name}}</div> 
    <div>{{.msg}}</div>
    <div>{{.code}}</div>

    {{$b1 := .var}}
    <section style="border: 1px solid #eee;margin-top:16px;">
       <b>$b1 := .var</b>
      <p>= {{$b1}}</p>
    </section>
  </body>
</html>
{{end}}
```

- templates/default/index1.html

```go
{{define "default/index1.html"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <h1>这是default模版</h1>
</body>
</html>
{{end}}
```

### gin-post & Dynamic Routing

```go
//# mdemo2/index.go
package mdemo2

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Age      int    `json:"age" form:"age"`
}

func TestPostRouteMain(e *gin.Engine) {

	e.GET("/post", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "post/index.html", gin.H{
			"status": "ok",
		})
	})
	e.POST("/post/user", func(ctx *gin.Context) {

		// username := ctx.PostForm("username")
		// password := ctx.PostForm("password")
		// age := ctx.DefaultPostForm("age", "20") // 默认值 不传是20
		// ctx.JSON(http.StatusOK, gin.H{
		// 	"username": username,
		// 	"password": password,
		// 	"age":      age,
		// })

		user := &UserInfo{}
		err := ctx.ShouldBind(&user)
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		} else {
			ctx.JSON(200, user)
		}
		fmt.Printf("userInfo====: %v", user)

	})

	// todo: 绑定结构体  http://localhost:8000/get/user?username=cwt&password=123&age=18
	e.GET("/get/user", func(ctx *gin.Context) {
		user := &UserInfo{}
		if err := ctx.ShouldBind(&user); err == nil {
			ctx.JSON(200, user)
		} else {
			ctx.JSON(400, err.Error())
		}

		fmt.Printf("userInfo: %v", user)
	})
	// todo: 动态路由
	e.GET("/list/:cid", func(ctx *gin.Context) {
		cid := ctx.Param("cid")
		ctx.String(200, "cid: %v", cid)
	})
}

---------------
//# main.go
package main

import (
	"go-gin/mdemo2"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	e := gin.Default()
	// todo: 如果要返回html模版,必须设置模版目录 html模版放到templates下
	// 如果你在templates文件夹下有一个html 你可以设置/templates/*
	// 如果你在templates文件夹下有一个html,还有一个文件夹，文件夹下有html 你可以设置/templates/*/*
	// 如果你在templates文件夹下有一层文件夹,文件夹下有html 你可以设置/templates/**/*
	e.LoadHTMLGlob("templates/*/*")
	// mdemo1.TestRouteMain(e)
	mdemo2.TestPostRouteMain(e)
	// 监听并在 0.0.0.0:8000 上启动服务 不传默认8080
	e.Run(":8000")
}

----
//# post/index.html

{{define "post/index.html"}}
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Document</title>
  </head>
  <body>
    <form action="/post/user" method="post">
      用户名：<input type="text" name="username" /><br />
      密码：<input type="text" name="password" /><br />
      年龄：<input type="text" name="age" /><br />

      <button type="submit">提交</button>
    </form>
  </body>
</html>
{{end}}

```

### Routing Grouping - Pull Out Controller

- 路由分组-抽离控制器

---

**Load Routing Grouping Function**

- 加载路由分组函数

```go
package main

import (
	"go-gin/mdemo2"
	"go-gin/routers"

	"github.com/gin-gonic/gin"
)

// ps1: 加载路由分组函数
func routersGroup(r *gin.Engine) {
	routers.DefaultRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.AdminRoutersInit(r)
}
func main() {
	// 创建一个默认的路由引擎
	e := gin.Default()
	// todo: 如果要返回html模版,必须设置模版目录 html模版放到templates下
	// 如果你在templates文件夹下有一个html 你可以设置/templates/*
	// 如果你在templates文件夹下有一个html,还有一个文件夹，文件夹下有html 你可以设置/templates/*/*
	// 如果你在templates文件夹下有一层文件夹,文件夹下有html 你可以设置/templates/**/*
	e.LoadHTMLGlob("templates/*/*")
	// mdemo1.TestRouteMain(e)
	mdemo2.TestPostRouteMain(e)
	// 监听并在 0.0.0.0:8000 上启动服务 不传默认8080

	routersGroup(e)
	e.Run(":8000")
}
```

**new-built** **routers**

- 新建routers

```go
package routers

import (
	"go-gin/controllers/admin"

	"github.com/gin-gonic/gin"
)

// admin:
func AdminRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/admin")
	apiRouters.GET("/user", admin.UserController{}.Index)
	apiRouters.POST("/user/add", admin.UserController{}.Add)
	apiRouters.PUT("/user/put", admin.UserController{}.Put)
	apiRouters.DELETE("/user/delete", admin.UserController{}.Delete)
}
```

**new-built controllers**

- 新建controllers

```go
package admin

import (
	"go-gin/controllers/common"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	common.BaseController
}

// todo: 把方法绑定到结构体上
func (con UserController) Index(r *gin.Context) {
	// r.String(200, "get admin-user")
	con.Success(r)
}

func (con UserController) Add(r *gin.Context) {
	r.String(200, "post admin-user/add")
}

func (con UserController) Put(r *gin.Context) {
	r.String(200, "put admin-user/put")
}
func (con UserController) Delete(r *gin.Context) {
	r.String(200, "delete admin-user/delete")
}
```

**new-built common**

- 新建common

```go
package common

import "github.com/gin-gonic/gin"

type BaseController struct{}

func (con BaseController) Success(c *gin.Context) {
	c.String(200, "success")
}

func (con BaseController) Error(c *gin.Context) {
	c.String(200, "Error")
}
```


### 中间件

- c.Next()

```go
-- main
func routersGroup(r *gin.Engine) {
	routers.Test(r)
}

func initMiddleware(r *gin.Context) {
	fmt.Println("aaaa-initMiddleware-中间件")
}

func main() {
	// 创建一个默认的路由引擎
	e := gin.Default()
	// todo: 如果要返回html模版,必须设置模版目录 html模版放到templates下
	// 如果你在templates文件夹下有一个html 你可以设置/templates/*
	// 如果你在templates文件夹下有一个html,还有一个文件夹，文件夹下有html 你可以设置/templates/*/*
	// 如果你在templates文件夹下有一层文件夹,文件夹下有html 你可以设置/templates/**/*
	e.LoadHTMLGlob("templates/*/*")
	// mdemo1.TestRouteMain(e)
	// 监听并在 0.0.0.0:8000 上启动服务 不传默认8080

	routersGroup(e)
	e.Run(":8000")
}

-- routers
package routers

import (
	"fmt"
	"go-gin/controllers/test"
	"time"

	"github.com/gin-gonic/gin"
)

// 中间件
func initMiddleware(r *gin.Context) {
	startTime := time.Now().Nanosecond()
	fmt.Println("【我是第一个执行的】 aaaa-initMiddleware-中间件")
	// 调用改请求的剩余处理程序
	r.Next() //todo: 等待数据返回之后才会执行下面的代码
	fmt.Println("【我是最后执行的】!!!!!!!!-我是等数据返回之后 最后才执行的")
	endTime := time.Now().Nanosecond()
	fmt.Println("统计请求的时间======:", endTime-startTime)
}

func Test(r *gin.Engine) {
	testRouters := r.Group("/test")
	testRouters.GET("/user", initMiddleware, test.TestController{}.Index)
}

-- controllers
package test

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type TestController struct{}

func (con TestController) Index(r *gin.Context) {
	fmt.Println("【我是第二个执行的】我是返回的数据 get test")
	time.Sleep(time.Second)
	r.String(200, "get test")
}
```

- c.Abort()

```go
-- main
func routersGroup(r *gin.Engine) {
	routers.Test(r)
}

func initMiddleware(r *gin.Context) {
	fmt.Println("aaaa-initMiddleware-中间件")
}

func main() {
	// 创建一个默认的路由引擎
	e := gin.Default()
	// todo: 如果要返回html模版,必须设置模版目录 html模版放到templates下
	// 如果你在templates文件夹下有一个html 你可以设置/templates/*
	// 如果你在templates文件夹下有一个html,还有一个文件夹，文件夹下有html 你可以设置/templates/*/*
	// 如果你在templates文件夹下有一层文件夹,文件夹下有html 你可以设置/templates/**/*
	e.LoadHTMLGlob("templates/*/*")
	// mdemo1.TestRouteMain(e)
	// 监听并在 0.0.0.0:8000 上启动服务 不传默认8080

	routersGroup(e)
	e.Run(":8000")
}

-- routers
package routers

import (
	"fmt"
	"go-gin/controllers/test"
	"time"

	"github.com/gin-gonic/gin"
)

// 中间件
func initMiddleware(r *gin.Context) {
	startTime := time.Now().Nanosecond()
	fmt.Println("【我是第一个执行的】 aaaa-initMiddleware-中间件")
	// 调用改请求的剩余处理程序
	r.Abort() //todo: 返回数据不会执行 只会执行当前中间件里面的
	fmt.Println("【我是最后执行的】!!!!!!!!-我是等数据返回之后 最后才执行的")
	endTime := time.Now().Nanosecond()
	fmt.Println("统计请求的时间======:", endTime-startTime)
}

func Test(r *gin.Engine) {
	testRouters := r.Group("/test")
	testRouters.GET("/user", initMiddleware, test.TestController{}.Index)
}

-- controllers
package test

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type TestController struct{}

func (con TestController) Index(r *gin.Context) {
	fmt.Println("【我是第二个执行的】我是返回的数据 get test")
	time.Sleep(time.Second)
	r.String(200, "get test")
}
```

### 配置多个中间件

- 执行顺序 两个中间件都是c.Next()的话 ｜第一个c.Next上的数据 ⇒第二个c.Next上的数据⇒ 要返回的数据执行⇒第二个c.Next下的数据⇒第一个c.Next下的数据

```go
-- main 同上
-- controllers 同上
-- router
package routers

import (
	"fmt"
	"go-gin/controllers/test"
	"time"

	"github.com/gin-gonic/gin"
)
func initMiddlewareNext(r *gin.Context) {
	startTime := time.Now().Nanosecond()
	fmt.Println("【我是第一个执行的】 aaaa-initMiddleware-中间件")
	// 调用改请求的剩余处理程序
	r.Next() //todo: 等待数据返回之后才会执行下面的代码
	fmt.Println("【我是最后执行的】!!!!!!!!-我是等数据返回之后 最后才执行的")
	endTime := time.Now().Nanosecond()
	fmt.Println("统计请求的时间======:", endTime-startTime)
}
func initMiddlewareAbort(r *gin.Context) {
	startTime := time.Now().Nanosecond()
	fmt.Println("【我是第一个执行的】 aaaa-initMiddleware-中间件")
	// 调用改请求的剩余处理程序
	r.Abort() //todo: 等待数据返回之后才会执行下面的代码
	fmt.Println("【我是最后执行的】!!!!!!!!-我是等数据返回之后 最后才执行的")
	endTime := time.Now().Nanosecond()
	fmt.Println("统计请求的时间======:", endTime-startTime)
}

func Test(r *gin.Engine) {
	testRouters := r.Group("/test")
	testRouters.GET("/user", initMiddlewareNext, initMiddlewareAbort, test.TestController{}.Index)
}
```

### 全局中间件

- e.Use(x1,x2,…..)

```go
package main

import (
	"fmt"
	"go-gin/mdemo2"
	"go-gin/routers"

	"github.com/gin-gonic/gin"
)

func routersGroup(r *gin.Engine) {
	routers.DefaultRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.AdminRoutersInit(r)
	routers.Test(r)
}

func initMiddleware(r *gin.Context) {
	fmt.Println("全局中间件")
}
func initMiddlewareTwo(r *gin.Context) {
	fmt.Println("全局中间件2")
}

func main() {
	// 创建一个默认的路由引擎
	e := gin.Default()
	e.Use(initMiddleware, initMiddlewareTwo)   //todo: 全局中间件
	// todo: 如果要返回html模版,必须设置模版目录 html模版放到templates下
	// 如果你在templates文件夹下有一个html 你可以设置/templates/*
	// 如果你在templates文件夹下有一个html,还有一个文件夹，文件夹下有html 你可以设置/templates/*/*
	// 如果你在templates文件夹下有一层文件夹,文件夹下有html 你可以设置/templates/**/*
	e.LoadHTMLGlob("templates/*/*")
	// mdemo1.TestRouteMain(e)
	mdemo2.TestPostRouteMain(e)
	// 监听并在 0.0.0.0:8000 上启动服务 不传默认8080

	routersGroup(e)
	e.Run(":8000")
}
```

### 路由分组中间件配置

```go
-- main 同上
package routers

import (
	"go-gin/controllers/api"
	"go-gin/middlewares"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api", middlewares.InitMiddleware) //第一种方式配置
  // apiRouters.Use(middlewares.InitMiddleware)   //第二种方式配置
	apiRouters.GET("/user", api.UserIndex)
	apiRouters.POST("/user/add", api.UserAdd)
	apiRouters.PUT("/user/put", api.UserPut)
	apiRouters.DELETE("/user/delete", api.UserDelete)
}
// * 只会在请求/api/* 下才会执行中间件

-- middleware/init.go
package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(c *gin.Context) {
	fmt.Println(time.Now())
	fmt.Println(c.Request.URL)
}
```

### 中间件和控制器之间共享数据

- c.Set()  设置数据
- c.Get()  获取数据

```go
value, exists := r.Get("username")
	fmt.Printf("我是控制器 获取中间件的value: %v, exists: %v\n", value, exists)

	// 类型断言
	v, ok := value.(string)
	if ok == true {
		r.String(200, "get api-user111"+v)
	} else {
		r.String(400, "获取数据失败", value)
	}
```

```go
--middlewares/init.go

package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(c *gin.Context) {
	fmt.Println(time.Now())
	fmt.Println(c.Request.URL)
	c.Set("username", "kingcwt 我是来自中间件的username")

	// 使用goroutine 使用协程来执行注意事项
	cCp := c.Copy() // 如果要使用c 就是Context的话 需要复制一份
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("done1===" + cCp.Request.URL.Path)
	}()
}

-- routers/apiRouters.go

package routers

import (
	"go-gin/controllers/api"
	"go-gin/middlewares"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api", middlewares.InitMiddleware)
	// apiRouters.Use(middlewares.InitMiddleware)
	apiRouters.GET("/user", api.UserIndex)
	apiRouters.POST("/user/add", api.UserAdd)
	apiRouters.PUT("/user/put", api.UserPut)
	apiRouters.DELETE("/user/delete", api.UserDelete)
}

-- controllers/api/userController.go

package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func UserIndex(r *gin.Context) {
	value, exists := r.Get("username")
	fmt.Printf("我是控制器 获取中间件的value: %v, exists: %v\n", value, exists)

	// 类型断言
	v, ok := value.(string)
	if ok == true {
		r.String(200, "get api-user111"+v)
	} else {
		r.String(400, "获取数据失败", value)
	}
}

func UserAdd(r *gin.Context) {
	r.String(200, "post api-user/add")
}

func UserPut(r *gin.Context) {
	r.String(200, "put api-user/put")
}
func UserDelete(r *gin.Context) {
	r.String(200, "delete api-user/delete")
}
```

### 自定义models

- 放置公共的工具函数

```go
--models/tools.go

package models

import "time"

func GetFormattedTime() string {
	layout := "2006-01-02 15:04:05"
	return time.Now().Format(layout)
}

--api/userController.go

package api

import (
	"fmt"
	"go-gin/models"

	"github.com/gin-gonic/gin"
)

func UserGetIndex(r *gin.Context) {
	nowTime := models.GetFormattedTime()
	fmt.Printf("获取的时间: %v", nowTime)
	r.String(200, "get api/unix: %v", nowTime)
}

-- routers/apiRouters.go

package routers

import (
	"go-gin/controllers/api"
	"go-gin/middlewares"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api", middlewares.InitMiddleware)

	apiRouters.GET("/time", api.UserGetIndex)
}

--main
package main

import (
	"fmt"
	"go-gin/mdemo2"
	"go-gin/routers"

	"github.com/gin-gonic/gin"
)

func routersGroup(r *gin.Engine) {
	routers.DefaultRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.AdminRoutersInit(r)
	routers.Test(r)
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

	routersGroup(e)
	e.Run(":8000")
}
```

### 单文件和多文件上传

```go
--controllers/test.go
package test

import (
	"fmt"
	"net/http"
	"path"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type TestController struct{}

func (con TestController) Index(r *gin.Context) {
	fmt.Println("【我是第二个执行的】我是返回的数据 get test")
	time.Sleep(time.Second)
	r.String(200, "get test")
}

// 单文件上传
func fileUpload(r *gin.Context) {
	username := r.PostForm("username")
	file, err := r.FormFile("files") //单文件

	if err != nil {
		fmt.Println(err)
	} else {
		dst := path.Join("./static/upload/", file.Filename)
		r.SaveUploadedFile(file, dst)
	}
	r.JSON(200, gin.H{
		"username": username,
		"file":     file.Filename,
		"size":     file.Size,
		"type":     file.Header["Content-Type"],
		"key":      file.Filename,
		"createAt": time.Now().Format("2006-01-02 15:04:05"),
		"updateAt": time.Now().Format("2006-01-02 15:04:05"),
		"deleteAt": time.Now().Format("2006-01-02 15:04:05"),
		"status":   "ok",
	})

}

// 多文件上传
func MultipartFileUpload(r *gin.Context) {
	username := r.PostForm("username")
	form, err := r.MultipartForm()
	if err != nil {
		fmt.Println(err)
	} else {
		files := form.File["files"]
		for _, file := range files {
			filename := filepath.Base(file.Filename)
			dst := path.Join("./static/upload/", filename)
			if err := r.SaveUploadedFile(file, dst); err != nil {
				r.String(http.StatusBadRequest, "upload file err: %s", err.Error())
				return
			}
		}
		r.String(http.StatusOK, "Uploaded successfully %d files with fields name=%s", len(files), username)

	}
}

func (con TestController) Upload(r *gin.Context) {
	fileUpload(r)
	// MultipartFileUpload(r)

}

func (con TestController) GetUpload(r *gin.Context) {
	r.HTML(200, "test/index1.html", gin.H{
		"num":  8 << 20,
		"num2": 8 << 10,
	})
}

-- routers/test.go
package routers

import (
	"fmt"
	"go-gin/controllers/test"

	"github.com/gin-gonic/gin"
)

// 中间件

func initMiddlewareNext(r *gin.Context) {
	fmt.Println("【ssssss1】")
	// 调用改请求的剩余处理程序
	r.Next() //todo: 等待数据返回之后才会执行下面的代码
	fmt.Println("【eeeeeee1】")
}
func initMiddlewareAbort(r *gin.Context) {
	fmt.Println("【sssss2】")
	// 调用改请求的剩余处理程序
	r.Next() //todo: 等待数据返回之后才会执行下面的代码
	fmt.Println("【eeeee2】")
}

func Test(r *gin.Engine) {
	testRouters := r.Group("/test")
	testRouters.GET("/user", initMiddlewareNext, initMiddlewareAbort, test.TestController{}.Index)
	testRouters.GET("/get/upload", test.TestController{}.GetUpload)
	testRouters.POST("/upload", test.TestController{}.Upload)
}

-- templates/test/index1.html
{{define "test/index1.html"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <h2>文件上传</h2>
    <p>num:{{.num}}</p>
    <p>num2:{{.num2}}</p>
    <form action="/test/upload" method="POST" enctype="multipart/form-data">
        用户名：<input type="text" name="username"><br />
        头像：<input type="file" name="files"><br />
        头像2：<input type="file" name="files"><br />
        <input type="submit" value="上传">
    </form>
</body>
</html>
{{end}}
-- main.go

package main

import (
	"fmt"
	"go-gin/mdemo2"
	"go-gin/routers"

	"github.com/gin-gonic/gin"
)

func routersGroup(r *gin.Engine) {
	routers.DefaultRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.AdminRoutersInit(r)
	routers.Test(r)
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

	routersGroup(e)
	e.Run(":8000")
}
```

### 按日期存储上传的图片

```go
-- routers/test.go
func Test(r *gin.Engine) {
	testRouters := r.Group("/test")
	testRouters.GET("/get/file", test.TestController{}.GetFile)
	testRouters.POST("/upload/file", test.TestController{}.UploadFile)
}

-- controllers/test/index.go
package test

import (
	"fmt"
	"go-gin/models"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type TestController struct{}

// 图片上传和根据当前时间分组保存
func (con TestController) GetFile(r *gin.Context) {
	r.HTML(200, "test/index2.html", gin.H{})
}
func (con TestController) UploadFile(r *gin.Context) {
	username := r.PostForm("username")
	file, err := r.FormFile("files")
	if err != nil {
		fmt.Println(err)
		return
	}
	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
		".svg":  true,
	}
	if _, ok := allowExtMap[extName]; !ok {
		r.String(200, "文件格式不正确")
		return
	}
	day := models.GetDay()
	dir := "./static/upload/" + day

	fileDir, err := os.Open(dir)
	if err != nil {
		error := os.MkdirAll(dir, os.ModePerm)
		if error != nil {
			r.String(200, "创建文件夹失败")
			return
		}
	}
	defer fileDir.Close()
	filName := models.GetFormattedTime() + extName
	dst := path.Join(dir, filName)
	if err := r.SaveUploadedFile(file, dst); err != nil {
		r.String(200, "上传文件失败")
		return
	}
	fmt.Println("----------------------------------------返回数据执行")
	// r.Redirect(http.StatusMovedPermanently, "/test/get/file")
	r.JSON(200, gin.H{
		"username": username,
		"file":     filName,
		"size":     file.Size,
		"createAt": time.Now().Format("2006-01-02 15:04:05"),
		"msg":      "上传成功 3秒钟后 跳转原页面",
	})
}

-- models/tools.go
package models

import "time"

func GetFormattedTime() string {
	layout := "2006-01-02 15:04:05"
	return time.Now().Format(layout)
}

func GetDay() string {
	layout := "2006-01-02"
	return time.Now().Format(layout)
}

-- templates/test/index2.html

{{define "test/index2.html"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <h2>文件上传</h2>
    <p>num:{{.num}}</p>
    <p>num2:{{.num2}}</p>
    <form action="/test/upload/file" method="POST" enctype="multipart/form-data">
        用户名：<input type="text" name="username"><br />
        头像：<input type="file" name="files"><br />
        <input type="submit" value="上传">
    </form>
</body>
</html>
{{end}}

-- main.go 同上
```

### cookie

```go
c.SetCookie(name,value string, maxAge int, path,domain string,secure,httpOnly bool)
// 名称,值,过期时间秒,路径,域,是否https,是否在客户端可操作

// cookie 在二级域名下共享 a.chr.com 可以访问 b.chr.com也可以访问
c.SetCookie("username","kingcwt",3600,"/",".chr.com",false,true)
```

- sercure 为true时,cookie在http中是无效的,在https中才有效
- httpOnly,是微软对cookie做的扩展,如果在cookie中设置了httpOnly属性,则通过程序(js脚本,applet等)将无法读取到cookie信息,放置xss攻击产生

```go
-- router
func Test(r *gin.Engine) {
	testRouters := r.Group("/test")
	testRouters.GET("/set/cookie", test.TestController{}.SetCookie)
	testRouters.GET("/get/cookie", test.TestController{}.GetCookie)
	testRouters.GET("/delete/cookie", test.TestController{}.DeleteCookie)

}

-- controller
func (c TestController) SetCookie(r *gin.Context) {
	r.SetCookie("name", "kingcwt", 10, "/", "localhost", false, true)
	r.String(http.StatusOK, "ok")
}

func (c TestController) GetCookie(r *gin.Context) {
	s, err := r.Cookie("name")
	if err != nil {
		fmt.Println(err)
		r.String(http.StatusOK, "获取cookie失败")
		return
	}
	r.String(http.StatusOK, "cookie="+s)
}
func (c TestController) DeleteCookie(r *gin.Context) {
	r.SetCookie("name", "kingcwt", -1, "/", "localhost", false, true)
	r.String(http.StatusOK, "删除cookie成功")
}

```

### Session

- 介绍

session是另一种记录客户端状态的机制,不同的是Cookie保存在客户端浏览器中,而Session保存在服务器上

- 流程

当客户端浏览器第一次访问服务器并发送请求时,服务器端会创建一个session对象,生成一个类似于key,value的键值对,然后将value保存到服务器。将key(cookie)返回到浏览器(客户端),浏览器下次访问会携带key(cookie),找到对应的session(value)

 https://github.com/gin-contrib/sessions

- cookie
- memstore
- redis
- memcached
- mongodb

**全局引入Session中间件**

- `go get github.com/gin-contrib/sessions`

```go
// 直接引入gin会自动安装包或则手动安装
"github.com/gin-contrib/sessions"
"github.com/gin-contrib/sessions/cookie"

--main.go
package main

import (
	"fmt"
	"go-gin/mdemo2"
	"go-gin/routers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func routersGroup(r *gin.Engine) {
	routers.DefaultRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.AdminRoutersInit(r)
	routers.Test(r)
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

	// 配置Session中间件
	store := cookie.NewStore([]byte("secret"))
	e.Use(sessions.Sessions("mysession", store))
	routersGroup(e)
	e.Run(":80")
}
```

**详细步骤:**

- session中间件包地址：https://github.com/gin-contrib/sessions
- 下载session包

```go
go get github.com/gin-contrib/sessions
```

- 在main.go引入和设置session中间件

```go
package main

import (
	
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	e := gin.Default() // 默认中加载了Logger(), Recovery() 中间件
	e.Use(initMiddleware)
	e.LoadHTMLGlob("templates/*/*")
	// todo: 配置Session中间件 创建基于cookie的存储引擎 secret 是加密的密钥
	store := cookie.NewStore([]byte("secret")) // store是存储引擎
	e.Use(sessions.Sessions("mysession", store))
	routersGroup(e)
	e.Run(":80")
}
```

- 在routers设置对应的接口地址

```go
package routers

import (
	"fmt"
	"go-gin/controllers/test"

	"github.com/gin-gonic/gin"
)

func Test(r *gin.Engine) {
	testRouters := r.Group("/test")
	testRouters.GET("/get/session", test.TestController{}.GetSession)
	testRouters.GET("/set/session", test.TestController{}.SetSession)
	// testRouters.GET("/delete/session", test.TestController{}.DeleteCookie)
}
```

- 在对应的controller里设置和获取session

```go
package test

import (
	"fmt"
	"go-gin/models"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type TestController struct{}

func (c TestController) SetSession(r *gin.Context) {
	session := sessions.Default(r)
	session.Set("age", "27")
	session.Save()
	r.String(200, "设置session成功")
}

func (c TestController) GetSession(r *gin.Context) {

	session := sessions.Default(r)
	age := session.Get("age")
	r.String(200, "获取session成功", age)
}
```

- 下载redis 。首先，确保您已安装 Homebrew。从终端运行：

```bash
brew install redis
```

- 启动redis

```bash
redis-server
```

- 停止redis 请输入 **Ctrl-C**

****使用 launchd 启动和停止 Redis 在后台启动进程****

```bash
brew services start redis
```

这将启动 Redis 并在登录时重新启动。**`launchd`**您可以通过运行以下命令来检查托管 Redis 的状态：

```bash
brew services info redis
```

如果服务正在运行 您将看到以下输出：

```bash
redis (homebrew.mxcl.redis)
Running: ✔
Loaded: ✔
Schedulable: ✘
User: cuihongran
PID: 20893
```

要停止该服务，请运行：

```bash
brew services stop redis
```

****连接到 Redis****

Redis 运行后，您可以通过运行以下命令来测试它**`redis-cli`**：

```bash
redis-cli
```

- 配置redis存储引擎
- 下载session包下对应的redis引擎

```bash
go get github.com/gin-contrib/sessions/redis@v0.0.5
```

- 设置redis存储引擎

```bash
store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret")) // 使用redis localhost:6379是本地的redis地址
```

- 设置数据

```bash
set key1 kingcwt
```

- 查看数据

```bash
keys * 查看所有
get key1 根据键名获取对应的值
// "kingcwt"
```

设置redis以后 这样在请求的时候 session就会存储到redis数据库中 可以通过keys *来查看是否设置成功

- 设置过期时间

```go
func (c TestController) SetSession(r *gin.Context) {
	session := sessions.Default(r)
  // 设置过期时间
	session.Options(sessions.Options{ 
		MaxAge: 10, // //6hrs MaxAge单位是秒
		// Path:   "/",
		// HttpOnly: true,
		// Secure: false,
	})
	session.Set("age", "27")
	session.Save()
	r.String(200, "设置session成功")
}
```

### MySQL字段类型，查询语句详解

- 创建school数据库

```sql
create database school;
```

**MySQL字段的常用数据类型**

- 整数型：TINYINT,SMALLINT,MEDIUMINT,INT,BIGINT  占用字节数[1,2,3,4,8] ,最大长度[4,6,8,11,20] 长度中第一位是符号
- 浮点数：FLOAT,DOUBLE,DECIMAL(M,D)
- 字符型：CHAR,VARCHAR【varchar是动态的,占用字节数和你保存的字符长度有关,char是固定的长度】
- 备注型：TINYTEXT,TEXT,MEDIUMTEXT,LONGTEXT 占用的字节数[255,65535,16777215,4294967295] 【一个汉字占用3个字节】

---

- 创建表

```sql
CREATE TABLE class (
    id INT(11) NOT NULL AUTO_INCREMENT,
    name VARCHAR(255),
    email VARCHAR(255),
    score TINYINT(4),
    PRIMARY KEY (id)
);
```

- 新增10条数据

```sql
INSERT INTO class (name, email, score) VALUES
    ('John Doe', 'john.doe@example.com', 85),
    ('Jane Smith', 'jane.smith@example.com', 92),
    ('Mike Johnson', 'mike.johnson@example.com', 78),
    ('Sarah Lee', 'sarah.lee@example.com', 88),
    ('Alex Brown', 'alex.brown@example.com', 76),
    ('Emily Davis', 'emily.davis@example.com', 95),
    ('James Wilson', 'james.wilson@example.com', 80),
    ('Lisa Chen', 'lisa.chen@example.com', 89),
    ('Kevin Wang', 'kevin.wang@example.com', 82),
    ('Anna Taylor', 'anna.taylor@example.com', 91);
```

- 查找null或则””

```sql
// OR 或的意思
select * from class where email is null OR email="";

// AND 同时成立的意思
select * from class where score>=70 AND score<90;
```

- 查找分数最高的这一条数据

```sql
select * from class where score in(select max(score) from class); // 查找分数最高的一条数据
```

- MySQL别名

```sql
select name as a from class;
```

### gorm

[GORM](https://gorm.io/zh_CN/)

- 下载gorm包

```sql
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

- 数据库连接

```go
--models/core.go

package models

// document: https://gorm.io/zh_CN/docs/connecting_to_the_database.html
import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func int() {
	dsn := "root:password@qq.com@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("数据库连接失败======error: %v\n", err)
	}
}
```

- CRUD

```go

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

======

package routers

import (
	"golang-gin-note/controllers/users"

	"github.com/gin-gonic/gin"
)

func UsersRouters(r *gin.Engine) {
	routes := r.Group("/users")
	routes.GET("/list", users.UsersController{}.GetList)
}

==============
package users

import (
	"fmt"
	"golang-gin-note/models"
	"time"

	"github.com/gin-gonic/gin"
)

type UsersController struct{}

func find() []models.Users {
	userList := []models.Users{}
	// models.DB.First(&userList)
	models.DB.Where("age>20").Find(&userList)
	return userList
}

func create() string {
	user := models.Users{
		Username: "李四",
		Age:      30,
		Email:    "lisi@qq.com",
		AddTime:  time.Now().Unix(),
	}

	d := models.DB.Create(&user)
	fmt.Println(d)
	return "创建成功"
}

func update() string {

	user := models.Users{}
	models.DB.Where("id = ?", 2).Find(&user)
	user.Username = "张三在写代码"
	user.Email = "zzzz@qq.com"
	user.AddTime = time.Now().Unix()
	models.DB.Save(&user)

	return "更新成功"
}

func update2() string {

	user := models.Users{Id: 1}
	models.DB.Find(&user)
	user.Username = "king"
	user.Email = "king@qq.com"
	models.DB.Save(&user)

	return "更新成功"
}

func update3() string {
	user := models.Users{}
	models.DB.Model(&user).Where("id = ?", 3).Updates(models.Users{Username: "李四update", Email: "lisiupdate@qq.com"})
	return "更新成功1"
}

func delete() string {
	user := models.Users{}
	models.DB.Where("id = ?", 3).Delete(&user)
	return "删除成功"
}

func (u UsersController) GetList(r *gin.Context) {

	// 查询数据库
	// data := find()
	// data := create()
	data := delete()
	r.JSON(200, gin.H{
		"data": data,
	})
}

=====

package models

// document: https://gorm.io/zh_CN/docs/connecting_to_the_database.html
import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	dsn := "root:1003835955@qq.com@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println("ddddddddd1", DB)
	if err != nil {
		fmt.Printf("数据库连接失败======error: %v\n", err)
		log.Fatal(err)
	}
}

===== 
package models

// 结构体名称首字母必须大写 并和数据库表名称对应
type Users struct {
	Id       int32
	Username string
	Age      int16
	Email    string
	AddTime  int64
}

// 配置操作数据库表名称
func (Users) TableName() string {
	return "users"
}
```

### gin-gorm 详细CRUD操作和执行原生SQL语句

```go
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
	routers.NavsRouters(r)
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

-- 
package routers

import (
	"golang-gin-note/controllers/navs"

	"github.com/gin-gonic/gin"
)

func NavsRouters(r *gin.Engine) {
	routes := r.Group("/navs")
	routes.GET("/list", navs.NavsController{}.GetList)
}

--
package navs

import (
	"golang-gin-note/models"
	"time"

	"github.com/gin-gonic/gin"
)

type NavsController struct{}

// 临时测试
type NavJson struct {
	Id    int32  `gorm:"primary_key"`
	Title string `json:"title"`
}

func (NavJson) TableName() string {
	return "navs"
}

func find(c *gin.Context) {
	// navResult := []models.Navs{}
	// AND
	// models.DB.Where("id > ? AND id < ?", 1, 9).Find(&navResult)

	// IN
	// models.DB.Where("id in (?) ", []int{3, 5}).Find(&navResult)

	// like 查询标题里面 包含代码的内容
	// models.DB.Where("title like ? ", "%代码%").Find(&navResult)

	// between 查询id在 5-7之间
	// models.DB.Where("id between ? and ?", 5, 7).Find(&navResult)

	// or 查询id = 3 or id = 6
	// models.DB.Where("id = ? or id = ?", 3, 6).Find(&navResult)
	// models.DB.Where("id = ?", 2).Or("id = ?", 6).Or("id = ?", 3).Find(&navResult)

	// select 指定返回的字段
	// navJsonResult := []NavJson{}
	// models.DB.Select("id,title").Find(&navJsonResult)

	// order 排序 limit offset   // desc 降序  // limit 2 展示降序前两条  // offset 2 跳过前两条数据
	// models.DB.Order("id desc").Offset(2).Limit(2).Find(&navResult)

	// count 统计数量
	// var num int64
	// models.DB.Find(&navResult).Count(&num)
	// c.JSON(200, gin.H{
	// 	"data":  navResult,
	// 	"count": num,
	// })

}

func nativeSqlOperat(c *gin.Context) {
	navResult := []models.Navs{}

	// --- 使用原生sql ---
	// todo Exec 执行sql语句 Raw 查询sql语句

	// 删除一条数据
	// models.DB.Exec("delete from navs where id = ?", 4)

	// 修改数据
	// models.DB.Exec("update navs set title = ? where id = ?", "id1的在大半夜学习golang", 1)

	// 查询数据
	models.DB.Raw("select * from navs where id > 5").Scan(&navResult)

	// 统计数据
	var num int64
	models.DB.Raw("select count(*) from navs").Scan(&num)
	models.DB.Find(&navResult)
	c.JSON(200, gin.H{
		"data":  navResult,
		"count": num,
	})
}

func create(c *gin.Context) {
	nav := models.Navs{
		Title:   "小米在写代码",
		Url:     "http://www.baidu.com",
		Status:  7,
		AddTime: time.Now().Unix(),
	}

	models.DB.Create(&nav)
	c.JSON(200, gin.H{
		"data": nav,
		"msg":  "创建成功",
	})
}

func (n NavsController) GetList(c *gin.Context) {
	nativeSqlOperat(c)
	// find(c)
	// create(c)

}

--
package models

type Navs struct {
	Id      int32  `gorm:"primary_key"`
	Title   string `json:"title"`
	Url     string `json:"url"`
	Status  int32  `json:"status"`
	AddTime int64  `json:"add_time"`
}

func (Navs) TableName() string {
	return "navs"
}
```

### go-gin之Preload 预加载

[预加载](https://gorm.io/zh_CN/docs/preload.html)

- 如果你在Article表中关联了一个ArticleCate表 在执行preload的时候 默认就会加载ArticleCate表中的数据,如果你在Article中指定了一个名字叫ArticleCateId的id,默认会把ArticleCate表中的id作为主键,把当前表中的ArticleCateId作为外键关联起来
- **文章属于哪个分类**

```go
package models

type Article struct {
	Id            int
	Title         string
	ArticleCateId int //外键
	State         int
	ArticleCate   ArticleCate
}

// 当你在关联一个表的时候 比方现在关联了一个ArticleCate表 默认会把ArticleCate的id作为主键 如果你在当前表中设置了ArticleCateId 就会把这个作为外键关联起来
func (Article) TableName() string {
	return "article"
}

---
package models

type ArticleCate struct {
	Id    int
	Title string
	State int
}

func (ArticleCate) TableName() string {
	return "article_cate"
}
```

- 指定外键
- 通过`gorm:"foreignKey:CateId"`指定id. 也可以重新定义返回的关联表名称

```go
package models

type Article struct {
	Id        int
	Title     string
	CateId    int //外键
	State     int
	ArticleCa ArticleCate `gorm:"foreignKey:CateId"`
}

// 当你在关联一个表的时候 比方现在关联了一个ArticleCate表 默认会把ArticleCate的id作为主键 如果你在当前表中设置了ArticleCateId 就会把这个作为外键关联起来
func (Article) TableName() string {
	return "article"
}
```

- **文章分类下包含哪些文章**

```go
package models

type ArticleCate struct {
	Id      int
	Title   string
	State   int
	Article []Article `gorm:"foreignKey:CateId"` // 这里把文章关联到分类下 并指定关联id
}

func (ArticleCate) TableName() string {
	return "article_cate"
}

package models

type Article struct {
	Id     int
	Title  string
	CateId int //外键
	State  int
}

func (Article) TableName() string {
	return "article"
}
```

### gorm多表关联查询

- main.go

```go
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
	routers.StudentRoutersInit(r)
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
```

- routers

```go
package routers

import (
	"golang-gin-note/controllers/student"

	"github.com/gin-gonic/gin"
)

func StudentRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/student")
	apiRouters.GET("/list", student.StudentController{}.Index)
}
```

- controllers

```go
package student

import (
	"golang-gin-note/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StudentController struct{}

func (stu StudentController) Index(c *gin.Context) {
	// 获取学生信息
	// studentList := []models.Student{}
	// models.DB.Find(&studentList)
	// c.JSON(200, gin.H{"data": studentList})

	// 获取所有课 程信息
	// lessonList := []models.Lesson{}
	// models.DB.Find(&lessonList)
	// c.JSON(200, gin.H{"data": lessonList})

	// 查询学生信息的时候 展示学生选修的课程
	// studentList := []models.Student{}
	// models.DB.Preload("Lesson").Find(&studentList)
	// c.JSON(200, gin.H{"data": studentList})
	// 获取学生选修的课程信息
	// studentList := []models.Student{}
	// models.DB.Find(&studentList)
	// c.JSON(200, gin.H{"data": studentList})

	// 查询张三选修了哪些课程
	// studentList := []models.Student{}
	// models.DB.Preload("Lesson").Where("name = ?", "张三").Find(&studentList)
	// c.JSON(200, gin.H{"data": studentList})

	// 课程被哪些学生选修了
	// lessonList := []models.Lesson{}
	// models.DB.Preload("Student").Limit(2).Order("id desc").Find(&lessonList)
	// c.JSON(200, gin.H{"data": lessonList})

	//查询课程被哪些学生选修的同时去掉张三
	// lessonList := []models.Lesson{}
	// models.DB.Preload("Student", "name != ?", "张三").Find(&lessonList)
	// c.JSON(200, gin.H{"data": lessonList})

	//查询课程被哪些学生选修的同时去掉张三和李四
	// lessonList := []models.Lesson{}
	// models.DB.Preload("Student", "id not in (1,2)").Find(&lessonList)
	// c.JSON(200, gin.H{"data": lessonList})

	// 预加载SQL 在查询课程被哪些学生选修的同时 学生数据倒叙排列
	lessonList := []models.Lesson{}
	models.DB.Preload("Student", func(db *gorm.DB) *gorm.DB {
		return models.DB.Where("id > 1").Order("id desc")
	}).Find(&lessonList)
	c.JSON(200, gin.H{"data": lessonList})
}
```

- models

```go
package models

type Lesson struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Student []Student `gorm:"many2many:lesson_student;"`
}

func (Lesson) TableName() string {
	return "lesson"
}

//
package models

type LessonStudent struct {
	LessonId  int
	StudentId int
}

func (LessonStudent) TableName() string {
	return "lesson_student"
}

//
package models

type Student struct {
	Id       int
	Number   string
	Password string
	ClassId  int
	Name     string
	// Lesson   []Lesson `gorm:"many2many:lesson_student;"`
}

func (Student) TableName() string {
	return "student"
}
```

### gin [gorm-事务]

- main.go

```go
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
```

- routers

```go
package routers

import (
	"golang-gin-note/controllers/bank"

	"github.com/gin-gonic/gin"
)

func BankRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/bank")
	apiRouters.GET("/list", bank.BankController{}.Index)
}
```

- controllers

```go
package bank

import (
	"golang-gin-note/controllers/common"
	"golang-gin-note/models"

	"github.com/gin-gonic/gin"
)

type BankController struct {
	common.BaseController
}

func (bank BankController) Index(c *gin.Context) {

	// c.JSON(200, gin.H{
	// 	"code": 200,
	// 	"data": "ok",
	// })
	// 开启事务
	tx := models.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	u1 := models.Bank{Id: 1}
	tx.Find(&u1)
	u1.Balance = u1.Balance + 100
	if err := tx.Save(&u1).Error; err != nil {
		tx.Rollback()
		bank.Error(c)
		return
	}

	u2 := models.Bank{Id: 2}
	tx.Find(&u2)
	u2.Balance = u2.Balance - 100
	if err := tx.Save(&u2).Error; err != nil {
		tx.Rollback()
		bank.Error(c)
		return
	}
	tx.Commit()
	bank.Success(c)

	c.JSON(200, gin.H{
		"u1": u1,
		"u2": u2,
	})
}
```

- models

```go
package models

type Bank struct {
	Id       int
	Username string
	Balance  float32
}

func (Bank) TableName() string {
	return "bank"
}
```