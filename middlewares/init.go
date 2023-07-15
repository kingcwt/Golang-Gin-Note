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
