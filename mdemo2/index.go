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
