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
