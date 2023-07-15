package routers

import (
	"fmt"
	"golang-gin-note/controllers/test"

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
	testRouters.GET("/get/file", test.TestController{}.GetFile)
	testRouters.POST("/upload/file", test.TestController{}.UploadFile)
	testRouters.GET("/set/cookie", test.TestController{}.SetCookie)
	testRouters.GET("/get/cookie", test.TestController{}.GetCookie)
	testRouters.GET("/delete/cookie", test.TestController{}.DeleteCookie)
	testRouters.GET("/get/session", test.TestController{}.GetSession)
	testRouters.GET("/set/session", test.TestController{}.SetSession)
	// testRouters.GET("/delete/session", test.TestController{}.DeleteCookie)
}
