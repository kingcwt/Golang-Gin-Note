package api

import (
	"fmt"
	"golang-gin-note/models"

	"github.com/gin-gonic/gin"
)

func UserGetIndex(r *gin.Context) {
	nowTime := models.GetFormattedTime()
	fmt.Printf("获取的时间戳: %v", nowTime)
	r.String(200, "get api/unix: %v", nowTime)
}
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
