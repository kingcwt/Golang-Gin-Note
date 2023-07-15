package routers

import (
	"golang-gin-note/controllers/admin"

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
