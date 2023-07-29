package routers

import (
	"golang-gin-note/controllers/student"

	"github.com/gin-gonic/gin"
)

func StudentRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/student")
	apiRouters.GET("/list", student.StudentController{}.Index)
}
